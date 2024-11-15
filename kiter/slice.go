package kiter

import "iter"

func FromSlice[V any](slice []V) iter.Seq[V] {
	return func(yield func(V) bool) {
		for _, v := range slice {
			if !yield(v) {
				break
			}
		}
	}
}