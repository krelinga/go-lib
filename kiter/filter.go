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
