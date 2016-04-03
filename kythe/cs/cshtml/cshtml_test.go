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

package cshtml

import (
	"testing"
)

func TestSrcData(t *testing.T) {
	type tcase struct {
		byteOffset int32
		lineNumber int
	}
	tests := []struct {
		src   string
		cases []tcase
	}{
		{
			"line 1\nline 2\n",
			[]tcase{{0, 1}, {6, 1}, {7, 2}, {13, 2}},
		},
		{
			"line 1\nline 2\nline 3\nline 4\n",
			[]tcase{{0, 1}, {6, 1}, {7, 2}, {13, 2}, {14, 3}, {20, 3}, {21, 4}, {27, 4}},
		},
	}

	for _, test := range tests {
		srcd := MakeSrcData([]byte(test.src))
		for _, c := range test.cases {
			if got := srcd.LineNumber(c.byteOffset); got != c.lineNumber {
				t.Errorf("%v; expected line number of byte offset %d to be %d, got %d", test.src, c.byteOffset, c.lineNumber, got)
			}
		}
	}
}
