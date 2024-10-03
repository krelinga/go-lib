package pipe

// spell-checker:ignore chans

import (
	"context"
	"sync"
)

// A more-generic implementation of Merge() that can merge any readable channel type.
// This is useful in situations where generic code needs to call Merge(), because the go compiler
// is very picky about how types are inferred in generic functions.
// Not exposing this on the public API because the standard go method of expressing a read-only channel
// is more readable.
func mergeImpl[t any, chanT readable[t]](ctx context.Context, channels ...chanT) <-chan t {
	out := make(chan t, len(channels))
	var wg sync.WaitGroup
	wg.Add(len(channels))
	for _, c := range channels {
		c := c
		go func() {
			defer wg.Done()
			for v := range c {
				if !TryWrite(ctx, out, v) {
					return
				}
			}
		}()
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

// Merge() merges multiple channels into a single channel.
func Merge[t any](ctx context.Context, channels ...<-chan t) <-chan t {
	return mergeImpl[t](ctx, channels...)
}