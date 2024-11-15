package kiter

import "iter"

func Map[Vin, Vout any](in iter.Seq[Vin], fn func(Vin) Vout) iter.Seq[Vout] {
	return func(yield func(Vout) bool) {
		for x := range in {
			if !yield(fn(x)) {
				return
			}
		}
	}
}