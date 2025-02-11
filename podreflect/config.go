package podreflect

import "reflect"

type fieldConfigSetter func(*fieldConfig)

func WithExtraValidator[T any](val func(T, ErrorReporter)) fieldConfigSetter {
	return func(fc *fieldConfig) {
		if fc.fieldType != reflect.TypeFor[T]() {
			panic("WithExtraValidator: type mismatch")
		}
		fc.extraValidator = func(v any, reporter ErrorReporter) {
			val(v.(T), reporter)
		}
	}
}

func FieldConfig[T any](fieldName string, setters ...fieldConfigSetter) fieldConfig {
	fc := fieldConfig{
		fieldName: fieldName,
		fieldType: reflect.TypeFor[T](),
	}
	for _, setter := range setters {
		setter(&fc)
	}
	return fc
}

func StructConfig[T any](fieldConfigs ...fieldConfig) structConfig {
	sc := structConfig{
		structType: reflect.TypeFor[T](),
	}
	sc.FieldConfigs = append(sc.FieldConfigs, fieldConfigs...)
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
