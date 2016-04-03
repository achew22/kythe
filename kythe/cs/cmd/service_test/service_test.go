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

package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"golang.org/x/net/html"

	"kythe.io/kythe/cs/service"
)

type mockWriter struct {
	header  http.Header
	content bytes.Buffer
	code    int
}

func (w *mockWriter) Header() http.Header {
	if w.header == nil {
		w.header = make(http.Header)
	}
	return w.header
}

func (w *mockWriter) Write(p []byte) (int, error) {
	return w.content.Write(p)
}

func (w *mockWriter) WriteHeader(code int) {
	w.code = code
}

type link struct {
	href, text string
}

func (w *mockWriter) links() ([]link, error) {
	doc, err := html.Parse(bytes.NewReader(w.content.Bytes()))
	if err != nil {
		return nil, err
	}
	var getLinkText func(*html.Node) string
	getLinkText = func(n *html.Node) string {
		if n.Type == html.TextNode {
			return n.Data
		} else {
			var t string
			for c := n.FirstChild; c != nil; c = c.NextSibling {
				t += getLinkText(c)
			}
			return t
		}

	}
	var links []link
	var findLinks func(*html.Node)
	findLinks = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			var href string
			for _, a := range n.Attr {
				if a.Key == "href" {
					href = a.Val
				}
			}
			if href != "" {
				links = append(links, link{href, getLinkText(n)})
			}
		} else {
			for c := n.FirstChild; c != nil; c = c.NextSibling {
				findLinks(c)
			}
		}
	}
	findLinks(doc)
	return links, nil
}

type testHarness struct {
	handler http.Handler
	history []*url.URL
	curPage *mockWriter
}

func (h *testHarness) navigate(rawurl string) error {
	newUrl, err := url.Parse(rawurl)
	if err != nil {
		return err
	}
	if len(h.history) == 0 {
		if !newUrl.IsAbs() {
			return errors.New("first url should be absolute")
		}
	} else {
		newUrl = h.history[len(h.history)-1].ResolveReference(newUrl)
	}
	h.history = append(h.history, newUrl)
	h.requestPage()
	return nil
}

func (h *testHarness) back() {
	h.history = h.history[:len(h.history)-1]
	err := h.requestPage()
	if err != nil {
		log.Fatal("couldn't go back")
	}
}

func (h *testHarness) requestPage() error {
	req, err := http.NewRequest("GET", h.Location().String(), nil)
	if err != nil {
		return err
	}
	newPage := new(mockWriter)
	h.handler.ServeHTTP(newPage, req)
	h.curPage = newPage
	return nil
}

func (h *testHarness) Location() *url.URL {
	return h.history[len(h.history)-1]
}

func (h *testHarness) run(r io.Reader, path string) {
	s := bufio.NewScanner(r)
	lineno := 1
	checkOffset := 0
	for s.Scan() {
		line := s.Text()
		if strings.HasPrefix(line, "// navigate: ") {
			err := h.navigate(line[13:])
			if err != nil {
				log.Fatalf("%s:%d: error while navigating: %s", path, lineno, err.Error())
			}
			checkOffset = 0
		}
		var i, n int
		var pat string
		if sn, err := fmt.Sscanf(line, "// click %d/%d: %s", &i, &n, &pat); err == nil && sn == 3 {
			links, err := h.curPage.links()
			if err != nil {
				log.Fatalf("%s:%d: error extracting links: %s", path, lineno, err.Error())
			}
			var matchingLinks []link
			for _, l := range links {
				if strings.Contains(l.text, pat) {
					matchingLinks = append(matchingLinks, l)
				}
			}
			if len(matchingLinks) != n {
				log.Fatalf("%s:%d: error: expected %d match(es), got %d: %v", path, lineno, n, len(matchingLinks), matchingLinks)
			}
			err = h.navigate(matchingLinks[i].href)
			if err != nil {
				log.Fatalf("%s:%d: error while navigating: %s", path, lineno, err.Error())
			}
			checkOffset = 0
		}
		if line == "// back" {
			h.back()
		}
		if strings.HasPrefix(line, "// check: ") {
			if i := bytes.Index(h.curPage.content.Bytes()[checkOffset:], []byte(line[10:])); i != -1 {
				checkOffset += i + len(line) - 10
			} else {
				log.Fatalf("%s:%d: error: expected string not found", path, lineno)
			}
		}
		if strings.HasPrefix(line, "// url: ") {
			expectedUrl := line[8:]
			if expectedUrl != h.Location().String() {
				log.Fatalf("%s:%d: error: expected url %s, got %s", path, lineno, expectedUrl, h.Location().String())
			}
		}
		if line == "// print" {
			os.Stdout.Write(h.curPage.content.Bytes())
		}
		lineno++
	}
}

func main() {
	var h testHarness

	var s service.Service
	ix, err := service.LoadIndex(os.Args[1])
	if err != nil {
		log.Fatalf("error loading index %s: %s\n", os.Args[1], err)
	}
	s.SetIndex(ix)

	h.handler = &s
	err = h.navigate("http://localhost/")
	if err != nil {
		log.Fatalf("error making first navigation: %s", err.Error())
	}

	f, err := os.Open(os.Args[2])
	if err != nil {
		log.Fatalf("error opening test file %s: %s", os.Args[2], err)
	}
	h.run(f, os.Args[2])
}
