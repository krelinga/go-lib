package kiter

import "iter"

func Filter[V any](in iter.Seq[V], pred func(V) bool) iter.Seq[V] {
	return func(yield func(V) bool) {
		for v := range in {
			if !pred(v) {
				continue
			}
			if !yield(v) {
				break
			}
		}
	}
}

func Filter2[K, V any](in iter.Seq2[K, V], pred func(K, V) bool) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for k, v := range in {
			if !pred(k, v) {
				continue
			}
			if !yield(k, v) {
				break
			}
		}
	}
}