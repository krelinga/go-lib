package kfscopier

import (
	"fmt"
	"os"
)

const (
	DefaultMaxParallelCopies = 10
	DefaultChunkSize		 = 1024 * 1024
)

type Req struct {
	Src, Dest string
}

func (r *Req) Validate() error {
	if stat, err := os.Stat(r.Src); err != nil {
		return r.newError(fmt.Errorf("%w: %w", ErrReqSrcNotStat, err))
	} else if !stat.Mode().IsRegular() {
		return r.newError(ErrReqSrcNotFile)
	}

	if _, err := os.Stat(r.Dest); err == nil {
		return r.newError(ErrReqDestExists)
	} else if !os.IsNotExist(err) {
		return r.newError(fmt.Errorf("%w: %w", ErrReqDestNotStat, err))
	}

	return nil
}

func (r *Req) newError(err error) error {
	return &reqError{
		req: r,
		err: err,
	}
}