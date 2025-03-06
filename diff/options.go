package diff

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

type options struct {
	methods []string
}

type Option func(*options)

func WithMethods(methods ...string) Option {
	return func(o *options) {
		o.methods = methods
	}
}

var (
	errEmptyPkgPath      = errors.New("empty package path")
	errAlreadyRegistered = errors.New("type already registered")
	errInvalidMethod     = errors.New("invalid method")
)

func typeKey(t reflect.Type) (string, error) {
	if t.PkgPath() == "" {
		return "", errEmptyPkgPath
	}
	var name string
	if !strings.HasSuffix(t.PkgPath(), "]") {
		name = t.Name()
	} else {
		// This is a generic type, we need to remove the trailing [].
		var level, stop int
		for i := len(t.Name()) - 1; i >= 0; i-- {
			if t.Name()[i] == ']' {
				level++
			}
			if t.Name()[i] == '[' {
				level--
			}
			if level == 0 {
				stop = i
				break
			}
		}
		if level == 0 {
			panic(fmt.Sprintf("unbalanced generic type: %s", t.Name()))
		}
		name = fmt.Sprintf("%s[]", t.Name()[:stop])
	}
	return fmt.Sprintf("%s.%s", t.PkgPath(), name), nil
}

func checkMethodNames(t reflect.Type, methods []string) error {
	var wantMethodInputs int
	switch t.Kind() {
	case reflect.Interface:
		wantMethodInputs = 0
	default:
		wantMethodInputs = 1
	}
	for _, name := range methods {
		if m, found := t.MethodByName(name); !found {
			return fmt.Errorf("%w: %s", errInvalidMethod, name)
		} else if m.Type.NumIn() != wantMethodInputs || m.Type.NumOut() != 1 {
			return fmt.Errorf("%w: %s", errInvalidMethod, name)
		}
	}
	return nil
}

type optionsDb map[string]*options

func (db *optionsDb) register(t reflect.Type, opts ...Option) error {
	key, err := typeKey(t)
	if err != nil {
		return err
	}
	if _, exists := (*db)[key]; exists {
		return errAlreadyRegistered
	}
	optStruct := &options{}
	for _, opt := range opts {
		opt(optStruct)
	}
	if err := checkMethodNames(t, optStruct.methods); err != nil {
		return err
	}
	(*db)[key] = optStruct

	return nil
}

func (db *optionsDb) lookup(t reflect.Type) *options {
	key, err := typeKey(t)
	if err != nil {
		return nil
	}
	return (*db)[key]
}

var globalDb = make(optionsDb)

func Register[T any](opts ...Option) {
	if err := globalDb.register(reflect.TypeFor[T](), opts...); err != nil {
		panic(err)
	}
}
