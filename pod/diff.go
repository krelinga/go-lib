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
	return false // TODO: Implement
}
