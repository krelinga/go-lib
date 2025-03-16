package kfscopier

import "errors"

var (
	ErrReqSrcNotStat = errors.New("request Src failed to stat")
	ErrReqSrcNotFile = errors.New("request Src is not a file")
	ErrReqDestNotStat = errors.New("request Dest failed to stat")
	ErrReqDestExists = errors.New("request Dest already exists")

	ErrReqSrcOpen = errors.New("request Src failed to open")
	ErrReqDestCreate = errors.New("request Dest failed to create")
	ErrReqSrcRead = errors.New("request Src failed to read")
	ErrReqDestWrite = errors.New("request Dest failed to write")
)

type ReqError interface {
	Req() *Req
}

type reqError struct {
	req *Req
	err error
}

func (re *reqError) Req() *Req {
	return re.req
}

func (re *reqError) Error() string {
	return re.err.Error()
}

func (re *reqError) Unwrap() error {
	return re.err
}
