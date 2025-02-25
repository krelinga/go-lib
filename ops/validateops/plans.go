package validateops

import (
	"errors"
	"reflect"
)

func ByMethod[T ValidateOper]() Plan[T] {
	return func(op T, sink Sink) {
		if !sink.WantMore() {
			return
		}

		val := reflect.ValueOf(op)
		if !val.IsValid() {
			return
		}
		if val.Kind() == reflect.Ptr && val.IsNil() {
			return
		}

		op.ValidateOp(sink)
	}
}

func MapOf[K comparable, V any](in map[K]V, p Plan[KV[K, V]]) Plan[map[K]V] {
	return func(in map[K]V, sink Sink) {
		for k, v := range in {
			if !sink.WantMore() {
				break
			}
			p(KV[K, V]{k, v}, sink.Key(k))
		}
	}
}

func SliceOf[V any](in []V, p Plan[KV[int, V]]) Plan[[]V] {
	return func(in []V, sink Sink) {
		for i, v := range in {
			if !sink.WantMore() {
				break
			}
			p(KV[int, V]{i, v}, sink.Key(i))
		}
	}
}

func AllOf[T any](p ...Plan[T]) Plan[T] {
	return func(in T, sink Sink) {
		for _, pp := range p {
			if !sink.WantMore() {
				break
			}
			pp(in, sink)
		}
	}
}

func Keys[K comparable, V any](p Plan[K]) Plan[KV[K, V]] {
	return func(in KV[K, V], sink Sink) {
		if !sink.WantMore() {
			return
		}
		// TODO: what sink should be passed here?
		p(in.K, sink.Note("key"))
	}
}

func Values[K comparable, V any](p Plan[V]) Plan[KV[K, V]] {
	return func(in KV[K, V], sink Sink) {
		if !sink.WantMore() {
			return
		}
		// TODO: what sink should be passed here?
		p(in.V, sink.Note("value"))
	}
}

var ErrWantNonZero = errors.New("want non-zero value")

func NonZero[T comparable]() Plan[T] {
	return func(in T, sink Sink) {
		if !sink.WantMore() {
			return
		}
		var zero T
		if in == zero {
			sink.Error(ErrWantNonZero)
		}
	}
}

var ErrWantZero = errors.New("want zero value")

func Zero[T comparable]() Plan[T] {
	return func(in T, sink Sink) {
		if !sink.WantMore() {
			return
		}
		var zero T
		if in != zero {
			sink.Error(ErrWantZero)
		}
	}
}
