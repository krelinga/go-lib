package podreflect

import (
	"reflect"
)

type ErrorReporter interface {
	Err(err error)
}

func CollectErrors(in POD, reporter ErrorReporter) {
	collectErrorsImpl(reflect.ValueOf(in), reporter)
}

type LocalValidator interface {
	PODLocalValidate(ErrorReporter)
}

func collectErrorsImpl(in reflect.Value, reporter ErrorReporter) {
	switch in.Kind() {
	case reflect.Struct:
		if in.Type().Implements(reflect.TypeFor[LocalValidator]()) {
			in.Interface().(LocalValidator).PODLocalValidate(reporter)
		}
		for i := 0; i < in.NumField(); i++ {
			collectErrorsImpl(in.Field(i), reporter)
		}
	case reflect.Ptr:
		if in.IsNil() {
			return
		}
		if in.Type().Implements(reflect.TypeFor[LocalValidator]()) {
			in.Interface().(LocalValidator).PODLocalValidate(reporter)
		}
		collectErrorsImpl(in.Elem(), reporter)
	case reflect.Interface:
		if in.IsNil() {
			return
		}
		collectErrorsImpl(in.Elem(), reporter)
	default:
		return
	}
}
