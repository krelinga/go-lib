package kiter

import "iter"

func FromMap[K comparable, V any](m map[K]V) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for k, v := range m {
			if !yield(k, v) {
				break
			}
		}
	}
}

func ToMap[K comparable, V any](in iter.Seq2[K, V]) map[K]V {
	m := map[K]V{}
	ToMapInsert(m, in)
	return m
}

func ToMapInsert[K comparable, V any](m map[K]V, in iter.Seq2[K, V]) {
	for k, v := range in {
		m[k] = v
	}
}

func FromMapKeys[K comparable, V any](m map[K]V) iter.Seq[K] {
	return func(yield func(K) bool) {
		for k := range m {
			if !yield(k) {
				break
			}
		}
	}
}
