package base

// Any type that satisfies the Validator interface can be used with the ValidOrPanic function.
// Implementations of the Validator interface must satisfy the following contract:
// - The Validate method must return an error if the value is not valid.
// - The Validate method must not mutate the value.
// - The Validate method must not panic.
type Validator interface {
	Validate() error
}

// ValidOrPanic panics if the input value is not valid.
// The input value is returned, so ValidOrPanic can be used in a single line with another access to the input value.
func ValidOrPanic[V Validator](in V) V {
	if err := in.Validate(); err != nil {
		panic(err)
	}
	return in
}
