package diffops

func Equals[T DiffOper](lhs T, rhs any) bool {
	var ads anyDiffSink
	lhs.DiffOp(rhs, &ads)
	return !bool(ads)
}

func EqualsPlan[T any](lhs, rhs T, p Plan[T]) bool {
	var ads anyDiffSink
	p(lhs, rhs, &ads)
	return !bool(ads)
}

type anyDiffSink bool

func (ads *anyDiffSink) TypeDiff(any, any) {
	*ads = true
}

func (ads *anyDiffSink) ValueDiff(any, any) {
	*ads = true
}

func (ads *anyDiffSink) Extra(any) {
	*ads = true
}

func (ads *anyDiffSink) Missing(any) {
	*ads = true
}

func (ads *anyDiffSink) Field(string) Sink {
	return ads
}

func (ads *anyDiffSink) Key(any) Sink {
	return ads
}

func (ads *anyDiffSink) WantMore() bool {
	return !bool(*ads)
}
