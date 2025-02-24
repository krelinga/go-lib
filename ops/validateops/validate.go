package validateops

import (
	"errors"
	"fmt"
)

func Validate(op ValidateOper) error {
	c := collectorSink{MaxErrors: 1}
	ByMethod[ValidateOper]()(op, &c)
	if len(c.Errors) == 0 {
		return nil
	}
	return c.Errors[0]
}

func ValidateAll(op ValidateOper) error {
	c := collectorSink{}
	ByMethod[ValidateOper]()(op, &c)
	return errors.Join(c.Errors...)
}

type Error struct {
	Context string
	Err     error
}

func (e *Error) Error() string {
	var contextPart string
	if e.Context != "" {
		contextPart = fmt.Sprintf("%s: ", e.Context)
	}
	return fmt.Sprintf("%s%s", contextPart, e.Err.Error())
}

type collectorSink struct {
	// Set to 0 for no max.
	MaxErrors int
	Errors    []error
}

func (s *collectorSink) Error(err error) {
	if s.WantMore() {
		s.Errors = append(s.Errors, &Error{Err: err})
	}
}

func (s *collectorSink) Field(field string) Sink {
	return &childSink{
		Parent:  s,
		Context: field,
	}
}

func (s *collectorSink) Key(key any) Sink {
	return &childSink{
		Parent:  s,
		Context: fmt.Sprintf("[%v]", key),
	}
}

func (s *collectorSink) Note(note string) Sink {
	return &childSink{
		Parent:  s,
		Context: fmt.Sprintf("(%s)", note),
	}
}

func (s *collectorSink) WantMore() bool {
	return s.MaxErrors == 0 || len(s.Errors) < s.MaxErrors
}

type childSink struct {
	Parent  *collectorSink
	Context string
}

func (s *childSink) Error(err error) {
	if s.Parent.WantMore() {
		s.Parent.Errors = append(s.Parent.Errors, &Error{Context: s.Context, Err: err})
	}
}

func (s *childSink) Field(field string) Sink {
	return &childSink{
		Parent:  s.Parent,
		Context: fmt.Sprintf("%s.%s", s.Context, field),
	}
}

func (s *childSink) Key(key any) Sink {
	return &childSink{
		Parent:  s.Parent,
		Context: fmt.Sprintf("%s[%v]", s.Context, key),
	}
}

func (s *childSink) Note(note string) Sink {
	return &childSink{
		Parent:  s.Parent,
		Context: fmt.Sprintf("%s (%s)", s.Context, note),
	}
}

func (s *childSink) WantMore() bool {
	return s.Parent.WantMore()
}
