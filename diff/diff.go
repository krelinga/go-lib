package diff

import (
	"errors"
	"reflect"
)

var ErrUnsupportedType = errors.New("unsupported type")

func Diff[T any](lhs, rhs T) (bool, error) {
	return diffWithReflection(lhs, rhs)
}

func diffWithReflection(lhs, rhs any) (bool, error) {
	lhsType := reflect.TypeOf(lhs)
	rhsType := reflect.TypeOf(rhs)
	if lhsType != rhsType {
		return true, nil
	}
	
	if lhsType.Kind() == reflect.Pointer {
		return diffPointer(lhs, rhs)
	} else if lhsType.Comparable() {
		return diffComparable(lhs, rhs), nil
	}

	return false, ErrUnsupportedType
}

func valueOk(value reflect.Value) bool {
	return value.IsValid() && !value.IsNil()
}

func diffPointer(lhs, rhs any) (bool, error) {
	lhsValue := reflect.ValueOf(lhs)
	rhsValue := reflect.ValueOf(rhs)
	lhsOk := valueOk(lhsValue)
	rhsOk := valueOk(rhsValue)
	if !lhsOk && !rhsOk {
		return false, nil
	} else if lhsOk && rhsOk {
		return diffWithReflection(lhsValue.Elem().Interface(), rhsValue.Elem().Interface())
	}

	return true, nil
}

func diffComparable(lhs, rhs any) bool {
	return rhs != lhs
}