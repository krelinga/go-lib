package basetest

import (
	"fmt"

	"github.com/krelinga/go-lib/base"
)

type MockDiffReporter struct {
	ReportedTypeDiffs  []ReportedTypeDiff
	ReportedValueDiffs []ReportedValueDiff
	ReportedMissing    []ReprtedMissing
	ReportedExtra      []ReportedExtra
}

func (mdr *MockDiffReporter) TypeDiff(a, b interface{}) {
	mdr.ReportedTypeDiffs = append(mdr.ReportedTypeDiffs, ReportedTypeDiff{A: a, B: b})
}

func (mdr *MockDiffReporter) ValueDiff(a, b interface{}) {
	mdr.ReportedValueDiffs = append(mdr.ReportedValueDiffs, ReportedValueDiff{A: a, B: b})
}

func (mdr *MockDiffReporter) Missing(a interface{}) {
	mdr.ReportedMissing = append(mdr.ReportedMissing, ReprtedMissing{A: a})
}

func (mdr *MockDiffReporter) Extra(b interface{}) {
	mdr.ReportedExtra = append(mdr.ReportedExtra, ReportedExtra{B: b})
}

func (mdr *MockDiffReporter) ChildField(name string) base.DiffReporter {
	return &mockChildFieldDiffReporter{parent: mdr, path: name}
}

func (mdr *MockDiffReporter) ChildKey(key interface{}) base.DiffReporter {
	return &mockChildFieldDiffReporter{parent: mdr, path: fmt.Sprintf("[%v]", key)}
}

type ReportedTypeDiff struct {
	Path string
	A, B interface{}
}

type ReportedValueDiff struct {
	Path string
	A, B interface{}
}

type ReprtedMissing struct {
	Path string
	A    interface{}
}

type ReportedExtra struct {
	Path string
	B    interface{}
}

type mockChildFieldDiffReporter struct {
	parent *MockDiffReporter
	path   string
}

func (mcd *mockChildFieldDiffReporter) TypeDiff(a, b interface{}) {
	mcd.parent.ReportedTypeDiffs = append(mcd.parent.ReportedTypeDiffs, ReportedTypeDiff{Path: mcd.path, A: a, B: b})
}

func (mcd *mockChildFieldDiffReporter) ValueDiff(a, b interface{}) {
	mcd.parent.ReportedValueDiffs = append(mcd.parent.ReportedValueDiffs, ReportedValueDiff{Path: mcd.path, A: a, B: b})
}

func (mcd *mockChildFieldDiffReporter) Missing(a interface{}) {
	mcd.parent.ReportedMissing = append(mcd.parent.ReportedMissing, ReprtedMissing{Path: mcd.path, A: a})
}

func (mcd *mockChildFieldDiffReporter) Extra(b interface{}) {
	mcd.parent.ReportedExtra = append(mcd.parent.ReportedExtra, ReportedExtra{Path: mcd.path, B: b})
}

func (mcd *mockChildFieldDiffReporter) ChildField(name string) base.DiffReporter {
	return &mockChildFieldDiffReporter{parent: mcd.parent, path: mcd.path + "." + name}
}

func (mcd *mockChildFieldDiffReporter) ChildKey(key interface{}) base.DiffReporter {
	return &mockChildFieldDiffReporter{parent: mcd.parent, path: mcd.path + fmt.Sprintf("[%v]", key)}
}
