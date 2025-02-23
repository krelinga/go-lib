package diffops

type Plan[T any] func(lhs, rhs T, s Sink)

type Sink interface {
	TypeDiff(lhs, rhs any)
	ValueDiff(lhs, rhs any)
	Extra(rhs any)
	Missing(rhs any)

	Field(name string) Sink
	Key(key any) Sink

	WantMore() bool
}

type DiffOper interface {
	DiffOp(rhs any, s Sink)
}
