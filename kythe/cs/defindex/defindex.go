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

// Package defindex defines the data structures for definition records, which
// are used to prioritize search results corresponding to definitions.
package defindex

import (
	"fmt"
)

// A Definition is a definition record.
type Definition struct {
	// The full name of the entity, as it may be written in the source code.
	// This is in lower case in order to simplify case insensitive matching.
	FullName string

	// The number of references to this definition. We prioritize the most
	// frequently referenced entities.
	RefCount uint32

	// The start and end bytes for the part of the definition that names the
	// entity.
	StartByte, EndByte uint32
}

// A File is a file record. File records also control which files are added to
// the full-text search index.
type File struct {
	// The name of the file, both in original case and in lower case in
	// order to simplify case insensitive matching.
	Name, NameLower string

	// The list of definition records for this file.
	Defs []Definition
}

// An Index is just a list of files.
type Index []File

// Dump dumps the defindex to stdout.
func (ix Index) Dump() {
	for _, f := range ix {
		for _, d := range f.Defs {
			fmt.Printf("%s @ %s:%d:%d\n", d.FullName, f.Name, d.StartByte, d.EndByte)
		}
	}
}
