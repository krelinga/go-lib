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

func Map2[Kin, Vin, Kout, Vout any](in iter.Seq2[Kin, Vin], fn func(Kin, Vin) (Kout, Vout)) iter.Seq2[Kout, Vout] {
	return func(yield func(Kout, Vout) bool) {
		for k, v := range in {
			k2, v2 := fn(k, v)
			if !yield(k2, v2) {
				return
			}
		}
	}
}
