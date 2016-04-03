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
	"encoding/binary"
	"errors"
	"io"
	"sort"
)

type key struct {
	name               string
	startByte, endByte uint32
}

type Writer struct {
	bytes.Buffer
	keys []key
}

func (w *Writer) NewName(name string) {
	if len(w.keys) != 0 {
		w.keys[len(w.keys)-1].endByte = uint32(w.Len())
	}
	w.keys = append(w.keys, key{name, uint32(w.Len()), 0})
}

type byName []key

func (b byName) Len() int {
	return len(b)
}
func (b byName) Less(i, j int) bool {
	return b[i].name < b[j].name
}
func (b byName) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func (w *Writer) WriteTo(fw io.Writer) {
	if len(w.keys) != 0 {
		w.keys[len(w.keys)-1].endByte = uint32(w.Len())
	}
	sort.Sort(byName(w.keys))

	strslen := 0
	for _, k := range w.keys {
		strslen += len(k.name)
	}

	buf := make([]byte, 8+12*len(w.keys)+4+strslen)
	binary.LittleEndian.PutUint32(buf, uint32(len(w.keys)))
	binary.LittleEndian.PutUint32(buf[4:], uint32(strslen))

	tab := buf[8:]
	strs := buf[8+12*len(w.keys)+4:]

	tabpos := 0
	strspos := 0
	for _, k := range w.keys {
		binary.LittleEndian.PutUint32(tab, uint32(strspos))
		binary.LittleEndian.PutUint32(tab[4:], k.startByte)
		binary.LittleEndian.PutUint32(tab[8:], k.endByte)
		tabpos += 12
		tab = tab[12:]

		copy(strs, k.name)
		strspos += len(k.name)
		strs = strs[len(k.name):]
	}
	binary.LittleEndian.PutUint32(tab, uint32(strspos))

	fw.Write(buf)
	fw.Write(w.Bytes())
}

type Reader struct {
	io.ReadSeeker
	keyslen uint32
	header  []byte
}

func NewReader(fr io.ReadSeeker) (Reader, error) {
	header := make([]byte, 8)
	n, err := fr.Read(header)
	if err != nil {
		return Reader{}, err
	} else if n != 8 {
		return Reader{}, errors.New("short read")
	}

	keyslen := binary.LittleEndian.Uint32(header)
	strslen := binary.LittleEndian.Uint32(header[4:])
	fullheader := make([]byte, 12*keyslen+4+strslen)
	n, err = fr.Read(fullheader)
	if err != nil {
		return Reader{}, err
	} else if n != len(fullheader) {
		return Reader{}, errors.New("short read")
	}

	return Reader{fr, keyslen, fullheader}, nil
}

func (r *Reader) SeekToName(name string) (int, error) {
	namebytes := []byte(name)
	strsoffset := 12*r.keyslen + 4
	found := false
	var startByte, endByte uint32
	sort.Search(int(r.keyslen), func(i int) bool {
		keyOffset := 12 * i
		keyStartByte := strsoffset + binary.LittleEndian.Uint32(r.header[keyOffset:])
		keyEndByte := strsoffset + binary.LittleEndian.Uint32(r.header[keyOffset+12:])
		c := bytes.Compare(namebytes, r.header[keyStartByte:keyEndByte])
		if c == 0 {
			found = true
			startByte = binary.LittleEndian.Uint32(r.header[keyOffset+4:])
			endByte = binary.LittleEndian.Uint32(r.header[keyOffset+8:])
		}
		return c <= 0
	})
	if !found {
		return -1, nil
	}
	_, err := r.Seek(int64(8+len(r.header)+int(startByte)), 0)
	return int(endByte - startByte), err
}
