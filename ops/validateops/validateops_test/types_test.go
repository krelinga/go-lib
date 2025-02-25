package validateops_test

import "github.com/krelinga/go-lib/ops/validateops"

type ErrWrapper struct {
	Err error
}

func (e ErrWrapper) ValidateOp(sink validateops.Sink) {
	if e.Err != nil {
		sink.Error(e.Err)
	}
}

type PtrErrWrapper struct {
	Err error
}

func (e *PtrErrWrapper) ValidateOp(sink validateops.Sink) {
	if e.Err != nil {
		sink.Error(e.Err)
	}
}