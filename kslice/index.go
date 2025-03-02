package kslice

import "errors"

type Index int

var ErrNegativeIndex = errors.New("kslice.Index must be >= 0")

func (i Index) Validate() error {
	if i < 0 {
		return ErrNegativeIndex
	}
	return nil
}