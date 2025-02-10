package basetest

import "github.com/krelinga/go-lib/base"

type MockDiffReporter struct {
	ReportedTypeDiffs   []ReportedTypeDiff
	ReportedValueDiffs  []ReportedValueDiff
	ReportedStringDiffs []ReportedStringDiff
}

func (mdr *MockDiffReporter) ReportTypeDiff(a, b interface{}) {
	mdr.ReportedTypeDiffs = append(mdr.ReportedTypeDiffs, ReportedTypeDiff{A: a, B: b})
}

func (mdr *MockDiffReporter) ReportDiffValues(a, b interface{}) {
	mdr.ReportedValueDiffs = append(mdr.ReportedValueDiffs, ReportedValueDiff{A: a, B: b})
}

func (mdr *MockDiffReporter) ReportDiffStrings(a, b string) {
	mdr.ReportedStringDiffs = append(mdr.ReportedStringDiffs, ReportedStringDiff{A: a, B: b})
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

type ReportedStringDiff struct {
	Path string
	A, B string
}

type mockChildFieldDiffReporter struct {
	parent *MockDiffReporter
	path   string
}

func (mcd *mockChildFieldDiffReporter) ReportTypeDiff(a, b interface{}) {
	mcd.parent.ReportedTypeDiffs = append(mcd.parent.ReportedTypeDiffs, ReportedTypeDiff{Path: mcd.path, A: a, B: b})
}

func (mcd *mockChildFieldDiffReporter) ReportDiffValues(a, b interface{}) {
	mcd.parent.ReportedValueDiffs = append(mcd.parent.ReportedValueDiffs, ReportedValueDiff{Path: mcd.path, A: a, B: b})
}

func (mcd *mockChildFieldDiffReporter) ReportDiffStrings(a, b string) {
	mcd.parent.ReportedStringDiffs = append(mcd.parent.ReportedStringDiffs, ReportedStringDiff{Path: mcd.path, A: a, B: b})
}

func (mcd *mockChildFieldDiffReporter) ChildField(name string) base.DiffReporter {
	return &mockChildFieldDiffReporter{parent: mcd.parent, path: mcd.path + "." + name}
}
