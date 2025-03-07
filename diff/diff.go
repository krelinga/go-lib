package diff

import (
	"fmt"
	"reflect"

	"github.com/krelinga/go-lib/datapath"
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
	Path datapath.Path
	Lhs  interface{}
	Rhs  interface{}
	Kind Kind
}

func Diff[T any](lhs, rhs T) []Result {
	results := diffWithReflection(datapath.Path{}, newVt(lhs), newVt(rhs))
	if len(results) == 0 {
		return nil
	}
	return results
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

func (v vt) TypeName() string {
	// TODO: this may not work for types other than pointers & specifically named types.
	prefix := ""
	for v.Type.Kind() == reflect.Ptr {
		prefix += "*"
		v = v.Elem()
	}
	return prefix + v.Type.Name()
}

func (v vt) CallIsSafe() bool {
	if !v.Value.IsValid() {
		return false
	}
	switch v.Value.Kind() {
	case reflect.Interface: fallthrough
	case reflect.Ptr: fallthrough
	case reflect.Slice: fallthrough
	case reflect.Map: fallthrough
	case reflect.Chan: fallthrough
	case reflect.Func:
		if v.Value.IsNil() {
			return false
		}
	}
	return true
}

func (v vt) Call(name string) vt {

	typeMethod, ok := v.Type.MethodByName(name)
	if !ok {
		panic(fmt.Sprintf("method %s not found on type %s", name, v.Type))
	}
	if typeMethod.Type.NumOut() != 1 {
		panic(fmt.Sprintf("method %s must have exactly one return value", name))
	}
	valueMethod := v.Value.MethodByName(name)
	if !valueMethod.IsValid() {
		panic(fmt.Sprintf("method %s not found on value %s", name, v.Type))
	}
	outs := valueMethod.Call([]reflect.Value{})
	return vt{
		Value: outs[0],
		Type:  typeMethod.Type.Out(0),
	}
}

func diffWithReflection(p datapath.Path, lhs, rhs vt) []Result {
	if lhs.Type != rhs.Type {
		panic("lhs and rhs must be of the same type")
	}

	results := []Result{}
	opts := globalDb.lookup(lhs.Type)
	if opts != nil {
		results = append(results, diffMethods(opts.methods, p, lhs, rhs)...)
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
		results = append(results, diffInterface(p, lhs, rhs)...)
	} else if lhs.Type.Kind() == reflect.Pointer {
		results = append(results, diffPointer(p, lhs, rhs)...)
	} else if lhs.Type.Kind() == reflect.Slice {
		results = append(results, diffSlice(p, lhs, rhs)...)
	} else if lhs.Type.Kind() == reflect.Map {
		results = append(results, diffMap(p, lhs, rhs)...)
	} else if lhs.Type.Kind() == reflect.Struct {
		results = append(results, diffStruct(p, lhs, rhs)...)
	} else if lhs.Type.Comparable() {
		results = append(results, diffComp(p, lhs, rhs)...)
	} else {
		panic(fmt.Sprintf("unsupported type: %v", lhs.Type))
	}
	return results
}

func diffMethods(methods []string, p datapath.Path, lhs, rhs vt) []Result {
	// Only diff methods if both lhs and rhs are valid and non-nil.  The other
	// cases will be caught by the other diff functions.
	if !lhs.CallIsSafe() || !rhs.CallIsSafe() {
		return nil
	}
	results := []Result{}
	for _, name := range methods {
		lhsOut := lhs.Call(name)
		rhsOut := rhs.Call(name)
		if diff := diffWithReflection(p.Method(name), lhsOut, rhsOut); diff != nil {
			results = append(results, diff...)
		}
	}
	return results
}

func diffPointer(p datapath.Path, lhs, rhs vt) []Result {
	if !lhs.Value.IsValid() && !rhs.Value.IsValid() {
		return nil
	}
	if !lhs.Value.IsValid() || !rhs.Value.IsValid() {
		return []Result{
			{
				Path: p,
				Lhs:  lhs.Interface(),
				Rhs:  rhs.Interface(),
				Kind: Different,
			},
		}
	}
	if lhs.Value.IsNil() && rhs.Value.IsNil() {
		// Both are nil
		return nil
	}
	if lhs.Value.IsNil() || rhs.Value.IsNil() {
		return []Result{
			{
				Path: p,
				Lhs:  lhs.Interface(),
				Rhs:  rhs.Interface(),
				Kind: Different,
			},
		}
	}
	return diffWithReflection(p.PtrDeref(), lhs.Elem(), rhs.Elem())
}

func diffComp(p datapath.Path, lhs, rhs vt) []Result {
	if !lhs.Value.Equal(rhs.Value) {
		return []Result{
			{
				Path: p,
				Lhs:  lhs.Interface(),
				Rhs:  rhs.Interface(),
				Kind: Different,
			},
		}
	}
	return nil
}

func diffInterface(p datapath.Path, lhs, rhs vt) []Result {
	lhs = lhs.ResolveInterface()
	rhs = rhs.ResolveInterface()
	if lhs.Type == nil && rhs.Type == nil {
		// This triggers if lhs and rhs were both nil interface values.
		return nil
	}
	if lhs.Type != rhs.Type {
		// This triggers if lhs and rhs were both non-nil interface values with different underlying types.
		return []Result{
			{
				Path: p,
				Lhs:  lhs.Interface(),
				Rhs:  rhs.Interface(),
				Kind: Different,
			},
		}
	}
	return diffWithReflection(p.TypeAssert(lhs.TypeName()), lhs, rhs)
}

func diffSlice(p datapath.Path, lhs, rhs vt) []Result {
	// TODO: is this block actually needed?  These should probably be reported as extra/missing.
	if lhs.Value.IsValid() != rhs.Value.IsValid() {
		// Only one of the instances is nil.
		return []Result{
			{
				Path: p,
				Lhs:  lhs.Interface(),
				Rhs:  rhs.Interface(),
				Kind: Different,
			},
		}
	}
	if !lhs.Value.IsValid() && !rhs.Value.IsValid() {
		// Both instances are nil.
		return nil
	}
	results := []Result{}
	i := 0
	for ; i < lhs.Value.Len() && i < rhs.Value.Len(); i++ {
		if diff := diffWithReflection(p.Index(i), lhs.Index(i), rhs.Index(i)); diff != nil {
			results = append(results, diff...)
		}
	}
	for ; i < lhs.Value.Len(); i++ {
		results = append(results, Result{
			Path: p.Index(i),
			Lhs:  lhs.Index(i).Interface(),
			Kind: Missing,
		})
	}
	for ; i < rhs.Value.Len(); i++ {
		results = append(results, Result{
			Path: p.Index(i),
			Rhs:  rhs.Index(i).Interface(),
			Kind: Extra,
		})
	}
	return results
}

func diffMap(p datapath.Path, lhs, rhs vt) []Result {
	// TODO: is this block actually needed?  These should probably be reported as extra/missing.
	if lhs.Value.IsValid() != rhs.Value.IsValid() {
		// Only one of the instances is nil.
		return []Result{
			{
				Path: p,
				Lhs:  lhs.Interface(),
				Rhs:  rhs.Interface(),
				Kind: Different,
			},
		}
	}
	if !lhs.Value.IsValid() && !rhs.Value.IsValid() {
		// Both instances are nil.
		return nil
	}

	results := []Result{}
	for _, key := range lhs.Value.MapKeys() {
		lhsFound := lhs.MapIndex(key)
		rhsFound := rhs.MapIndex(key)
		if !rhsFound.Value.IsValid() {
			results = append(results, Result{
				Path: p.Key(key.Interface()),
				Lhs:  lhsFound.Interface(),
				Kind: Missing,
			})
		} else if diff := diffWithReflection(p.Key(key.Interface()), lhsFound, rhsFound); diff != nil {
			results = append(results, diff...)
		}
	}
	for _, key := range rhs.Value.MapKeys() {
		lhsFound := lhs.MapIndex(key)
		if !lhsFound.Value.IsValid() {
			results = append(results, Result{
				Path: p.Key(key.Interface()),
				Rhs:  rhs.MapIndex(key).Interface(),
				Kind: Extra,
			})
		}
	}

	return results
}

func diffStruct(p datapath.Path, lhs, rhs vt) []Result {
	if lhs.Value.IsValid() != rhs.Value.IsValid() {
		// Only one of the instances is nil.
		return []Result{
			{
				Path: p,
				Lhs:  lhs.Interface(),
				Rhs:  rhs.Interface(),
				Kind: Different,
			},
		}
	}
	if !lhs.Value.IsValid() && !rhs.Value.IsValid() {
		// Both instances are nil.
		return nil
	}
	results := []Result{}
	for _, f := range reflect.VisibleFields(lhs.Type) {
		isFieldOfNestedStruct := len(f.Index) > 1
		if !f.IsExported() || isFieldOfNestedStruct {
			// Skip non-exported fields since we can't access their contents.
			// Skip fields of nested structs since we'll visit them when we visit the field
			// that corresponds to the nested struct.
			continue
		}
		if diff := diffWithReflection(p.Field(f.Name), lhs.StructField(f), rhs.StructField(f)); diff != nil {
			results = append(results, diff...)
		}
	}
	return results
}

// There's some really subtle stuff going on between reflect.Value.IsNil() and reflect.Value.IsValid().
// AFAICT the difference is that nil interface values and the literal nil are considered invalid (i.e. things where we have no way to tell their type), but
// nil pointers are considered valid (i.e. we know their type, and we know they're nil).
