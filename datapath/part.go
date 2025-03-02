package datapath

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/krelinga/go-lib/kslice"
	"github.com/krelinga/go-lib/valid"
)

var (
	ErrInvalidField = errors.New("invalid field")
	ErrNegativeIndex = kslice.ErrNegativeIndex
)

type Part interface {
	partIsClosedType()

	valid.Validator
	fmt.Stringer
}

type Field string

func (f Field) partIsClosedType() {}

func (f Field) String() string {
	return string(f)
}

func (f Field) Validate() error {
	// TODO: look at the Go spec and implement other restrictions on field names.
	if f == "" {
		return ErrInvalidField
	}
	return nil
}

type Index kslice.Index

func (i Index) partIsClosedType() {}

func (i Index) String() string {
	return kslice.Index(i).String()
}

func (i Index) Validate() error {
	return kslice.Index(i).Validate()
}

type Key interface {
	Part

	Get() any
}

func NewKey[T comparable](k T) Key {
	return typedKey[T]{k: k}
}

type typedKey[T comparable] struct {
	k T
}

func (k typedKey[T]) partIsClosedType() {}

func (k typedKey[T]) String() string {
	return fmt.Sprintf("%v", k.k)
}

func (k typedKey[T]) Validate() error {
	return nil
}

func (k typedKey[T]) Get() any {
	return k.k
}

func NewReflectKey(k reflect.Value) Key {
	return reflectKey{k: k}
}

type reflectKey struct {
	k reflect.Value
}

func (k reflectKey) partIsClosedType() {}

func (k reflectKey) String() string {
	valid.OrPanic(k)
	return "" // TODO
}

func (k reflectKey) Validate() error {
	return nil  // TODO
}

func (k reflectKey) Get() any {
	valid.OrPanic(k)
	return k.k.Interface()
}