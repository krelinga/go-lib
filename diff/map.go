package diff

import "reflect"

type Map[K comparable, V any] struct {
	// TODO: think more about this config ... do I like all of these being separate?
	IgnoreExtra bool
	IgnoreMissing bool
	Explicit bool
	Entries map[K]TypedDiffer[V]
	Methods map[string]AnyDiffer
}

func (m Map[K, V]) typedDiff(lhs, rhs map[K]V) Results {
	return nil // TODO
}

func (m Map[K, V]) anyDiff(lhs, rhs any) (Results, error) {
	return nil, nil // TODO
}

func (m Map[K, V]) Validate() error {
	return nil // TODO
}

func (m Map[K, V]) accepts(t reflect.Type) bool {
	return false // TODO
}