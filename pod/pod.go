package pod

type POD interface {
	InternalValidate(reporter ErrorReporter)
	InternalDiff(rhs POD, reporter DiffReporter)
	// TODO: rename this to InternalCopyTo ... not all copies will be deep.
	InternalDeepCopyTo(out POD)
}
