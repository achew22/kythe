/*
 * Copyright 2015 The Kythe Authors. All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// Binary http_server exposes HTTP interfaces for the xrefs and filetree
// services backed by a combined serving table.
package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"kythe.io/kythe/go/services/filetree"
	"kythe.io/kythe/go/services/xrefs"
	ftpb "kythe.io/kythe/proto/filetree_go_proto"
	xpb "kythe.io/kythe/proto/xref_go_proto"

	"github.com/google/codesearch/index"
	"github.com/google/codesearch/regexp"
)

type result struct {
	lineNo uint
	ticket string
	line   string
}

type codeSearch struct {
	index       *index.Index
	indexed     bool
	indexingErr error
	buf         []byte

	ft filetree.Service
	xs xrefs.Service
}

func newCodeSearch(ft filetree.Service, xs xrefs.Service) *codeSearch {
	return &codeSearch{
		ft: ft,
		xs: xs,
	}
}

func (cs *codeSearch) Index() {
	defer func() {
		if cs.indexingErr != nil {
			fmt.Fprintf(os.Stderr, "Indexing error: %v", cs.indexingErr)
		}
	}()

	indexTempFile, err := ioutil.TempFile("", "csindex")
	if err != nil {
		cs.indexingErr = fmt.Errorf("error making csindex tempfile: %v", err)
		return
	}

	w := index.Create(indexTempFile.Name())

	ctx := context.Background()

	corpusRoots, err := cs.ft.CorpusRoots(ctx, &ftpb.CorpusRootsRequest{})
	if err != nil {
		cs.indexingErr = fmt.Errorf("error getting corpus roots: %v", err)
		return
	}

	for _, corpusRoot := range corpusRoots.Corpus {
		for _, root := range corpusRoot.Root {
			if err := cs.parseDir(ctx, w, corpusRoot.Name, root, ""); err != nil {
				cs.indexingErr = fmt.Errorf("error parsingDir(%q, %q, /): %v", corpusRoot.Name, root, err)
				return
			}
		}
	}

	w.Flush()
	cs.index = index.Open(indexTempFile.Name())
	cs.indexed = true
}

func (cs *codeSearch) parseDir(ctx context.Context, w *index.IndexWriter, corpus, root, path string) error {
	var err error
	var dir *ftpb.DirectoryReply
	if dir, err = cs.ft.Directory(ctx, &ftpb.DirectoryRequest{
		Corpus: corpus,
		Root:   root,
		Path:   path,
	}); err != nil {
		return fmt.Errorf("cs.ft.Directory(Corpys: %q, Root: %q, Path: %q) error: %v", corpus, root, path, err)
	}

	for _, entry := range dir.Entry {
		newPath := strings.TrimPrefix(path+"/"+entry.Name, "/")
		switch entry.Kind {
		case ftpb.DirectoryReply_DIRECTORY:
			cs.parseDir(ctx, w, corpus, root, newPath)
		case ftpb.DirectoryReply_FILE:
			ticket := fmt.Sprintf("kythe://%s?path=%s",
				corpus,
				url.QueryEscape(newPath))
			var buf io.Reader
			if buf, err = cs.getTicket(ctx, ticket); err != nil {
				return fmt.Errorf("getTicket(%q): %v", ticket, err)
			}
			fmt.Printf("Indexing ticket: %v\n", ticket)
			w.Add(ticket, buf)
		}
	}

	return nil
}

func (cs *codeSearch) getTicket(ctx context.Context, ticket string) (io.Reader, error) {

	file, err := cs.xs.Decorations(ctx, &xpb.DecorationsRequest{
		Location: &xpb.Location{
			Ticket: ticket,
			Kind:   xpb.Location_FILE,
		},
		SourceText: true,
	})
	if err != nil {
		return nil, fmt.Errorf("cs.ft.Decorations(Ticket: %q) error: %v", ticket, err)
	}

	return bytes.NewBuffer(file.SourceText), nil
}

func (cs *codeSearch) Search(w http.ResponseWriter, r *http.Request) {
	if !cs.indexed {
		w.WriteHeader(http.StatusServiceUnavailable)
		w.Write([]byte("The server is not finished indexing yet. Please try again later."))
		return
	}

	if cs.indexingErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("The server errored indexing: %v", cs.indexingErr)))
		return
	}

	w.Header().Set("content-type", "text/html")

	vals, ok := r.URL.Query()["q"]
	if !ok || len(vals) != 1 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Search expects a query string parameter, \"q\" that is your query."))
		return
	}

	query := vals[0]

	pat := "(?m)" + query
	// To do case insensitivity, do this
	//pat = "(?i)" + pat

	g := regexp.Grep{
		Stdout: w,
		Stderr: os.Stderr,
	}

	re, err := regexp.Compile(pat)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("regexp.Compile(%q): %v", pat, err)))
		return
	}
	g.Regexp = re
	q := index.RegexpQuery(re.Syntax)

	post := cs.index.PostingQuery(q)

	for _, fileid := range post {
		ticket := cs.index.Name(fileid)

		u, err := url.Parse(ticket)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf("url.Parse(%q): %v", ticket, err)))
			return
		}

		rawPath := u.Query()["path"][0]
		path, err := url.PathUnescape(rawPath)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf("url.PathUnescape(%q): %v", rawPath, err)))
			return
		}
		corpus := u.Hostname()
		w.Write([]byte(fmt.Sprintf("<h1>%s</h1>", path)))

		w.Write([]byte("<ul>"))
		// Load the object from the index
		var results []result
		if results, err = cs.find(r.Context(), ticket, re); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf("find(%q): %v", ticket, err)))
			return
		}

		filename := filepath.Base(path)

		for _, result := range results {
			w.Write([]byte(fmt.Sprintf("<li><a href=\"/#%s?corpus=%s&signature&line=%d\">%s:%d</a>: %s</li>", path, corpus, result.lineNo, filename, result.lineNo, result.line)))
		}
		w.Write([]byte("</ul>"))
	}

}

var nl = []byte{'\n'}

func countNL(b []byte) uint {
	n := uint(0)
	for {
		i := bytes.IndexByte(b, '\n')
		if i < 0 {
			break
		}
		n++
		b = b[i+1:]
	}
	return n
}

// find searches the corpus for the provided regexp and returns a set of results.
// Liberally stolen from https://github.com/google/codesearch/blob/master/regexp/match.go
func (cs *codeSearch) find(ctx context.Context, ticket string, re *regexp.Regexp) ([]result, error) {
	var res []result

	if cs.buf == nil {
		cs.buf = make([]byte, 1<<20)
	}
	var (
		buf        = cs.buf[:0]
		needLineno = true
		lineno     = uint(1)
		beginText  = true
		endText    = false
	)

	source, err := cs.getTicket(ctx, ticket)
	if err != nil {
		return res, fmt.Errorf("getTicket(%q): %v", ticket, err)
	}

	for {
		n, err := io.ReadFull(source, buf[len(buf):cap(buf)])
		buf = buf[:len(buf)+n]
		end := len(buf)
		if err == nil {
			i := bytes.LastIndex(buf, nl)
			if i >= 0 {
				end = i + 1
			}
		} else {
			endText = true
		}
		chunkStart := 0
		for chunkStart < end {
			m1 := re.Match(buf[chunkStart:end], beginText, endText) + chunkStart
			beginText = false
			if m1 < chunkStart {
				break
			}
			lineStart := bytes.LastIndex(buf[chunkStart:m1], nl) + 1 + chunkStart
			lineEnd := m1 + 1
			if lineEnd > end {
				lineEnd = end
			}
			if needLineno {
				lineno += countNL(buf[chunkStart:lineStart])
			}
			line := buf[lineStart:lineEnd]

			res = append(res, result{
				lineNo: lineno,
				line:   string(line),
				ticket: ticket,
			})
			lineno++
			chunkStart = lineEnd
		}
		if err == nil {
			lineno += countNL(buf[chunkStart:end])
		}
		n = copy(buf, buf[end:])
		buf = buf[:n]
		if len(buf) == 0 && err != nil {
			if err != io.EOF && err != io.ErrUnexpectedEOF {
				return res, fmt.Errorf("%s: %v\n", ticket, err)
			}
			break
		}
	}

	return res, nil
}
