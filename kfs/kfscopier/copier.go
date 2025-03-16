package kfscopier

import (
	"context"

	"github.com/krelinga/go-lib/pipe"
)

func New(ctx context.Context, reqs <-chan *Req, opts Options) <-chan error {
	opts.setDefaults()
	errChans := make([]<-chan error, opts.MaxParallelCopies)
	for i := range errChans {
		errChans[i] = worker(ctx, reqs, opts.ChunkSize)
	}

	return pipe.Merge(ctx, errChans...)
}