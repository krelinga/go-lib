package diff

import (
	"fmt"
	"reflect"
)

type Result int

func (r Result) String() string {
	switch r {
	case None:
		return "None"
	case ValueDiff:
		return "ValueDiff"
	case TypeDiff:
		return "TypeDiff"
	case Missing:
		return "Missing"
	case Extra:
		return "Extra"
	default:
		panic(fmt.Sprintf("unknown diff result: %d", int(r)))
	}
}

const (
	None Result = iota
	ValueDiff
	TypeDiff
	Missing
	Extra
)

func Diff[T any](lhs, rhs T) Result {
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

func (v vt) MapIndex(key reflect.Value) vt {
	return vt{
		Value: v.Value.MapIndex(key),
		Type:  v.Type.Elem(),
	}
}

func (v vt) StructField(sf reflect.StructField) vt {
	return vt{
		Value: v.Value.FieldByIndex(sf.Index),
		Type:  sf.Type,
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

func diffWithReflection(lhs, rhs vt) Result {
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
	} else if lhs.Type.Kind() == reflect.Map {
		return diffMap(lhs, rhs)
	} else if lhs.Type.Kind() == reflect.Struct {
		return diffStruct(lhs, rhs)
	} else if lhs.Type.Comparable() {
		return diffComp(lhs, rhs)
	}

	panic(fmt.Sprintf("unsupported type: %v", lhs.Type))
}

func diffComp(lhs, rhs vt) Result {
	if !lhs.Value.Equal(rhs.Value){
		return ValueDiff
	}
	return None
}

func diffInterface(lhs, rhs vt) Result {
	lhs = lhs.ResolveInterface()
	rhs = rhs.ResolveInterface()
	if lhs.Type == nil && rhs.Type == nil {
		// This triggers if lhs and rhs were both nil interface values.
		return None
	}
	if lhs.Type != rhs.Type {
		// This triggers if lhs and rhs were both non-nil interface values with different underlying types.
		return TypeDiff
	}
	return diffWithReflection(lhs, rhs)
}

func diffSlice(lhs, rhs vt) Result {
	if lhs.Value.IsValid() != rhs.Value.IsValid() {
		// Only one of the instances is nil.
		return ValueDiff
	}
	if !lhs.Value.IsValid() && !rhs.Value.IsValid() {
		// Both instances are nil.
		return None
	}
	if lhs.Value.Len() < rhs.Value.Len() {
		return Extra
	} else if lhs.Value.Len() > rhs.Value.Len() {
		return Missing
	}
	for i := 0; i < lhs.Value.Len(); i++ {
		if diff := diffWithReflection(lhs.Index(i), rhs.Index(i)); diff != None {
			return diff
		}
	}
	return None
}

func diffMap(lhs, rhs vt) Result {
	if lhs.Value.IsValid() != rhs.Value.IsValid() {
		// Only one of the instances is nil.
		return ValueDiff
	}
	if !lhs.Value.IsValid() && !rhs.Value.IsValid() {
		// Both instances are nil.
		return None
	}
	if lhs.Value.Len() < rhs.Value.Len() {
		return Extra
	} else if lhs.Value.Len() > rhs.Value.Len() {
		return Missing
	}
	for _, key := range lhs.Value.MapKeys() {
		if diff := diffWithReflection(lhs.MapIndex(key), rhs.MapIndex(key)); diff != None {
			return diff
		}
	}
	return None
}

func diffStruct(lhs, rhs vt) Result {
	if lhs.Value.IsValid() != rhs.Value.IsValid() {
		// Only one of the instances is nil.
		return ValueDiff
	}
	if !lhs.Value.IsValid() && !rhs.Value.IsValid() {
		// Both instances are nil.
		return None
	}
	for _, f := range reflect.VisibleFields(lhs.Type) {
		isFieldOfNestedStruct := len(f.Index) > 1
		if !f.IsExported() || isFieldOfNestedStruct {
			// Skip non-exported fields since we can't access their contents.
			// Skip fields of nested structs since we'll visit them when we visit the field
			// that corresponds to the nested struct.
			continue
		}
		if diff := diffWithReflection(lhs.StructField(f), rhs.StructField(f)); diff != None {
			return diff
		}
	}
	return None
}
