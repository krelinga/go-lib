package diff

import (
	"errors"
	"reflect"
)

var ErrUnsupportedType = errors.New("unsupported type")

func Diff[T any](lhs, rhs T) (bool, error) {
	return diffWithReflection(newVt(lhs), newVt(rhs))
}

type vt struct {
	Value reflect.Value
	Type  reflect.Type
}

func newVt[T any](v T) vt {
	return vt{
		Value: reflect.ValueOf(v),
		Type:  reflect.TypeFor[T](),
	}
}

func (v vt) Elem() vt {
	return vt{
		Value: v.Value.Elem(),
		Type:  v.Type.Elem(),
	}
}

func diffWithReflection(lhs, rhs vt) (bool, error) {
	if lhs.Type != rhs.Type {
		panic("lhs and rhs must be of the same type")
	}

	if lhs.Type.Kind() == reflect.Pointer {
		return diffPointer(lhs, rhs)
	} else if lhs.Type.Comparable() {
		return diffComparable(lhs, rhs), nil
	}

	return false, ErrUnsupportedType
}

func diffPointer(lhs, rhs vt) (bool, error) {
	return diffWithReflection(lhs.Elem(), rhs.Elem())
}

func diffComparable(lhs, rhs vt) bool {
	return !lhs.Value.Equal(rhs.Value)
}
