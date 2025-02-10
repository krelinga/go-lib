package base

type Differ interface {
	Diff(other interface{}, reporter DiffReporter)
}

type DiffReporter interface {
	TypeDiff(a, b interface{})
	ValueDiff(a, b interface{})

	Missing(a interface{})
	Extra(b interface{})

	ChildField(name string) DiffReporter
	ChildKey(key interface{}) DiffReporter
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

func (bdr *boolDiffReporter) TypeDiff(a, b interface{}) {
	bdr.HadDiffs = true
}

func (bdr *boolDiffReporter) ValueDiff(a, b interface{}) {
	bdr.HadDiffs = true
}

func (bdr *boolDiffReporter) Missing(a interface{}) {
	bdr.HadDiffs = true
}

func (bdr *boolDiffReporter) Extra(b interface{}) {
	bdr.HadDiffs = true
}

func (bdr *boolDiffReporter) ChildField(name string) DiffReporter {
	return bdr
}

func (bdr *boolDiffReporter) ChildKey(key interface{}) DiffReporter {
	return bdr
}

func DiffMap[K comparable, V Differ](a, b map[K]V, reporter DiffReporter) {
	for k, va := range a {
		if vb, ok := b[k]; ok {
			Diff(va, vb, reporter.ChildKey(k))
		} else {
			reporter.ChildKey(k).Missing(va)
		}
	}
	for k, vb := range b {
		if _, ok := a[k]; !ok {
			reporter.ChildKey(k).Extra(vb)
		}
	}
}