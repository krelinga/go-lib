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

func (v vt) ResolveInterface() vt {
	// TODO: this needs more testing.
	// It looks like a reflect.Value will never have Kind() == reflect.Interface?
	// If that's true then only reflect.Type will have Kind() == reflect.Interface, and
	// any values associated with that time will have Kind() == their underlying kind.
	if v.Type.Kind() != reflect.Interface {
		panic("type is not an interface")
	}
	return vt{
		Value: v.Value,
		Type:  v.Value.Type(),
	}
}

func diffWithReflection(lhs, rhs vt) (bool, error) {
	if lhs.Type != rhs.Type {
		panic("lhs and rhs must be of the same type")
	}

	if lhs.Type.Kind() == reflect.Interface {
		// This is really subtle. If I take this if condition out, then we only get into
		// trouble because the interface type is comparable, and it happens by pointer
		// instead of by value of thing pointed to.
		//
		// On the other hand, the rules for how reflect.Value.Equal() works are not very
		// clear ... the docs say it will panic if the underlying types are not comparable,
		// but then proceeds to describe how slice comparison works ... but slices are not
		// comparable, so I don't know what to make of that.
		return diffInterface(lhs, rhs)
	} else if lhs.Type.Kind() == reflect.Pointer {
		return diffWithReflection(lhs.Elem(), rhs.Elem())
	} else if lhs.Type.Comparable() {
		return !lhs.Value.Equal(rhs.Value), nil
	}

	return false, ErrUnsupportedType
}

func diffInterface(lhs, rhs vt) (bool, error) {
	lhs = lhs.ResolveInterface()
	rhs = rhs.ResolveInterface()
	if lhs.Type == nil && rhs.Type == nil {
		// This triggers if lhs and rhs were both nil interface values.
		return false, nil
	}
	if lhs.Type != rhs.Type {
		// This triggers if lhs and rhs were both non-nil interface values with different underlying types.
		return true, nil
	}
	return diffWithReflection(lhs, rhs)
}
