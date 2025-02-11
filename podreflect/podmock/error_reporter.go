package podmock

type ErrorReporter struct {
	Errs []error
}

func (er *ErrorReporter) Err(err error) {
	er.Errs = append(er.Errs, err)
}