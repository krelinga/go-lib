package diff

func Diff[T comparable](lhs, rhs T) bool {
	return lhs != rhs
}