package base

// Any type that implements DeepCopier can be used with the DeepCopy function.
// Callers should not use this method directly, and instead use the DeepCopy function at package level.
// Implementations of the DeepCopier interface must satisfy the following contract:
// - The DeepCopy method must return a deep copy of the value.
// - The DeepCopy method must not mutate the value.
// - The DeepCopy method must not panic.
// - The type of DeepCopy's return value must be the same as the type of the input value.
type DeepCopier interface {
	DeepCopy() interface{}
}

// DeepCopy returns a deep copy of the input value.
func DeepCopy[DC DeepCopier](in DC) DC {
	return in.DeepCopy().(DC)
}
