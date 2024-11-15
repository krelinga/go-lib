package kiter

import (
	"cmp"
	"iter"
	"slices"
)

func Sorted2[K cmp.Ordered, V any](in iter.Seq2[K, V]) iter.Seq2[K, V] {
	slice := ToSlice(ToPairs(in))
	slices.SortFunc(slice, func(a, b KV[K, V]) int {
		return cmp.Compare(a.K, b.K)
	})
	return FromPairs(FromSlice(slice))
}
