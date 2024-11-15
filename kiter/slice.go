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

func ToSlice[V any](in iter.Seq[V]) []V {
	slice := []V{}
	return AppendToSlice(slice, in)
}

func AppendToSlice[V any](slice []V, in iter.Seq[V]) []V {
	for v := range in {
		slice = append(slice, v)
	}
	return slice
}
