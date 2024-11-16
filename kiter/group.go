package kiter

import "iter"

func Grouped[K comparable, V any](in iter.Seq2[K, V]) iter.Seq2[K, iter.Seq[V]] {
	return func(yield func(K, iter.Seq[V]) bool) {
		m := map[K][]V{}
		for k, v := range in {
			m[k] = append(m[k], v)
		}
		for k, vs := range m {
			if !yield(k, FromSlice(vs)) {
				break
			}
		}
	}
}