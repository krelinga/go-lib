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

func (v vt) Index(i int) vt {
	return vt{
		Value: v.Value.Index(i),
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
	var newType reflect.Type
	if v.Value.IsValid() {
		newType = v.Value.Type()
	}
	return vt{
		Value: v.Value,
		Type:  newType,
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
	} else if lhs.Type.Kind() == reflect.Slice {
		return diffSlice(lhs, rhs)
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

func diffSlice(lhs, rhs vt) (bool, error) {
	if lhs.Value.IsValid() != rhs.Value.IsValid() {
		// Only one of the instances is nil.
		return true, nil
	}
	if !lhs.Value.IsValid() && !rhs.Value.IsValid() {
		// Both instances are nil.
		return false, nil
	}
	if lhs.Value.Len() != rhs.Value.Len() {
		return true, nil
	}
	for i := 0; i < lhs.Value.Len(); i++ {
		if diff, err := diffWithReflection(lhs.Index(i), rhs.Index(i)); diff || err != nil {
			return diff, err
		}
	}
	return false, nil
}