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

func FromKVSlice[K, V any](slice []KV[K, V]) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for _, kv := range slice {
			if !yield(kv.K, kv.V) {
				break
			}
		}
	}
}

func ToKVSlice[K, V any](in iter.Seq2[K, V]) []KV[K, V] {
	slice := []KV[K, V]{}
	return AppendToKVSlice(slice, in)
}

func AppendToKVSlice[K, V any](slice []KV[K, V], in iter.Seq2[K, V]) []KV[K, V] {
	for k, v := range in {
		slice = append(slice, KV[K, V]{K: k, V: v})
	}
	return slice
}