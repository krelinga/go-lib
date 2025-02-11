package pod

func WrapMapComp[K comparable, V comparable](in map[K]V) wrappedMapComp[K, V] {
	return wrappedMapComp[K, V](in)
}

type wrappedMapComp[K comparable, V comparable] map[K]V

func (wmc wrappedMapComp[K, V]) InternalValidate(reporter ErrorReporter) {
	// TODO: Implement
}

func (wmc wrappedMapComp[K, V]) InternalDiff(rhs POD, reporter DiffReporter) {
	typedRhs, ok := rhs.(wrappedMapComp[K, V])
	if !ok {
		reporter.TypeDiff(wmc, rhs)
		return
	}
	for k, lhsV := range wmc {
		rhsV, ok := typedRhs[k]
		if !ok {
			reporter.ChildKey(k).Missing(lhsV)
			continue
		}
		if lhsV != rhsV {
			reporter.ChildKey(k).ValueDiff(lhsV, rhsV)
		}
	}
	for k, rhsV := range typedRhs {
		if _, ok := wmc[k]; !ok {
			reporter.ChildKey(k).Extra(rhsV)
		}
	}
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
	typedRhs, ok := rhs.(wrappedMapPOD[K, V])
	if !ok {
		reporter.TypeDiff(wmp, rhs)
		return
	}
	for k, lhsV := range wmp {
		rhsV, ok := typedRhs[k]
		if !ok {
			reporter.ChildKey(k).Missing(lhsV)
			continue
		}
		Diff(lhsV, rhsV, reporter.ChildKey(k))
	}
	for k, rhsV := range typedRhs {
		if _, ok := wmp[k]; !ok {
			reporter.ChildKey(k).Extra(rhsV)
		}
	}
}

func (wmp wrappedMapPOD[K, V]) InternalDeepCopyTo(out POD) {
	// TODO: Implement
}
