package kfscopier

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/krelinga/go-lib/pipe"
)

func copyOneFile(ctx context.Context, req *Req, chunkSize int) error {
	if err := req.Validate(); err != nil {
		return err
	}
	in, err := os.Open(req.Src)
	if err != nil {
		return req.newError(ErrReqSrcOpen)
	}
	defer in.Close()
	out, err := os.Create(req.Dest)
	if err != nil {
		return req.newError(ErrReqDestCreate)
	}
	defer out.Close()

	buf := make([]byte, chunkSize)
	for {
		if err := ctx.Err(); err != nil {
			return err
		}
		b, err := in.Read(buf)
		if errors.Is(err, io.EOF) {
			break
		} else if err != nil {
			return req.newError(fmt.Errorf("%w: %w", ErrReqSrcRead, err))
		}

		if err := ctx.Err(); err != nil {
			return err
		}

		if _, err = out.Write(buf[:b]); err != nil {
			return req.newError(fmt.Errorf("%w: %w", ErrReqDestWrite, err))
		}
	}
	return nil
}

func worker(ctx context.Context, reqs <-chan *Req, chunkSize int) <-chan error {
	errs := make(chan error, 1)
	go func() {
		defer close(errs)
		for {
			select {
			case <-ctx.Done():
				return
			case req, ok := <-reqs:
				if !ok {
					return
				}
				if err := copyOneFile(ctx, req, chunkSize); err != nil {
					if !pipe.TryWrite(ctx, errs, err) {
						return
					}
					continue
				}
			}
		}
	}()
	return errs
}