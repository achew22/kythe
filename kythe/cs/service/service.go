/*
 * Copyright 2016 Google Inc. All rights reserved.
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

package service

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"html"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"sort"
	"strings"
	"sync/atomic"
	"unsafe"

	"github.com/google/codesearch/index"
	"github.com/google/codesearch/regexp"
	"kythe.io/kythe/cs/cshtml"
	"kythe.io/kythe/cs/defindex"
	"kythe.io/kythe/cs/refpack"
)

// A loaded corpus index.
type Index struct {
	indexdir string
	idx      *index.Index
	didx     defindex.Index
	files    map[string]cshtml.SrcData
}

func (ix *Index) init() error {
	ix.idx = index.Open(ix.indexdir + "/index")
	didxfh, err := os.Open(ix.indexdir + "/defs")
	if err != nil {
		return err
	}

	dec := gob.NewDecoder(didxfh)
	err = dec.Decode(&ix.didx)
	if err != nil {
		return err
	}

	ix.files = make(map[string]cshtml.SrcData)
	for _, dfile := range ix.didx {
		contents, err := ioutil.ReadFile(ix.indexdir + "/raw/" + dfile.Name)
		if err != nil {
			return err
		}
		ix.files[dfile.Name] = cshtml.MakeSrcData(contents)
	}

	return nil
}

// Returns the name of the directory from which the index was loaded.
func (ix *Index) Dir() string {
	return ix.indexdir
}

type textmatch struct {
	filename           string
	startByte, endByte int
	score              float32
}

type byScore []textmatch

func (b byScore) Len() int {
	return len(b)
}
func (b byScore) Less(i, j int) bool {
	return b[i].score > b[j].score
}
func (b byScore) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func (ix *Index) search(pat string) ([]textmatch, error) {
	var textmatches []textmatch
	const maxmatches int = 100

	patlower := strings.ToLower(pat)

	for _, dfile := range ix.didx {
		for _, def := range dfile.Defs {
			if match := strings.Index(def.FullName, patlower); match != -1 {
				matchEnd := match + len(pat)
				unqnameBegin := strings.LastIndex(def.FullName, "::") + 2
				if unqnameBegin == 1 {
					unqnameBegin = 0
				}
				if matchEnd < unqnameBegin {
					continue
				}
				if match < unqnameBegin {
					match = unqnameBegin
				}
				score := float32(def.RefCount) * float32(matchEnd-match) / float32(len(def.FullName)-unqnameBegin)
				textmatches = append(textmatches, textmatch{dfile.Name, int(def.StartByte), int(def.EndByte), score})
			}
		}
		if fmatch := strings.Index(dfile.NameLower, patlower); fmatch != -1 && fmatch+len(pat) == len(dfile.Name) {
			textmatches = append(textmatches, textmatch{dfile.Name, 0, 0, 0.0})
		}
	}

	sort.Sort(byScore(textmatches))
	if len(textmatches) >= maxmatches {
		return textmatches[:maxmatches], nil
	}

	patre := "(?i)" + pat
	cre, err := regexp.Compile(patre)
	if err != nil {
		return nil, err
	}
	q := index.RegexpQuery(cre.Syntax)
	post := ix.idx.PostingQuery(q)

fileloop:
	for _, fileid := range post {
		name := ix.idx.Name(fileid)
		contents, ok := ix.files[name]
		if !ok {
			panic("File not present in def index")
		}

		offset := 0
		beginText := true
	matchloop:
		for {
			m := cre.Match(contents.Src[offset:], beginText, true)
			beginText = false
			if m == -1 {
				break
			}
			offset += m
			for _, xmatch := range textmatches {
				if xmatch.filename != name {
					continue
				}
				if xmatch.startByte < offset && offset <= xmatch.endByte {
					continue matchloop
				}
			}
			textmatches = append(textmatches, textmatch{name, offset, offset, 0.0})
			if len(textmatches) == maxmatches {
				break fileloop
			}
		}
	}

	return textmatches, nil
}

// An http.Handler that serves a given corpus index.
type Service struct {
	ix *Index
}

// Sets the corpus index served by this Service to ix.
func (s *Service) SetIndex(ix *Index) {
	atomic.StorePointer((*unsafe.Pointer)(unsafe.Pointer(&s.ix)), unsafe.Pointer(ix))
}

func (s *Service) serveSearchResults(ix *Index, w http.ResponseWriter, pat string) {
	cshtml.WriteHeader(w, "Search results for "+pat, pat)
	matches, err := ix.search(pat)
	if err != nil {
		fmt.Fprintf(w, "Invalid regular expression: \"%s\"", html.EscapeString(pat))
	} else if len(matches) == 0 {
		fmt.Fprintln(w, "No matches.")
	} else {
		for _, xmatch := range matches {
			fmt.Fprintln(w, html.EscapeString(xmatch.filename)+"<pre>")
			cshtml.WriteSnippet(w, xmatch.filename, ix.files[xmatch.filename], xmatch.startByte, xmatch.endByte, 1)
			fmt.Fprintln(w, "</pre>")
		}
	}
	cshtml.WriteFooter(w)
}

func (s *Service) serveDirectoryListing(ix *Index, w http.ResponseWriter, r *http.Request) {
	dir := r.URL.Path
	if dir[len(dir)-1] != '/' {
		http.Redirect(w, r, dir+"/", http.StatusFound)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	files, err := ioutil.ReadDir(ix.indexdir + "/raw/" + dir)
	if err != nil {
		fmt.Fprintln(w, "No directory listing available")
	}

	cshtml.WriteHeader(w, "Directory listing for "+dir, "")
	fmt.Fprintln(w, "<pre>")
	if dir != "/" {
		fmt.Fprintln(w, "<a href=\"../\">Parent Directory</a>")
	}
	for _, f := range files {
		name := html.EscapeString(f.Name())
		if f.IsDir() {
			name += "/"
		}
		fmt.Fprintln(w, "<a href=\""+name+"\">"+name+"</a>")
	}
	fmt.Fprintln(w, "</pre>")
	cshtml.WriteFooter(w)
}

func (s *Service) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ix := (*Index)(atomic.LoadPointer((*unsafe.Pointer)(unsafe.Pointer(&s.ix))))
	if ix == nil {
		http.Error(w, "index not ready yet", http.StatusServiceUnavailable)
		return
	}

	if strings.HasPrefix(r.URL.Path, "/!res/") {
		http.ServeFile(w, r, ix.indexdir+"/res/"+r.URL.Path[6:])
		return
	}
	if r.URL.Path == "/favicon.ico" {
		http.NotFound(w, r)
		return
	}

	if r.URL.Path == "/" {
		q, ok := r.URL.Query()["q"]
		if ok {
			s.serveSearchResults(ix, w, q[0])
			return
		}
	}

	if !strings.Contains(r.URL.Path, "/..") {
		htmlpath := ix.indexdir + "/html" + r.URL.Path
		info, err := os.Stat(htmlpath)
		if err == nil {
			if info.IsDir() {
				s.serveDirectoryListing(ix, w, r)
				return
			}
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			http.ServeFile(w, r, htmlpath)
			return
		}

		dir, name := path.Split(r.URL.Path)
		xrefpath := ix.indexdir + "/xrefs" + dir[:len(dir)-1]
		info, err = os.Stat(xrefpath)
		if err == nil && !info.IsDir() {
			f, err := os.Open(xrefpath)
			if err != nil {
				panic("could not open file")
			}

			rr, err := refpack.NewReader(f)
			if err != nil {
				panic("could not read refpack")
			}

			n, err := rr.SeekToName(name)
			if n != -1 {
				content := make([]byte, n)
				rn, err := rr.Read(content)
				if n != rn || err != nil {
					panic("could not read refpack")
				}

				w.Header().Set("Content-Type", "text/html; charset=utf-8")
				http.ServeContent(w, r, "", info.ModTime(), bytes.NewReader(content))
				return
			}
		}
	}

	s.serveSearchResults(ix, w, r.URL.Path[1:])
}

// Loads a corpus index from the given directory and returns it.
func LoadIndex(path string) (*Index, error) {
	var ix Index
	ix.indexdir = path
	err := ix.init()
	if err != nil {
		return nil, err
	}
	return &ix, nil
}
