package kiter

import "iter"

type KV[K, V any] struct {
	K K
	V V
}

func ToPairs[K, V any](in iter.Seq2[K, V]) iter.Seq[KV[K, V]] {
	return func(yield func(KV[K, V]) bool) {
		for k, v := range in {
			if !yield(KV[K, V]{k, v}) {
				return
			}
		}
	}
}

func FromPairs[K, V any](pairs iter.Seq[KV[K, V]]) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for p := range pairs {
			if !yield(p.K, p.V) {
				return
			}
		}
	}
}
