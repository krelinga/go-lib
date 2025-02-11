package pod

func WrapSliceComp[V comparable](in []V) wrappedSliceComp[V] {
	return wrappedSliceComp[V](in)
}

type wrappedSliceComp[V comparable] []V

func (wsc wrappedSliceComp[V]) InternalValidate(reporter ErrorReporter) {
	// TODO: Implement
}

func (wsc wrappedSliceComp[V]) InternalDiff(rhs POD, reporter DiffReporter) {
	// TODO: implement
}

func (wsc wrappedSliceComp[V]) InternalDeepCopyTo(out POD) {
	// TODO: Implement
}

func WrapSlicePOD[V POD](in []V) wrappedSlicePOD[V] {
	return wrappedSlicePOD[V](in)
}

type wrappedSlicePOD[V POD] []V

func (wsp wrappedSlicePOD[V]) InternalValidate(reporter ErrorReporter) {
	// TODO: Implement
}

func (wsp wrappedSlicePOD[V]) InternalDiff(rhs POD, reporter DiffReporter) {
	// TODO: Implement
}

func (wsp wrappedSlicePOD[V]) InternalDeepCopyTo(out POD) {
	// TODO: Implement
}