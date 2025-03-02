package diff

import (
	"errors"
	"reflect"

	"github.com/krelinga/go-lib/valid"
)

type Struct[T any] struct {
	AllExportedFields bool
	Fields            map[string]AnyDiffer
	Methods           map[string]AnyDiffer
}

func (s Struct[T]) typedDiff(lhs, rhs T) []Result {
	valid.OrPanic(s)
	return nil
}

func (Struct[T]) accepts(t reflect.Type) bool {
	return reflect.TypeFor[T]() == t
}

func (s Struct[T]) anyDiff(lhs, rhs any) []Result {
	return nil  // TODO
}

func (s Struct[T]) Validate() error {
	errs := []error{valid.Children(s)}

	// TODO: fix this
	// tType := s.diffType()
	// for fieldName, differ := range s.Fields {
	// 	if field, ok := tType.FieldByName(fieldName); !ok {
	// 		errs = append(errs, fmt.Errorf("%w: %s", ErrFieldNotFound, fieldName))
	// 	} else if !field.IsExported() {
	// 		errs = append(errs, fmt.Errorf("%w: %s", ErrFieldNotExported, fieldName))
	// 	} else if field.Type != differ.diffType() {
	// 		errs = append(errs, fmt.Errorf("%w: %s has type %s, differ supports type %s", ErrFieldWrongType, fieldName, field.Type, differ.diffType()))
	// 	}
	// }
	// for methodName, differ := range s.Methods {
	// 	if method, ok := tType.MethodByName(methodName); !ok {
	// 		errs = append(errs, fmt.Errorf("%w: %s", ErrMethodNotFound, methodName))
	// 	} else if !method.IsExported() {
	// 		errs = append(errs, fmt.Errorf("%w: %s", ErrMethodNotExported, methodName))
	// 	} else if method.Type != differ.diffType() {
	// 		errs = append(errs, fmt.Errorf("%w: %s() has type %s, differ supports type %s", ErrMethodWrongType, methodName, method.Type, differ.diffType()))
	// 	}
	// }

	return errors.Join(errs...)
}
