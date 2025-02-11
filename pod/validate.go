package pod

type ErrorReporter interface {
	Error(err error)

	ChildField(name string) ErrorReporter
	ChildKey(key interface{}) ErrorReporter
}

func Validate(p POD) error {
	return nil // TODO: Implement
}