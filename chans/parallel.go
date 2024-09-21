package chans

// spell-checker:ignore chans

import "fmt"

func each[inType any, outType any](f func(inType) outType, in ...inType) []outType {
	out := make([]outType, len(in))
	for i, v := range in {
		out[i] = f(v)
	}
	return out
}

func ParallelErr[inType any, outType any](parallelism int, in <-chan inType, f func(inType) (outType, error)) (<-chan outType, <-chan error) {
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
					errors[i] <- err
					continue
				}
				outs[i] <- outVal
			}
		}()
	}

	roOuts := each(ReadOnly, outs...)
	roErrors := each(ReadOnly, errors...)

	return Merge(roOuts...), Merge(roErrors...)
}

// Parallel() runs a function in parallel on each value from the input channel.
func Parallel[inType any, outType any](parallelism int, in <-chan inType, f func(inType) outType) <-chan outType {
	outs, errs := ParallelErr(parallelism, in, func(in inType) (outType, error) {
		return f(in), nil
	})
	go func() {
		for err := range errs {
			panic(fmt.Sprint("unexpected error in Parallel: ", err))
		}
	}()
	return outs
}
