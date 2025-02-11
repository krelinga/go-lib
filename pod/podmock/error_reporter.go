package podmock

import "fmt"

type GotErr struct {
	Path string
	Err error
}

type ErrorReporter struct {
	GotErrs []GotErr
}

func (er *ErrorReporter) Error(err error) {
	er.GotErrs = append(er.GotErrs, GotErr{Err: err})
}

func (er *ErrorReporter) ChildField(name string) *childErrorReporter {
	return &childErrorReporter{parent: er, path: name}
}

func (er *ErrorReporter) ChildKey(key interface{}) *childErrorReporter {
	return &childErrorReporter{parent: er, path: fmt.Sprintf("[%v]", key)}
}

type childErrorReporter struct {
	parent *ErrorReporter
	path string
}

func (cer *childErrorReporter) Error(err error) {
	cer.parent.GotErrs = append(cer.parent.GotErrs, GotErr{Path: cer.path, Err: err})
}

func (cer *childErrorReporter) ChildField(name string) *childErrorReporter {
	return &childErrorReporter{parent: cer.parent, path: fmt.Sprintf("%s.%s", cer.path, name)}
}

func (cer *childErrorReporter) ChildKey(key interface{}) *childErrorReporter {
	return &childErrorReporter{parent: cer.parent, path: fmt.Sprintf("%s[%v]", cer.path, key)}
}