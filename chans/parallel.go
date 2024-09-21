package chans

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