package validateops

type Plan[T any] func(T, Sink)

type Sink interface {
	Error(error)

	Field(string) Sink
	Key(any) Sink
	Note(string) Sink

	WantMore() bool
}

type ValidateOper interface {
	ValidateOp(Sink)
}

type KV[K comparable, V any] struct {
	K K
	V V
}

type Error struct {
	Context string
	Err     error
}
