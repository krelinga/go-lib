package podmock

import "fmt"

type DiffPair struct {
	A, B interface{}
}

type GotDiff struct {
	Path string

	TypeDiff  *DiffPair
	ValueDiff *DiffPair

	Missing interface{}
	Extra   interface{}
}

type DiffReporter struct {
	GotDiffs []GotDiff
}

func (dr *DiffReporter) TypeDiff(a, b interface{}) {
	dr.GotDiffs = append(dr.GotDiffs, GotDiff{TypeDiff: &DiffPair{A: a, B: b}})
}

func (dr *DiffReporter) ValueDiff(a, b interface{}) {
	dr.GotDiffs = append(dr.GotDiffs, GotDiff{ValueDiff: &DiffPair{A: a, B: b}})
}

func (dr *DiffReporter) Missing(a interface{}) {
	dr.GotDiffs = append(dr.GotDiffs, GotDiff{Missing: a})
}

func (dr *DiffReporter) Extra(b interface{}) {
	dr.GotDiffs = append(dr.GotDiffs, GotDiff{Extra: b})
}

func (dr *DiffReporter) ChildField(name string) *childDiffReporter {
	return &childDiffReporter{parent: dr, path: name}
}

func (dr *DiffReporter) ChildKey(key interface{}) *childDiffReporter {
	return &childDiffReporter{parent: dr, path: fmt.Sprintf("[%v]", key)}
}

type childDiffReporter struct {
	parent *DiffReporter
	path   string
}

func (cdr *childDiffReporter) TypeDiff(a, b interface{}) {
	cdr.parent.GotDiffs = append(cdr.parent.GotDiffs, GotDiff{Path: cdr.path, TypeDiff: &DiffPair{A: a, B: b}})
}

func (cdr *childDiffReporter) ValueDiff(a, b interface{}) {
	cdr.parent.GotDiffs = append(cdr.parent.GotDiffs, GotDiff{Path: cdr.path, ValueDiff: &DiffPair{A: a, B: b}})
}

func (cdr *childDiffReporter) Missing(a interface{}) {
	cdr.parent.GotDiffs = append(cdr.parent.GotDiffs, GotDiff{Path: cdr.path, Missing: a})
}

func (cdr *childDiffReporter) Extra(b interface{}) {
	cdr.parent.GotDiffs = append(cdr.parent.GotDiffs, GotDiff{Path: cdr.path, Extra: b})
}

func (cdr *childDiffReporter) ChildField(name string) *childDiffReporter {
	return &childDiffReporter{parent: cdr.parent, path: fmt.Sprintf("%s.%s", cdr.path, name)}
}

func (cdr *childDiffReporter) ChildKey(key interface{}) *childDiffReporter {
	return &childDiffReporter{parent: cdr.parent, path: fmt.Sprintf("%s.[%v]", cdr.path, key)}
}
