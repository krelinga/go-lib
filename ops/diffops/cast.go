package diffops

func CastRhs[T any](lhs T, rhs any, s Sink, fn func(rhs T)) {
	tRhs, ok := rhs.(T)
	if !ok {
		s.TypeDiff(lhs, rhs)
		return
	}
	fn(tRhs)
}
