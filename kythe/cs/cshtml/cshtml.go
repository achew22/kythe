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

// Package cshtml contains common utilities for producing HTML that are shared
// between the index and serve programs.
package cshtml

import (
	"fmt"
	"html"
	"io"
)

// WriteHeader writes the standard page header to w.
func WriteHeader(w io.Writer, title, q string) {
	fmt.Fprintf(w, `<!DOCTYPE html>
<html>
<head>
<title>%s</title>
<link rel="stylesheet" type="text/css" href="/!res/style.css">
<script src="/!res/script.js"></script>
</head>
<body onload="init()">
<table style="width: 100%%;">
<tr>
<td style="width: 75%%; text-align: center;">
<form action="/">
<input type="text" name="q" id="searchbox" value="%s" size="50">
<input type="submit" value="Search">
</form>
</td>
<td style="width: 25%%;">
<pre>
<span style="text-decoration: underline;">reference, declaration</span> &#8594; <span style="font-weight: bold;">definition</span>
<span style="font-weight: bold;">definition</span> &#8594; references, declarations, derived classes, virtual overrides
<span class="mref">reference to multiple definitions</span> &#8594; definitions
<span class="unref">unreferenced</span>
</pre>
</td>
</table>`, html.EscapeString(title), html.EscapeString(q))
}

// WriteFooter writes the standard page footer to w.
func WriteFooter(w io.Writer) {
	fmt.Fprintln(w, `</body>
</html>`)
}

// SrcData stores precomputed information about a source file.
type SrcData struct {
	// The content of the source file.
	Src []byte

	// An array whose length is the number of lines in the file.
	// LineTable[i] contains the byte offset of the start of line i+1.
	// This allows us to binary search for the line number corresponding to
	// a byte offset (see SrcData.LineNumber).
	LineTable []int32
}

// MakeSrcData makes a SrcData for src.
func MakeSrcData(src []byte) SrcData {
	if len(src) == 0 {
		return SrcData{src, []int32{0}}
	}

	linecount := 1
	for i := 0; i != len(src)-1; i++ {
		if src[i] == '\n' {
			linecount++
		}
	}

	linetable := make([]int32, linecount)
	linei := 1
	for i := 0; i != len(src)-1; i++ {
		if src[i] == '\n' {
			linetable[linei] = int32(i + 1)
			linei++
		}
	}

	return SrcData{src, linetable}
}

// LineNumber computes the line number for the given byte offset; O(log N) in
// the number of lines.
func (srcd *SrcData) LineNumber(byteOffset int32) int {
	min := 0
	max := len(srcd.LineTable) - 1
	for min != max {
		index := (min + max + 1) / 2
		byteindex := srcd.LineTable[index]
		if byteindex > byteOffset {
			max = index - 1
		} else {
			min = index
		}
	}
	return min + 1
}

// WriteSnippet writes an HTML snippet showing the line(s) covering the byte
// offsets [startByte, endByte) in srcd, linked to the full source file, which
// is assumed to be at path path. contextLines is the number of lines of context
// (on both sides) to include around the given location.
func WriteSnippet(w io.Writer, path string, srcd SrcData, startByte, endByte, contextLines int) {
	if startByte >= len(srcd.Src) {
		return
	}

	beginContext := -1
	startLine := startByte
	for {
		if startLine == 0 {
			break
		}
		if srcd.Src[startLine-1] == '\n' {
			beginContext++
		}
		if beginContext == contextLines {
			break
		}
		startLine--
	}

	lineno := srcd.LineNumber(int32(startLine))

	endLineno := lineno + 1 + (2 * contextLines)
	for i := startByte; i != len(srcd.Src) && i != endByte; i++ {
		if srcd.Src[i] == '\n' {
			endLineno++
		}
	}

	for startLine < len(srcd.Src) && lineno != endLineno {
		endLine := startLine
		for endLine != len(srcd.Src) && srcd.Src[endLine] != '\n' {
			endLine++
		}
		fmt.Fprintf(w, "<a href=\"/%s#%d\"><span class=\"lines\">%5d</span> ", path, lineno, lineno)
		if startLine <= startByte && endByte <= endLine {
			fmt.Fprintf(w, "%s<b>%s</b>%s</a>\n",
				html.EscapeString(string(srcd.Src[startLine:startByte])),
				html.EscapeString(string(srcd.Src[startByte:endByte])),
				html.EscapeString(string(srcd.Src[endByte:endLine])))
		} else {
			fmt.Fprintf(w, "%s</a>\n", html.EscapeString(string(srcd.Src[startLine:endLine])))
		}
		lineno++
		startLine = endLine + 1
	}
}
