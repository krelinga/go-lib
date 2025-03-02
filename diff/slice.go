package diff

import (
	"reflect"

	"github.com/krelinga/go-lib/kslice"
)

type Slice[T any] struct {
	Explicit bool
	IgnoreExtra bool
	IgnoreMissing bool
	Entries map[kslice.Index]TypedDiffer[T]
	Methods map[string]AnyDiffer
	EntryDiffer TypedDiffer[T]
}

func (s Slice[T]) typedDiff(lhs, rhs []T) Results {
	return nil // TODO
}

func (s Slice[T]) anyDiff(lhs, rhs any) Results {
	return nil // TODO
}

func (s Slice[T]) Validate() error {
	return nil // TODO
}

func (s Slice[T]) accepts(t reflect.Type) error {
	return nil // TODO
}