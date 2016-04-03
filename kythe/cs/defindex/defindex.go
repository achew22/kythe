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

package defindex

import (
	"fmt"
)

type Definition struct {
	FullName           string
	RefCount           uint32
	StartByte, EndByte uint32
}

type File struct {
	Name, NameLower string
	Defs            []Definition
}

type Index []File

func (ix Index) Dump() {
	for _, f := range ix {
		for _, d := range f.Defs {
			fmt.Printf("%s @ %s:%d:%d\n", d.FullName, f.Name, d.StartByte, d.EndByte)
		}
	}
}
