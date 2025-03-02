package kslice

import (
	"errors"
	"fmt"

	"github.com/krelinga/go-lib/valid"
)

type Index int

var ErrNegativeIndex = errors.New("kslice.Index must be >= 0")

func (i Index) Validate() error {
	if i < 0 {
		return ErrNegativeIndex
	}
	return nil
}

func (i Index) String() string {
	valid.OrPanic(i)
	return fmt.Sprintf("%d", int(i))
}