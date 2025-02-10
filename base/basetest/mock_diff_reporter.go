package basetest

import "github.com/krelinga/go-lib/base"

type MockDiffReporter struct {
	ReportedTypeDiffs  []ReportedTypeDiff
	ReportedValueDiffs []ReportedValueDiff
}

func (mdr *MockDiffReporter) TypeDiff(a, b interface{}) {
	mdr.ReportedTypeDiffs = append(mdr.ReportedTypeDiffs, ReportedTypeDiff{A: a, B: b})
}

func (mdr *MockDiffReporter) ValueDiff(a, b interface{}) {
	mdr.ReportedValueDiffs = append(mdr.ReportedValueDiffs, ReportedValueDiff{A: a, B: b})
}

func (mdr *MockDiffReporter) ChildField(name string) base.DiffReporter {
	return &mockChildFieldDiffReporter{parent: mdr, path: name}
}

type ReportedTypeDiff struct {
	Path string
	A, B interface{}
}

type ReportedValueDiff struct {
	Path string
	A, B interface{}
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

func (mcd *mockChildFieldDiffReporter) ChildField(name string) base.DiffReporter {
	return &mockChildFieldDiffReporter{parent: mcd.parent, path: mcd.path + "." + name}
}
