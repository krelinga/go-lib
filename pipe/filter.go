package pipe

import "context"

func ParDoFilterErr[inType any](ctx context.Context, parallelism int, in <-chan inType, f func(inType) (bool, error)) (<-chan inType, <-chan error) {
	out := make(chan inType)
	errors := make(chan error)

	go func() {
		defer close(out)
		defer close(errors)

		for inVal := range in {
			ok, err := f(inVal)
			if err != nil {
				if !TryWrite(ctx, errors, err) {
					return
				}
				continue
			}
			if ok {
				if !TryWrite(ctx, out, inVal) {
					return
				}
			}
		}
	}()

	return out, errors
}

func ParDoFilter[inType any](ctx context.Context, parallelism int, in <-chan inType, f func(inType) bool) <-chan inType {
	outs, errs := ParDoFilterErr(ctx, parallelism, in, func(in inType) (bool, error) {
		return f(in), nil
	})
	go func() {
		for err := range errs {
			panic("unexpected error in ParDoFilter: " + err.Error())
		}
	}()
	return outs
}