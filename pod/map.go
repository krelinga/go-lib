package pod

func WrapMapComp[K comparable, V comparable](in map[K]V) wrappedMapComp[K, V] {
	return wrappedMapComp[K, V](in)
}

type wrappedMapComp[K comparable, V comparable] map[K]V

func (wmc wrappedMapComp[K, V]) InternalValidate(reporter ErrorReporter) {
	// TODO: Implement
}

func (wmc wrappedMapComp[K, V]) InternalDiff(rhs POD, reporter DiffReporter) {
	// TODO: implement
}

func (wmc wrappedMapComp[K, V]) InternalDeepCopyTo(out POD) {
	// TODO: Implement
}

func WrapMapPOD[K comparable, V POD](in map[K]V) wrappedMapPOD[K, V] {
	return wrappedMapPOD[K, V](in)
}

type wrappedMapPOD[K comparable, V POD] map[K]V

func (wmp wrappedMapPOD[K, V]) InternalValidate(reporter ErrorReporter) {
	// TODO: Implement
}

func (wmp wrappedMapPOD[K, V]) InternalDiff(rhs POD, reporter DiffReporter) {
	// TODO: Implement
}

func (wmp wrappedMapPOD[K, V]) InternalDeepCopyTo(out POD) {
	// TODO: Implement
}
