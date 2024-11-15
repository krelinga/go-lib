package kiter

import "iter"

type Pair[K, V any] struct {
	Key   K
	Value V
}

func ToPairs[K, V any](in iter.Seq2[K, V]) iter.Seq[Pair[K, V]] {
	return func(yield func(Pair[K, V]) bool) {
		for k, v := range in {
			if !yield(Pair[K, V]{k, v}) {
				return
			}
		}
	}
}

func FromPairs[K, V any](pairs iter.Seq[Pair[K, V]]) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for p := range pairs {
			if !yield(p.Key, p.Value) {
				return
			}
		}
	}
}