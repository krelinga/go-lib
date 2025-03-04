package diff

import (
	"fmt"
	"reflect"
)

type Kind int

func (r Kind) String() string {
	switch r {
	case Same:
		return "Same"
	case Different:
		return "Different"
	case Missing:
		return "Missing"
	case Extra:
		return "Extra"
	default:
		panic(fmt.Sprintf("unknown diff result: %d", int(r)))
	}
}

const (
	Same Kind = iota
	Different
	Missing
	Extra
)

type Result struct {
	Lhs  interface{}
	Rhs  interface{}
	Kind Kind
}

func Diff[T any](lhs, rhs T) *Result {
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

func (v vt) Interface() any {
	if v.Value.IsValid() {
		return v.Value.Interface()
	}
	return nil
}

func diffWithReflection(lhs, rhs vt) *Result {
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
		return diffPointer(lhs, rhs)
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

func diffPointer(lhs, rhs vt) *Result {
	if !lhs.Value.IsValid() && !rhs.Value.IsValid() {
		return nil
	}
	if !lhs.Value.IsValid() || !rhs.Value.IsValid() {
		return &Result{
			Lhs:  lhs.Interface(),
			Rhs:  rhs.Interface(),
			Kind: Different,
		}
	}
	if lhs.Value.IsNil() && rhs.Value.IsNil() {
		// Both are nil
		return nil
	}
	if lhs.Value.IsNil() || rhs.Value.IsNil() {
		return &Result{
			Lhs:  lhs.Interface(),
			Rhs:  rhs.Interface(),
			Kind: Different,
		}
	}
	return diffWithReflection(lhs.Elem(), rhs.Elem())
}

func diffComp(lhs, rhs vt) *Result {
	if !lhs.Value.Equal(rhs.Value) {
		return &Result{
			Lhs:  lhs.Interface(),
			Rhs:  rhs.Interface(),
			Kind: Different,
		}
	}
	return nil
}

func diffInterface(lhs, rhs vt) *Result {
	lhs = lhs.ResolveInterface()
	rhs = rhs.ResolveInterface()
	if lhs.Type == nil && rhs.Type == nil {
		// This triggers if lhs and rhs were both nil interface values.
		return nil
	}
	if lhs.Type != rhs.Type {
		// This triggers if lhs and rhs were both non-nil interface values with different underlying types.
		return &Result{
			Lhs:  lhs.Interface(),
			Rhs:  rhs.Interface(),
			Kind: Different,
		}
	}
	return diffWithReflection(lhs, rhs)
}

func diffSlice(lhs, rhs vt) *Result {
	// TODO: is this block actually needed?  These should probably be reported as extra/missing.
	if lhs.Value.IsValid() != rhs.Value.IsValid() {
		// Only one of the instances is nil.
		return &Result{
			Lhs:  lhs.Interface(),
			Rhs:  rhs.Interface(),
			Kind: Different,
		}
	}
	if !lhs.Value.IsValid() && !rhs.Value.IsValid() {
		// Both instances are nil.
		return nil
	}
	if lhs.Value.Len() < rhs.Value.Len() {
		return &Result{
			Rhs:  rhs.Index(lhs.Value.Len()).Interface(),
			Kind: Extra,
		}
	} else if lhs.Value.Len() > rhs.Value.Len() {
		return &Result{
			Lhs:  lhs.Index(rhs.Value.Len()).Interface(),
			Kind: Missing,
		}
	}
	for i := 0; i < lhs.Value.Len(); i++ {
		if diff := diffWithReflection(lhs.Index(i), rhs.Index(i)); diff != nil {
			return diff
		}
	}
	return nil
}

func diffMap(lhs, rhs vt) *Result {
	// TODO: is this block actually needed?  These should probably be reported as extra/missing.
	if lhs.Value.IsValid() != rhs.Value.IsValid() {
		// Only one of the instances is nil.
		return &Result{
			Lhs:  lhs.Interface(),
			Rhs:  rhs.Interface(),
			Kind: Different,
		}
	}
	if !lhs.Value.IsValid() && !rhs.Value.IsValid() {
		// Both instances are nil.
		return nil
	}

	for _, key := range lhs.Value.MapKeys() {
		lhsFound := lhs.MapIndex(key)
		rhsFound := rhs.MapIndex(key)
		if !rhsFound.Value.IsValid() {
			return &Result{
				Lhs:  key.Interface(),
				Kind: Missing,
			}
		} else if diff := diffWithReflection(lhsFound, rhsFound); diff != nil {
			return diff
		}
	}
	for _, key := range rhs.Value.MapKeys() {
		lhsFound := lhs.MapIndex(key)
		if !lhsFound.Value.IsValid() {
			return &Result{
				Rhs:  key.Interface(),
				Kind: Extra,
			}
		}
	}

	return nil
}

func diffStruct(lhs, rhs vt) *Result {
	if lhs.Value.IsValid() != rhs.Value.IsValid() {
		// Only one of the instances is nil.
		return &Result{
			Lhs:  lhs.Interface(),
			Rhs:  rhs.Interface(),
			Kind: Different,
		}
	}
	if !lhs.Value.IsValid() && !rhs.Value.IsValid() {
		// Both instances are nil.
		return nil
	}
	for _, f := range reflect.VisibleFields(lhs.Type) {
		isFieldOfNestedStruct := len(f.Index) > 1
		if !f.IsExported() || isFieldOfNestedStruct {
			// Skip non-exported fields since we can't access their contents.
			// Skip fields of nested structs since we'll visit them when we visit the field
			// that corresponds to the nested struct.
			continue
		}
		if diff := diffWithReflection(lhs.StructField(f), rhs.StructField(f)); diff != nil {
			return diff
		}
	}
	return nil
}

// Next Steps:
// - Change Result to a struct, and rename the enum to Kind.
// - Result should surface Lhs and Rhs as any values, and Kind to indicate what sort of result it is.
// - Get rid of Same, and just return nil if there are no differences.

// There's some really subtle stuff going on between reflect.Value.IsNil() and reflect.Value.IsValid().
// AFAICT the difference is that nil interface values and the literal nil are considered invalid (i.e. things where we have no way to tell their type), but
// nil pointers are considered valid (i.e. we know their type, and we know they're nil).
