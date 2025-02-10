package base

type Differ interface {
	Diff(other interface{}, reporter DiffReporter)
}

type DiffReporter interface {
	ReportTypeDiff(a, b interface{})
	ReportDiffValues(a, b interface{})
	ChildField(name string) DiffReporter
}

func Diff[D Differ](a, b D, reporter DiffReporter) {
	a.Diff(b, reporter)
}

func Same[D Differ](a, b D) bool {
	reporter := &boolDiffReporter{}
	Diff(a, b, reporter)
	return !reporter.HadDiffs
}

type boolDiffReporter struct {
	HadDiffs bool
}

func (bdr *boolDiffReporter) ReportTypeDiff(a, b interface{}) {
	bdr.HadDiffs = true
}

func (bdr *boolDiffReporter) ReportDiffValues(a, b interface{}) {
	bdr.HadDiffs = true
}

func (bdr *boolDiffReporter) ReportDiffStrings(a, b string) {
	bdr.HadDiffs = true
}

func (bdr *boolDiffReporter) ChildField(name string) DiffReporter {
	return bdr
}
