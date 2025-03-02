package valid

import (
	"errors"
	"reflect"
)

// Children() uses reflection to recursively discover any Validator implementations in the input and call Validate() on them.
// The returned error is the combined result of all Validate() calls (created by errors.Join()), or nil if there are no errors.
//
// This function is intended to be called from within Validate() implementations, and so it does not call Validate() on the input itself (to avoid infinite recursion).
// The other intended use case for this function is to validate the contents of a struct, slice, or map that does not itself satisfy the Validator interface.
//
// Children() returns nil if the input is nil.
func Children(input any) error {
	v := reflect.ValueOf(input)
	if !v.IsValid() {
		return nil
	}

	// Unwrap any layers of pointers or interfaces, returning early if we find a nil.
	if v.Kind() == reflect.Ptr || v.Kind() == reflect.Interface {
		if v.IsNil() {
			return nil
		}
		return recurse(v.Elem().Interface())
	}

	var errs []error
	switch v.Kind() {
	case reflect.Struct:
		fields := reflect.VisibleFields(v.Type())
		for _, f := range fields {
			if !f.IsExported() {
				continue
			}
			if err := recurse(v.FieldByIndex(f.Index).Interface()); err != nil {
				errs = append(errs, err)
			}
		}
	case reflect.Slice:
		for i := 0; i < v.Len(); i++ {
			if err := recurse(v.Index(i).Interface()); err != nil {
				errs = append(errs, err)
			}
		}
	case reflect.Map:
		for _, key := range v.MapKeys() {
			if err := recurse(key.Interface()); err != nil {
				errs = append(errs, err)
			}
			if err := recurse(v.MapIndex(key).Interface()); err != nil {
				errs = append(errs, err)
			}
		}
	}

	return errors.Join(errs...)
}

func recurse(in any) error {
	if validator, ok := in.(Validator); ok {
		return validator.Validate()
	}

	return Children(in)
}
