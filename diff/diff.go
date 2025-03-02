package diff

func Diff[T any](lhs, rhs DefaultDifferer[T]) Results {
	return nil // TODO
}

func DiffUsing[T any](lhs, rhs T, differ TypedDiffer[T]) Results {
	return nil // TODO
}

func DiffAny[T any](lhs, rhs T, differ AnyDiffer) (Results, error) {
	return nil, nil // TODO
}
