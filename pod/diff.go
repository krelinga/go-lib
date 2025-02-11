package pod

type DiffReporter interface {
	TypeDiff(a, b interface{})
	ValueDiff(a, b interface{})
	Missing(a interface{})
	Extra(b interface{})

	ChildField(name string) DiffReporter
	ChildKey(key interface{}) DiffReporter
}

func Diff[P POD](a, b P, reporter DiffReporter) {
	a.InternalDiff(b, reporter)
}

func Same[P POD](a, b P) bool {
	bdr := boolDiffReporter(false)
	Diff(a, b, &bdr)
	return !bool(bdr)
}

type boolDiffReporter bool

func (bdr *boolDiffReporter) TypeDiff(a, b interface{}) {
	*bdr = true
}

func (bdr *boolDiffReporter) ValueDiff(a, b interface{}) {
	*bdr = true
}

func (bdr *boolDiffReporter) Missing(a interface{}) {
	*bdr = true
}

func (bdr *boolDiffReporter) Extra(b interface{}) {
	*bdr = true
}

func (bdr *boolDiffReporter) ChildField(name string) DiffReporter {
	return bdr
}

func (bdr *boolDiffReporter) ChildKey(key interface{}) DiffReporter {
	return bdr
}