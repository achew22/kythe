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
	"encoding/gob"
	"fmt"
	"html"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
	"runtime/pprof"
	"sort"
	"strconv"
	"strings"

	"github.com/google/codesearch/index"

	"kythe.io/kythe/go/platform/delimited"
	spb "kythe.io/kythe/proto/storage_proto"

	"kythe.io/kythe/cs/cshtml"
	"kythe.io/kythe/cs/defindex"
	"kythe.io/kythe/cs/refpack"
)

type StrTab struct {
	m    map[string]StrIndex
	Strs []string
}

type StrIndex int32

func (strs *StrTab) idx(str string) StrIndex {
	if strs.m == nil {
		if len(strs.Strs) != 0 {
			panic("appending to deserialized StrTab")
		}
		strs.m = make(map[string]StrIndex)
	}
	idx, ok := strs.m[str]
	if ok {
		return idx
	}
	idx = StrIndex(len(strs.Strs))
	strs.Strs = append(strs.Strs, str)
	strs.m[str] = idx
	return idx
}

func (i StrIndex) get(strs *StrTab) string {
	return strs.Strs[i]
}

type Loc struct {
	Path               StrIndex
	StartByte, EndByte int32
}

type LocNode struct {
	L       Loc
	Name    StrIndex
	VN      StrIndex
	NoUnref bool
}

type RelKind bool

const RelKindDerives RelKind = false
const RelKindOverrides RelKind = true

type Rel struct {
	Kind     RelKind
	Src, Dst StrIndex
}

type TUInfo struct {
	Incs, Refs, Defs, Comps []LocNode
	Rels                    []Rel
	Strs                    StrTab
}

func (info *TUInfo) mergeXrefList(target *[]LocNode, source []LocNode, strs *StrTab) {
	offset := len(*target)
	t := append(*target, source...)
	*target = t
	for i := offset; i != len(t); i++ {
		t[i].L.Path = info.Strs.idx(t[i].L.Path.get(strs))
		t[i].Name = info.Strs.idx(t[i].Name.get(strs))
		t[i].VN = info.Strs.idx(t[i].VN.get(strs))
	}
}

func (info *TUInfo) mergeRelList(target *[]Rel, source []Rel, strs *StrTab) {
	offset := len(*target)
	t := append(*target, source...)
	*target = t
	for i := offset; i != len(t); i++ {
		t[i].Src = info.Strs.idx(t[i].Src.get(strs))
		t[i].Dst = info.Strs.idx(t[i].Dst.get(strs))
	}
}

func (info *TUInfo) merge(other TUInfo) {
	info.mergeXrefList(&info.Incs, other.Incs, &other.Strs)
	info.mergeXrefList(&info.Refs, other.Refs, &other.Strs)
	info.mergeXrefList(&info.Defs, other.Defs, &other.Strs)
	info.mergeXrefList(&info.Comps, other.Comps, &other.Strs)
	info.mergeRelList(&info.Rels, other.Rels, &other.Strs)
}

type byLocNodeAll []LocNode

func (b byLocNodeAll) Len() int {
	return len(b)
}
func (b byLocNodeAll) Less(i, j int) bool {
	if b[i].L.Path < b[j].L.Path {
		return true
	}
	if b[i].L.Path > b[j].L.Path {
		return false
	}

	if b[i].L.StartByte < b[j].L.StartByte {
		return true
	}
	if b[i].L.StartByte > b[j].L.StartByte {
		return false
	}

	if b[i].L.EndByte < b[j].L.EndByte {
		return true
	}
	if b[i].L.EndByte > b[j].L.EndByte {
		return false
	}

	if b[i].Name < b[j].Name {
		return true
	}
	if b[i].Name > b[j].Name {
		return false
	}

	return b[i].VN < b[j].VN
}
func (b byLocNodeAll) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func dedupXrefList(xrefs *[]LocNode) {
	xr := *xrefs
	sort.Sort(byLocNodeAll(xr))
	ahead := 0
	behind := 0
	for ahead != len(xr) {
		entry := xr[ahead]
		xr[behind] = entry
		ahead++
		for ahead != len(xr) && xr[ahead] == entry {
			ahead++
		}
		behind++
	}
	*xrefs = xr[:behind]
}

type byRelAll []Rel

func (b byRelAll) Len() int {
	return len(b)
}
func (b byRelAll) Less(i, j int) bool {
	if !b[i].Kind && b[j].Kind {
		return true
	}
	if b[i].Kind && !b[j].Kind {
		return false
	}

	if b[i].Src < b[j].Src {
		return true
	}
	if b[i].Src > b[j].Src {
		return false
	}

	return b[i].Dst < b[j].Dst
}
func (b byRelAll) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func dedupRelList(xrefs *[]Rel) {
	xr := *xrefs
	sort.Sort(byRelAll(xr))
	ahead := 0
	behind := 0
	for ahead != len(xr) {
		entry := xr[ahead]
		xr[behind] = entry
		ahead++
		for ahead != len(xr) && xr[ahead] == entry {
			ahead++
		}
		behind++
	}
	*xrefs = xr[:behind]
}

func (info *TUInfo) dedup() {
	dedupXrefList(&info.Incs)
	dedupXrefList(&info.Refs)
	dedupXrefList(&info.Defs)
	dedupXrefList(&info.Comps)
	dedupRelList(&info.Rels)
}

func (i *TUInfo) dump() {
	for _, r := range i.Incs {
		fmt.Printf("inc %s:%d:%d : %s\n", r.L.Path.get(&i.Strs), r.L.StartByte, r.L.EndByte, r.VN.get(&i.Strs))
	}
	for _, r := range i.Refs {
		fmt.Printf("ref %s:%d:%d : %#v\n", r.L.Path.get(&i.Strs), r.L.StartByte, r.L.EndByte, r.VN.get(&i.Strs))
	}
	for _, r := range i.Defs {
		fmt.Printf("def %s:%d:%d (%s): %#v\n", r.L.Path.get(&i.Strs), r.L.StartByte, r.L.EndByte, r.Name.get(&i.Strs), r.VN.get(&i.Strs))
	}
	for _, r := range i.Comps {
		fmt.Printf("comp %s:%d:%d (%s): %#v\n", r.L.Path.get(&i.Strs), r.L.StartByte, r.L.EndByte, r.Name.get(&i.Strs), r.VN.get(&i.Strs))
	}
	for _, r := range i.Rels {
		var kind string
		switch r.Kind {
		case RelKindOverrides:
			kind = "overrides"
		case RelKindDerives:
			kind = "derives"
		}
		fmt.Printf("%s %#v -> %#v\n", kind, r.Src.get(&i.Strs), r.Dst.get(&i.Strs))
	}
}

func mkdirAndCreate(p string) (fh *os.File, err error) {
	err = os.MkdirAll(path.Dir(p), 0777)
	if err != nil {
		return nil, err
	}

	return os.Create(p)
}

func writeStaticFile(p, contents string) error {
	dir, file := path.Split(p)

	tmpname := dir + ".tmp" + strconv.Itoa(os.Getpid()) + "." + file
	fh, err := mkdirAndCreate(tmpname)
	if err != nil {
		return err
	}
	_, err = fh.WriteString(contents)
	if err != nil {
		return err
	}

	err = fh.Close()
	if err != nil {
		return err
	}

	return os.Rename(tmpname, p)
}

type tulocnode struct {
	path   string
	anchor string
	vn     *spb.VName
}

type tuinc struct {
	path   string
	anchor string
	target string
}

type tuanchor struct {
	start, end int32
}

type tuindex struct {
	outdir                string
	names                 map[string]string
	anchors               map[string]tuanchor
	nodeBlacklist         map[string]struct{}
	nodeDefIndexBlacklist map[string]bool
	nodeUnrefBlacklist    map[string]bool
	info                  TUInfo
	incs                  []tuinc
	refs, defs, comps     []tulocnode
}

func (ix *tuindex) init() {
	ix.names = make(map[string]string)
	ix.anchors = make(map[string]tuanchor)
	ix.nodeBlacklist = make(map[string]struct{})
	ix.nodeDefIndexBlacklist = make(map[string]bool)
	ix.nodeUnrefBlacklist = make(map[string]bool)
}

func (ix *tuindex) manglePath(path string) string {
	if path != "" && path[0] == '/' {
		path = path[1:]
	}
	return path
}

func vnameStr(vname spb.VName) string {
	return vname.Signature + "\000" + vname.Corpus + "\000" + vname.Root + "\000" + vname.Path + "\000" + vname.Language
}

func (ix *tuindex) addLocNode(path string, xrefs *[]tulocnode, entry *spb.Entry) {
	if strings.HasSuffix(entry.Target.Signature, "#builtin") {
		return
	}
	if path != "" {
		*xrefs = append(*xrefs, tulocnode{path, vnameStr(*entry.Source), entry.Target})
	}
}

func asLanguageName(qname string) string {
	qname = qname[:strings.LastIndex(qname, "#")]
	if strings.ContainsRune(qname, '#') { // operator overloading
		return ""
	}
	qname_parts := strings.Split(qname, ":")
	result := ""
	for i := len(qname_parts); i != 0; i-- {
		p := qname_parts[i-1]
		if len(p) == 0 || (p[0] >= '0' && p[0] <= '9') {
			return ""
		}
		result += "::"
		result += strings.ToLower(p)
	}
	if len(result) == 0 {
		return ""
	}
	return result[2:]
}

func (ix *tuindex) processEntry(entry *spb.Entry) {
	path := ix.manglePath(entry.Source.Path)
	if entry.FactName == "/kythe/text" && path != "" && entry.Source.Signature == "" {
		writeStaticFile(ix.outdir+"/raw/"+path, string(entry.FactValue))
	}
	if entry.FactName == "/kythe/node/kind" {
		kind := string(entry.FactValue)
		if kind == "tapp" || kind == "lookup" {
			ix.nodeBlacklist[vnameStr(*entry.Source)] = struct{}{}
		}
		if kind == "absvar" {
			ix.nodeDefIndexBlacklist[vnameStr(*entry.Source)] = true
		}
		if kind == "variable" {
			ix.nodeUnrefBlacklist[vnameStr(*entry.Source)] = true
		}
	}
	if entry.FactName == "/kythe/subkind" {
		kind := string(entry.FactValue)
		if kind == "field" {
			delete(ix.nodeUnrefBlacklist, vnameStr(*entry.Source))
		}
	}
	if entry.EdgeKind == "/kythe/edge/param" {
		ix.nodeDefIndexBlacklist[vnameStr(*entry.Target)] = true
	}
	if entry.FactName == "/kythe/loc/start" {
		srcname := vnameStr(*entry.Source)
		start, _ := strconv.Atoi(string(entry.FactValue))
		a := ix.anchors[srcname]
		a.start = int32(start + 1)
		ix.anchors[srcname] = a
	}
	if entry.FactName == "/kythe/loc/end" {
		srcname := vnameStr(*entry.Source)
		end, _ := strconv.Atoi(string(entry.FactValue))
		a := ix.anchors[srcname]
		a.end = int32(end + 1)
		ix.anchors[srcname] = a
	}
	if strings.HasPrefix(entry.EdgeKind, "/kythe/edge/extends") {
		ix.info.Rels = append(ix.info.Rels, Rel{
			RelKindDerives,
			ix.info.Strs.idx(vnameStr(*entry.Source)),
			ix.info.Strs.idx(vnameStr(*entry.Target)),
		})
	}
	if strings.HasPrefix(entry.EdgeKind, "/kythe/edge/overrides") {
		ix.info.Rels = append(ix.info.Rels, Rel{
			RelKindOverrides,
			ix.info.Strs.idx(vnameStr(*entry.Source)),
			ix.info.Strs.idx(vnameStr(*entry.Target)),
		})
	}
	if entry.EdgeKind == "/kythe/edge/named" {
		ix.names[vnameStr(*entry.Source)] = asLanguageName(entry.Target.Signature)
	}
	if entry.EdgeKind == "/kythe/edge/ref/includes" {
		ix.incs = append(ix.incs, tuinc{path, vnameStr(*entry.Source), ix.manglePath(entry.Target.Path)})
	}
	if entry.EdgeKind == "/kythe/edge/ref" || entry.EdgeKind == "/kythe/edge/ref/expands" {
		ix.addLocNode(path, &ix.refs, entry)
	}
	if entry.EdgeKind == "/kythe/edge/defines/binding" {
		ix.addLocNode(path, &ix.defs, entry)
	}
	if entry.EdgeKind == "/kythe/edge/completes" || entry.EdgeKind == "/kythe/edge/completes/uniquely" {
		ix.addLocNode(path, &ix.comps, entry)
	}
}

func (ix *tuindex) makeLoc(path, anchor string) (Loc, bool) {
	a := ix.anchors[anchor]
	if a.start == 0 || a.end == 0 {
		return Loc{}, false
	}

	return Loc{ix.info.Strs.idx(path), a.start - 1, a.end - 1}, true
}

func (ix *tuindex) prepareXrefs(xrefs []tulocnode, needname bool) []LocNode {
	var nodes []LocNode
	for _, xref := range xrefs {
		vnstr := vnameStr(*xref.vn)
		if _, ok := ix.nodeBlacklist[vnstr]; ok {
			continue
		}
		loc, ok := ix.makeLoc(xref.path, xref.anchor)
		if !ok {
			continue
		}

		var name StrIndex
		if needname && !ix.nodeDefIndexBlacklist[vnstr] {
			name = ix.info.Strs.idx(ix.names[vnstr])
		} else {
			name = ix.info.Strs.idx("")
		}

		nodes = append(nodes, LocNode{loc, name, ix.info.Strs.idx(vnstr), ix.nodeUnrefBlacklist[vnstr]})
	}
	return nodes
}

func (ix *tuindex) finish() {
	for _, inc := range ix.incs {
		loc, ok := ix.makeLoc(inc.path, inc.anchor)
		if !ok {
			continue
		}
		ix.info.Incs = append(ix.info.Incs, LocNode{loc, ix.info.Strs.idx(""), ix.info.Strs.idx(inc.target), true})
	}
	ix.info.Refs = ix.prepareXrefs(ix.refs, false)
	ix.info.Defs = ix.prepareXrefs(ix.defs, true)
	ix.info.Comps = ix.prepareXrefs(ix.comps, true)
}

func tu(outdir string, in io.Reader) TUInfo {
	var ix tuindex
	ix.outdir = outdir
	ix.init()

	rd := delimited.NewReader(in)
	var entry spb.Entry
	for {
		if err := rd.NextProto(&entry); err == io.EOF {
			break
		} else if err != nil {
			log.Fatalf("Error decoding Entry: %v", err)
		}
		ix.processEntry(&entry)
	}

	ix.finish()

	memprofile := os.Getenv("MEMPROFILE")
	if memprofile != "" {
		f, err := os.Create(memprofile + "-1")
		if err != nil {
			log.Fatal(err)
		}
		pprof.WriteHeapProfile(f)
		f.Close()
	}

	ix.info.dedup()

	if memprofile != "" {
		f, err := os.Create(memprofile + "-2")
		if err != nil {
			log.Fatal(err)
		}
		pprof.WriteHeapProfile(f)
		f.Close()
	}

	return ix.info
}

type xrefinfo struct {
	def                                          Loc
	qname                                        string
	decls, refs, derivers, overrides, overriders []Loc
	nounref                                      bool
}

type xref struct {
	startByte, endByte int32
	info               *xrefinfo
}

type srcinfo struct {
	xrefs []xref
}

func (srci *srcinfo) dump(name string, strs *StrTab) {
	for _, xr := range srci.xrefs {
		fmt.Printf("%s:%d:%d def @ %s:%d:%d ", name, xr.startByte, xr.endByte, xr.info.def.Path.get(strs), xr.info.def.StartByte, xr.info.def.EndByte)
		dumpLocList := func(name string, locs []Loc) {
			if len(locs) == 0 {
				return
			}
			fmt.Printf("%s @", name)
			for _, r := range locs {
				fmt.Printf(" %s:%d:%d", r.Path.get(strs), r.StartByte, r.EndByte)
			}
		}
		dumpLocList("decls", xr.info.decls)
		dumpLocList("refs", xr.info.refs)
		dumpLocList("derivers", xr.info.derivers)
		dumpLocList("overrides", xr.info.overrides)
		dumpLocList("overriders", xr.info.overriders)
	}
}

type corpusindex struct {
	outdir string

	info     TUInfo
	srcinfos map[StrIndex]*srcinfo
	srcpaths []StrIndex
	srcs     map[StrIndex]cshtml.SrcData
}

func (ix *corpusindex) dump() {
	for _, path := range ix.srcpaths {
		srci := ix.srcinfos[path]
		srci.dump(path.get(&ix.info.Strs), &ix.info.Strs)
	}
}

func (ix *corpusindex) init() {
	ix.srcinfos = make(map[StrIndex]*srcinfo)
	ix.srcs = make(map[StrIndex]cshtml.SrcData)
}

func (ix *corpusindex) srcinfoFor(name StrIndex) *srcinfo {
	if l, ok := ix.srcinfos[name]; ok {
		return l
	}
	ix.srcpaths = append(ix.srcpaths, name)
	newl := new(srcinfo)
	ix.srcinfos[name] = newl
	return newl
}

type byPath struct {
	names []StrIndex
	strs  *StrTab
}

func (b byPath) Len() int {
	return len(b.names)
}
func (b byPath) Less(i, j int) bool {
	return b.names[i].get(b.strs) < b.names[j].get(b.strs)
}
func (b byPath) Swap(i, j int) {
	b.names[i], b.names[j] = b.names[j], b.names[i]
}

func (ix *corpusindex) finish(dump bool) {
	for _, inc := range ix.info.Incs {
		ds := ix.srcinfoFor(inc.L.Path)
		ds.xrefs = append(ds.xrefs, xref{startByte: inc.L.StartByte, endByte: inc.L.EndByte, info: &xrefinfo{def: Loc{inc.VN, 0, 0}, nounref: true}})
	}
	ix.info.Incs = nil

	type xreflist struct {
		qname                           StrIndex
		refs, defs                      []Loc
		derivers, overrides, overriders *[]Loc
		toUpdate                        []*[]Loc
		nounref                         bool
	}

	xrefmap := make(map[StrIndex]*xreflist)
	xrefsFor := func(loc LocNode) *xreflist {
		if l, ok := xrefmap[loc.VN]; ok {
			return l
		}
		newl := new(xreflist)
		newl.qname = loc.Name
		xrefmap[loc.VN] = newl
		return newl
	}
	dumpXrefmap := func(when string) {
		if !dump {
			return
		}
		fmt.Printf("xrefmap (%s):\n", when)
		for sn, xrl := range xrefmap {
			fmt.Printf("%#v (%s): r(", sn.get(&ix.info.Strs), xrl.qname.get(&ix.info.Strs))
			for _, r := range xrl.refs {
				fmt.Printf(" %s:%d:%d", r.Path.get(&ix.info.Strs), r.StartByte, r.EndByte)
			}
			fmt.Printf(" ) d(")
			for _, r := range xrl.defs {
				fmt.Printf(" %s:%d:%d", r.Path.get(&ix.info.Strs), r.StartByte, r.EndByte)
			}
			fmt.Printf(" )\n")
		}
		fmt.Println()
	}

	deflocs := make(map[Loc]struct{})

	for _, def := range ix.info.Comps {
		deflocs[def.L] = struct{}{}
	}

	for _, def := range ix.info.Defs {
		deflocs[def.L] = struct{}{}
		xr := xrefsFor(def)
		xr.defs = append(xr.defs, def.L)
		if def.NoUnref {
			xr.nounref = true
		}
	}
	ix.info.Defs = nil

	for _, ref := range ix.info.Refs {
		if _, ok := deflocs[ref.L]; ok {
			continue
		}
		xr, ok := xrefmap[ref.VN]
		if !ok {
			continue
		}
		xr.refs = append(xr.refs, ref.L)
	}
	ix.info.Refs = nil

	for _, rel := range ix.info.Rels {
		src, ok := xrefmap[rel.Src]
		if !ok {
			continue
		}
		dst, ok := xrefmap[rel.Dst]
		if !ok {
			continue
		}

		switch rel.Kind {
		case RelKindOverrides:
			if src.overrides == nil {
				src.overrides = new([]Loc)
			}
			dst.toUpdate = append(dst.toUpdate, src.overrides)
			if dst.overriders == nil {
				dst.overriders = new([]Loc)
			}
			src.toUpdate = append(src.toUpdate, dst.overriders)

		case RelKindDerives:
			if dst.overriders == nil {
				dst.derivers = new([]Loc)
			}
			src.toUpdate = append(src.toUpdate, dst.derivers)
		}
	}

	dumpXrefmap("after defs/refs")

	type compXreflist struct {
		qname                           StrIndex
		refs, defs                      []Loc
		derivers, overrides, overriders []*[]Loc
		nounref                         bool
	}

	compmap := make(map[Loc]*compXreflist)
	mergeComp := func(loc Loc, name StrIndex, xr *xreflist) {
		comps, ok := compmap[loc]
		if !ok {
			comps = new(compXreflist)
			comps.qname = name
			compmap[loc] = comps
		}
		comps.refs = append(comps.refs, xr.refs...)
		comps.defs = append(comps.defs, xr.defs...)
		if xr.derivers != nil {
			comps.derivers = append(comps.derivers, xr.derivers)
		}
		if xr.overrides != nil {
			comps.overrides = append(comps.overrides, xr.overrides)
		}
		if xr.overriders != nil {
			comps.overriders = append(comps.overriders, xr.overriders)
		}
		for _, l := range xr.toUpdate {
			*l = append(*l, loc)
		}
		xr.toUpdate = nil
		if xr.nounref {
			comps.nounref = true
		}
	}
	dumpCompmap := func(when string) {
		if !dump {
			return
		}
		fmt.Printf("compmap (%s):\n", when)
		for loc, xrl := range compmap {
			fmt.Printf("%s:%d:%d (%s): r(", loc.Path.get(&ix.info.Strs), loc.StartByte, loc.EndByte, xrl.qname.get(&ix.info.Strs))
			for _, r := range xrl.refs {
				fmt.Printf(" %s:%d:%d", r.Path.get(&ix.info.Strs), r.StartByte, r.EndByte)
			}
			fmt.Printf(" ) d(")
			for _, r := range xrl.defs {
				fmt.Printf(" %s:%d:%d", r.Path.get(&ix.info.Strs), r.StartByte, r.EndByte)
			}
			fmt.Printf(" )\n")
		}
		fmt.Println()
	}

	for _, comp := range ix.info.Comps {
		xr, ok := xrefmap[comp.VN]
		if !ok {
			continue
		}
		mergeComp(comp.L, comp.Name, xr)
		delete(xrefmap, comp.VN)
	}
	ix.info.Comps = nil

	dumpCompmap("after comps")
	dumpXrefmap("after comps")

	for _, xr := range xrefmap {
		if len(xr.defs) == 0 {
			continue
		}
		def := xr.defs[0]
		xr.defs = xr.defs[1:]
		mergeComp(def, xr.qname, xr)
	}
	xrefmap = nil

	dumpCompmap("after xrefs")

	for def, refs := range compmap {
		ds := ix.srcinfoFor(def.Path)
		xri := &xrefinfo{def: def, qname: refs.qname.get(&ix.info.Strs), decls: refs.defs, refs: refs.refs, nounref: refs.nounref}

		for _, rel := range [...]struct {
			target *[]Loc
			src    []*[]Loc
		}{{&xri.derivers, refs.derivers},
			{&xri.overrides, refs.overrides},
			{&xri.overriders, refs.overriders}} {
			for _, s := range rel.src {
				*rel.target = append(*rel.target, *s...)
			}
		}

		ds.xrefs = append(ds.xrefs, xref{startByte: def.StartByte, endByte: def.EndByte, info: xri})

		for _, r := range append(refs.defs, refs.refs...) {
			s := ix.srcinfoFor(r.Path)
			s.xrefs = append(s.xrefs, xref{startByte: r.StartByte, endByte: r.EndByte, info: xri})
		}
	}
}

func (ix *corpusindex) getSrc(path StrIndex) cshtml.SrcData {
	srcd, ok := ix.srcs[path]
	if ok {
		return srcd
	}

	src, err := ioutil.ReadFile(ix.outdir + "/raw/" + path.get(&ix.info.Strs))
	if err != nil {
		panic(err)
	}

	srcd = cshtml.MakeSrcData(src)
	ix.srcs[path] = srcd
	return srcd
}

type byFileThenStartByte struct {
	locs []Loc
	strs *StrTab
}

func (b byFileThenStartByte) Len() int {
	return len(b.locs)
}
func (b byFileThenStartByte) Less(i, j int) bool {
	p1 := b.locs[i].Path.get(b.strs)
	p2 := b.locs[j].Path.get(b.strs)
	if p1 < p2 {
		return true
	}
	if p1 != p2 {
		return false
	}
	return b.locs[i].StartByte < b.locs[j].StartByte
}
func (b byFileThenStartByte) Swap(i, j int) {
	b.locs[i], b.locs[j] = b.locs[j], b.locs[i]
}

func (ix *corpusindex) locLineNumber(l Loc) int {
	srcd := ix.getSrc(l.Path)
	return srcd.LineNumber(l.StartByte)
}

func (ix *corpusindex) writeRefSnippets(w io.Writer, refs []Loc) {
	sort.Sort(byFileThenStartByte{refs, &ix.info.Strs})
	ahead := 0
	behind := 0
	for ahead != len(refs) {
		entry := refs[ahead]
		refs[behind] = entry
		ahead++
		for ahead != len(refs) && refs[ahead] == entry {
			ahead++
		}
		behind++
	}
	refs = refs[:behind]
	srcpath := ""
	for _, ref := range refs {
		newsrcpath := ref.Path.get(&ix.info.Strs)
		if newsrcpath != srcpath {
			if srcpath != "" {
				io.WriteString(w, "</pre>")
			}
			io.WriteString(w, newsrcpath+"<pre>")
			srcpath = newsrcpath
		}
		cshtml.WriteSnippet(w, newsrcpath, ix.getSrc(ref.Path), int(ref.StartByte), int(ref.EndByte), 0)
	}

	if srcpath == "" {
		io.WriteString(w, "No references")
	} else {
		io.WriteString(w, "</pre>")
	}
}

func (ix *corpusindex) writeRefPage(w io.Writer, xri *xrefinfo) {
	cshtml.WriteHeader(w, "Cross References", "")
	if len(xri.derivers) != 0 {
		io.WriteString(w, "<h1>Derived Classes</h1>")
		ix.writeRefSnippets(w, xri.derivers)
	}
	if len(xri.overrides) != 0 {
		io.WriteString(w, "<h1>Overrides</h1>")
		ix.writeRefSnippets(w, xri.overrides)
	}
	if len(xri.overriders) != 0 {
		io.WriteString(w, "<h1>Overridden By</h1>")
		ix.writeRefSnippets(w, xri.overriders)
	}
	if len(xri.decls) != 0 {
		io.WriteString(w, "<h1>Declarations</h1>")
		ix.writeRefSnippets(w, xri.decls)
	}
	if len(xri.refs) != 0 {
		io.WriteString(w, "<h1>References</h1>")
		ix.writeRefSnippets(w, xri.refs)
	}
	cshtml.WriteFooter(w)
}

func (ix *corpusindex) writeMultipleDefPage(w io.Writer, defs []Loc) {
	cshtml.WriteHeader(w, "Multiple Definitions", "")
	io.WriteString(w, "<h1>Definitions</h1>")
	ix.writeRefSnippets(w, defs)
	cshtml.WriteFooter(w)
}

type byStartByte []xref

func (b byStartByte) Len() int {
	return len(b)
}
func (b byStartByte) Less(i, j int) bool {
	return b[i].startByte < b[j].startByte
}
func (b byStartByte) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func filteredString(str []byte) string {
	result := make([]byte, 0, len(str))
	for _, b := range str {
		if (b >= '0' && b <= '9') || (b >= 'A' && b <= 'Z') || (b >= 'a' && b <= 'z') || b == '_' {
			result = append(result, b)
		}
	}
	return string(result)
}

func (ix *corpusindex) makeRefName(refNumbers map[string]int, prefix rune, id []byte) string {
	refname := string(prefix) + filteredString(id)
	num := refNumbers[refname]
	refNumbers[refname] = num + 1
	if num != 0 {
		refname += "-" + strconv.Itoa(num)
	}
	return refname
}

func (ix *corpusindex) writeAnnotatedSourcePage(w io.Writer, refw *refpack.Writer, fpath StrIndex, srci *srcinfo) {
	pathstr := fpath.get(&ix.info.Strs)
	_, name := path.Split(pathstr)

	srcd := ix.getSrc(fpath)
	src := srcd.Src
	linecount := len(srcd.LineTable)

	cshtml.WriteHeader(w, pathstr, "")

	io.WriteString(w, "<table><tr><td style=\"vertical-align: top;\"><pre class=\"lines\">")
	for i := 1; i <= 10; i++ {
		fmt.Fprintf(w, "<a id=\"%d\" />", i)
	}
	for i := 1; i <= linecount; i++ {
		fmt.Fprintf(w, "<a id=\"%d\" /><a href=\"#%d\">%5d</a>\n", i+10, i, i)
	}
	io.WriteString(w, "</pre></td><td style=\"vertical-align: top;\">")

	sort.Sort(byStartByte(srci.xrefs))

	type xrefGroup struct {
		startByte, endByte int32
		uniqueDef          Loc
		infos              []*xrefinfo
	}

	var xrgroups []xrefGroup
	if len(srci.xrefs) != 0 {
		xrgroups = []xrefGroup{{srci.xrefs[0].startByte, srci.xrefs[0].endByte, srci.xrefs[0].info.def, nil}}
		for _, r := range srci.xrefs {
			lastg := &xrgroups[len(xrgroups)-1]
			if lastg.startByte == r.startByte && lastg.endByte == r.endByte {
				if lastg.uniqueDef != r.info.def {
					lastg.uniqueDef = Loc{}
				}
				lastg.infos = append(lastg.infos, r.info)
			} else if r.startByte >= lastg.endByte {
				xrgroups = append(xrgroups, xrefGroup{r.startByte, r.endByte, r.info.def, []*xrefinfo{r.info}})
			}
		}
	}

	refcounts := make(map[Loc]int)
	for _, g := range xrgroups {
		if g.uniqueDef != (Loc{}) {
			refcounts[g.uniqueDef]++
		}
	}

	refgroupCount := 0
	refgroups := make(map[Loc]int)
	for k, v := range refcounts {
		if v >= 2 {
			refgroups[k] = refgroupCount
			refgroupCount++
		}
	}

	io.WriteString(w, "<pre class=\"code\">")
	lastEnd := 0
	refNumbers := make(map[string]int)

	for _, g := range xrgroups {
		io.WriteString(w, html.EscapeString(string(src[lastEnd:g.startByte])))

		var info *xrefinfo
		if len(g.infos) == 1 {
			info = g.infos[0]
		}
		close := ""
		if info != nil && !info.nounref && len(info.refs) == 0 && len(info.overrides) == 0 {
			io.WriteString(w, "<span class=\"unref\">")
			close = "</span>"
		}
		refgroup, hasRefgroup := refgroups[g.uniqueDef]
		if info != nil && info.def == (Loc{fpath, g.startByte, g.endByte}) {
			if len(info.derivers) != 0 || len(info.overrides) != 0 || len(info.overriders) != 0 ||
				len(info.decls) != 0 || len(info.refs) != 0 {
				io.WriteString(w, "<a class=\"")
				if hasRefgroup {
					io.WriteString(w, "r"+strconv.Itoa(refgroup)+" r ")
				}
				refname := ix.makeRefName(refNumbers, 'r', src[g.startByte:g.endByte])
				refw.NewName(refname)
				ix.writeRefPage(refw, info)
				io.WriteString(w, "def\" href=\""+html.EscapeString(name)+"/"+refname+"\">")
				close = "</a>" + close
			}
		} else if g.uniqueDef != (Loc{}) {
			target := "/" + html.EscapeString(g.uniqueDef.Path.get(&ix.info.Strs))
			target += "#" + strconv.Itoa(ix.locLineNumber(g.uniqueDef))
			io.WriteString(w, "<a")
			if hasRefgroup {
				io.WriteString(w, " class=\"r"+strconv.Itoa(refgroup)+" r\"")
			}
			io.WriteString(w, " href=\""+target+"\">")
			close = "</a>" + close
		} else {
			refname := ix.makeRefName(refNumbers, 'd', src[g.startByte:g.endByte])
			refw.NewName(refname)
			defs := make([]Loc, len(g.infos))
			for i := range g.infos {
				defs[i] = g.infos[i].def
			}
			ix.writeMultipleDefPage(refw, defs)
			fmt.Fprintf(w, "<a class=\"mref\" href=\"%s/%s\">", html.EscapeString(name), refname)
			close = "</a>" + close
		}

		io.WriteString(w, html.EscapeString(string(src[g.startByte:g.endByte])))
		io.WriteString(w, close)
		lastEnd = int(g.endByte)
	}
	io.WriteString(w, html.EscapeString(string(src[lastEnd:])))
	io.WriteString(w, "</pre></td></tr></table>")
	cshtml.WriteFooter(w)
}

func (ix *corpusindex) writeStatics() {
	err := writeStaticFile(ix.outdir+"/res/style.css", `
body {
	font-family: sans-serif;
}

a {
	color: inherit;
	text-decoration: none;
}

.lines {
	color: #808080;
}

.code a {
	text-decoration: underline;
}

.code a.def {
	font-weight: bold;
	text-decoration: none;
}

.code a.mref {
	text-decoration: none;
}

.unref {
	border-bottom: 1px dashed red;
}

.mref {
	border-bottom: 1px dashed black;
}
`)
	if err != nil {
		panic(err)
	}

	err = writeStaticFile(ix.outdir+"/res/script.js", `
var lineNodeStart = 0;
var lineNodeAlignment = 0;
var curHighlightedLine = 0;

function highlightLine() {
	if (location.hash == "")
		return;

	var line = parseInt(location.hash.substring(1));

	var lines = document.getElementsByClassName("lines")[0];
	var kids = lines.childNodes;
	if (lineNodeAlignment == 0) {
		var i = 0;
		while (i != kids.length) {
			if (kids[i].tagName == "A" && kids[i].hash == "#1") {
				lineNodeStart = i;
				break;
			}
			i++;
		}
		while (i != kids.length) {
			if (kids[i].tagName == "A" && kids[i].hash == "#2") {
				lineNodeAlignment = i - lineNodeStart;
				break;
			}
			i++;
		}
	}

	if (curHighlightedLine != 0) {
		kids[lineNodeStart + lineNodeAlignment * (curHighlightedLine - 1)].style.backgroundColor = "";
	}
	kids[lineNodeStart + lineNodeAlignment * (line - 1)].style.backgroundColor = "#e0e0e0";
	curHighlightedLine = line;
}

function setRefgroupColor(elem, color) {
	var index = elem.className.indexOf(' ');
	var refgroupClass;
	if (index == -1) {
		refgroupClass = elem.className;
	} else {
		refgroupClass = elem.className.substring(0, index);
	}
	var refgroup = document.getElementsByClassName(refgroupClass);
	for (var i = 0; i != refgroup.length; i++) {
		refgroup[i].style.backgroundColor = color;
	}
}

function highlightRefgroup() {
	setRefgroupColor(this, "yellow");
}

function unhighlightRefgroup() {
	setRefgroupColor(this, "");
}

function init() {
	highlightLine();
	window.onhashchange = highlightLine;

	var refgroups = document.getElementsByClassName("r");
	for (var j = 0; j != refgroups.length; j++) {
		refgroups[j].onmouseover = highlightRefgroup;
		refgroups[j].onmouseout = unhighlightRefgroup;
	}
}
`)
	if err != nil {
		panic(err)
	}
}

func (ix *corpusindex) makeDefIndex() defindex.Index {
	var didx defindex.Index
	for _, path := range ix.srcpaths {
		srci := ix.srcinfos[path]
		pathstr := path.get(&ix.info.Strs)
		var defs []defindex.Definition
		for _, xr := range srci.xrefs {
			if xr.info.qname != "" && xr.info.def == (Loc{path, xr.startByte, xr.endByte}) {
				defs = append(defs, defindex.Definition{xr.info.qname, uint32(len(xr.info.refs)), uint32(xr.startByte), uint32(xr.endByte)})
			}
		}
		didx = append(didx, defindex.File{pathstr, strings.ToLower(pathstr), defs})
	}
	return didx
}

func (ix *corpusindex) addRepo(repopath, corpuspath string) error {
	lsTreeCmd := exec.Command("git", "ls-tree", "-r", "HEAD:")
	lsTreeCmd.Dir = repopath
	stdout, err := lsTreeCmd.StdoutPipe()
	if err != nil {
		return err
	}
	if err := lsTreeCmd.Start(); err != nil {
		return err
	}
	lsTreeScan := bufio.NewScanner(stdout)
	for lsTreeScan.Scan() {
		fields := strings.Fields(lsTreeScan.Text())
		hash := fields[2]
		path := corpuspath + fields[3]

		pathidx := ix.info.Strs.idx(path)
		srci := ix.srcinfoFor(pathidx)

		if len(srci.xrefs) == 0 {
			f, err := mkdirAndCreate(ix.outdir + "/raw/" + path)
			if err != nil {
				return err
			}
			catCmd := exec.Command("git", "cat-file", "-p", hash)
			catCmd.Dir = repopath
			catCmd.Stdout = f
			if err := catCmd.Run(); err != nil {
				return err
			}
			f.Close()
		}
	}
	return lsTreeCmd.Wait()
}

func (ix *corpusindex) write() {
	os.RemoveAll(ix.outdir + "/defs")
	os.RemoveAll(ix.outdir + "/html")
	os.RemoveAll(ix.outdir + "/index")
	os.RemoveAll(ix.outdir + "/res")
	os.RemoveAll(ix.outdir + "/xrefs")

	ix.writeStatics()

	sort.Sort(byPath{ix.srcpaths, &ix.info.Strs})

	for _, path := range ix.srcpaths {
		srci := ix.srcinfos[path]
		pathstr := path.get(&ix.info.Strs)
		log.Printf("writing %s\n", pathstr)
		f, err := mkdirAndCreate(ix.outdir + "/html/" + pathstr)
		if err != nil {
			panic(err)
		}
		var rw refpack.Writer
		ix.writeAnnotatedSourcePage(f, &rw, path, srci)
		f.Close()

		rf, err := mkdirAndCreate(ix.outdir + "/xrefs/" + pathstr)
		if err != nil {
			panic(err)
		}
		rw.WriteTo(rf)
		rf.Close()
	}

	idx := index.Create(ix.outdir + "/index")
	for _, path := range ix.srcpaths {
		pathstr := path.get(&ix.info.Strs)
		idx.Add(pathstr, bytes.NewReader(ix.getSrc(path).Src))
	}
	idx.Flush()

	didx := ix.makeDefIndex()
	didxfh, err := mkdirAndCreate(ix.outdir + "/defs")
	if err != nil {
		panic(err)
	}
	enc := gob.NewEncoder(didxfh)
	err = enc.Encode(didx)
	if err != nil {
		panic(err)
	}

	err = didxfh.Close()
	if err != nil {
		panic(err)
	}
}

func corpus(outdir, indir string, dump bool) corpusindex {
	var ix corpusindex
	ix.outdir = outdir
	ix.init()

	files, err := ioutil.ReadDir(indir)
	if err != nil {
		panic(err)
	}

	for i, f := range files {
		log.Printf("reading %s\n", f.Name())

		fh, err := os.Open(indir + "/" + f.Name())
		if err != nil {
			panic(err)
		}

		var tuinfo TUInfo
		dec := gob.NewDecoder(fh)
		err = dec.Decode(&tuinfo)
		if err != nil {
			panic(err)
		}
		ix.info.merge(tuinfo)
		fh.Close()
		if i%64 == 63 {
			ix.info.dedup()
		}
	}

	ix.info.dedup()
	log.Println("finishing")
	ix.finish(dump)
	return ix
}

func tumerge(tus []string) {
	var result TUInfo

	for _, tu := range tus {
		fh, err := os.Open(tu)
		if err != nil {
			panic(err)
		}

		var tuinfo TUInfo
		dec := gob.NewDecoder(fh)
		err = dec.Decode(&tuinfo)
		if err != nil {
			panic(err)
		}
		result.merge(tuinfo)
		fh.Close()
	}

	result.dedup()

	enc := gob.NewEncoder(os.Stdout)
	err := enc.Encode(result)
	if err != nil {
		panic(err)
	}
}

func main() {
	cpuprofile := os.Getenv("CPUPROFILE")
	if cpuprofile != "" {
		f, err := os.Create(cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	if len(os.Args) < 2 {
		fmt.Println("commands: tu tumerge dumptu corpus dumpcorpus")
		return
	}

	switch os.Args[1] {
	case "tu":
		nti := tu(os.Args[2], os.Stdin)

		enc := gob.NewEncoder(os.Stdout)
		err := enc.Encode(nti)
		if err != nil {
			panic(err)
		}

		memprofile := os.Getenv("MEMPROFILE")
		if memprofile != "" {
			f, err := os.Create(memprofile + "-3")
			if err != nil {
				log.Fatal(err)
			}
			pprof.WriteHeapProfile(f)
			f.Close()
		}

	case "tumerge":
		tumerge(os.Args[2:])

	case "dumptu":
		ntis := tu(os.Args[2], os.Stdin)
		ntis.dump()

	case "corpus":
		ix := corpus(os.Args[2], os.Args[3], false)
		for _, repo := range os.Args[4:] {
			s := strings.SplitN(repo, "=", 2)
			if len(s) < 2 {
				log.Fatal("expected '=' in repo arg")
			}
			err := ix.addRepo(s[0], s[1])
			if err != nil {
				log.Fatal(err)
			}
		}
		ix.write()

	case "dumpcorpus":
		ix := corpus(os.Args[2], os.Args[3], true)
		ix.dump()
		ix.makeDefIndex().Dump()

	default:
		fmt.Println("unknown command: " + os.Args[1])
		os.Exit(1)
	}
}
