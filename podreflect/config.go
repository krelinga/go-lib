package podreflect

import "reflect"

type fieldConfigSetter func(*fieldConfig)

func WithFieldValidator[T any](val func(T, ErrorReporter)) fieldConfigSetter {
	return func(fc *fieldConfig) {
		if fc.fieldType != reflect.TypeFor[T]() {
			panic("WithExtraValidator: type mismatch")
		}
		fc.extraValidator = func(v any, reporter ErrorReporter) {
			val(v.(T), reporter)
		}
	}
}

type structConfigSetter func(*structConfig)

func FieldConfig[T any](fieldName string, setters ...fieldConfigSetter) structConfigSetter {
	return func(sc *structConfig) {
		fc := fieldConfig{
			fieldName: fieldName,
			fieldType: reflect.TypeFor[T](),
		}
		for _, setter := range setters {
			setter(&fc)
		}
		sc.FieldConfigs = append(sc.FieldConfigs, fc)
	}
}

func StructConfig[T any](setters ...structConfigSetter) structConfig {
	sc := structConfig{
		structType: reflect.TypeFor[T](),
	}
	for _, setter := range setters {
		setter(&sc)
	}
	return sc
}

type fieldConfig struct {
	fieldName      string
	fieldType      reflect.Type
	extraValidator func(any, ErrorReporter)
}

type structConfig struct {
	structType   reflect.Type
	FieldConfigs []fieldConfig
}
