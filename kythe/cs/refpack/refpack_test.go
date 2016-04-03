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

package refpack

import (
	"bytes"
	"testing"
)

func TestWriterReader(t *testing.T) {
	cases := []struct {
		name, content string
	}{
		{"file1", "test"},
		{"file3", "foo"},
		{"file2", "some data"},
	}
	var b bytes.Buffer

	var w Writer

	for _, c := range cases {
		w.NewName(c.name)
		w.WriteString(c.content)
	}

	w.WriteTo(&b)

	r, err := NewReader(bytes.NewReader(b.Bytes()))
	if err != nil {
		t.Error(err)
	}

	for _, c := range cases {
		n, err := r.SeekToName(c.name)
		if err != nil {
			t.Error(err)
		}
		if n == -1 {
			t.Errorf("%s: expected to find file", c.name)
		}
		buf := make([]byte, n)
		r.Read(buf)
		if string(buf) != c.content {
			t.Errorf("%s: unexpected content; want %#v, got %#v", c.name, c.content, string(buf))
		}
	}

	for _, name := range []string{"file0", "file2a", "file4"} {
		n, err := r.SeekToName(name)
		if err != nil {
			t.Error(err)
		}
		if n != -1 {
			t.Errorf("%s: expected not found", name)
		}
	}
}
