package pipe

// spell-checker:ignore chans

import (
	"context"
	"fmt"
)

func ParDoErr[inType any, outType any, chanInType readable[inType]](ctx context.Context, parallelism int, in chanInType, f func(inType) (outType, error)) (<-chan outType, <-chan error) {
	outs := make([]chan outType, parallelism)
	errors := make([]chan error, parallelism)
	for i := 0; i < parallelism; i++ {
		i := i
		outs[i] = make(chan outType)
		errors[i] = make(chan error)
		go func() {
			defer close(outs[i])
			defer close(errors[i])
			for inVal := range in {
				outVal, err := f(inVal)
				if err != nil {
					if !TryWrite(ctx, errors[i], err) {
						return
					}
					continue
				}
				if !TryWrite(ctx, outs[i], outVal) {
					return
				}
			}
		}()
	}

	return Merge(ctx, outs...), Merge(ctx, errors...)
}

// ParDo() runs a function in parallel on each value from the input channel.
func ParDo[inType any, outType any](ctx context.Context, parallelism int, in <-chan inType, f func(inType) outType) <-chan outType {
	outs, errs := ParDoErr(ctx, parallelism, in, func(in inType) (outType, error) {
		return f(in), nil
	})
	go func() {
		for err := range errs {
			panic(fmt.Sprint("unexpected error in Parallel: ", err))
		}
	}()
	return outs
}
