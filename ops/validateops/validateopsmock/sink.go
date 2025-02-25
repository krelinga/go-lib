package validateopsmock

import (
	"fmt"

	"github.com/krelinga/go-lib/ops/validateops"
)

type Entry struct {
	Context string
	Err     error
}

type childSink struct {
	parent *Sink
	path   string
}

func (cs *childSink) Error(err error) {
	if !cs.parent.WantMore() {
		return
	}
	cs.parent.Errors = append(cs.parent.Errors, Entry{
		Context: cs.path,
		Err:     err,
	})
}

func (cs *childSink) Field(name string) validateops.Sink {
	return &childSink{
		parent: cs.parent,
		path:   fmt.Sprintf("%s.%s", cs.path, name),
	}
}

func (cs *childSink) Key(key any) validateops.Sink {
	return &childSink{
		parent: cs.parent,
		path:   fmt.Sprintf("%s[%v]", cs.path, key),
	}
}

func (cs *childSink) Note(note string) validateops.Sink {
	return &childSink{
		parent: cs.parent,
		path:   fmt.Sprintf("%s(%s)", cs.path, note),
	}
}

func (cs *childSink) WantMore() bool {
	return cs.parent.WantMore()
}

type Sink struct {
	MaxErrors int
	Errors    []Entry
}

func (s *Sink) Error(err error) {
	if s.WantMore() {
		s.Errors = append(s.Errors, Entry{Err: err})
	}
}

func (s *Sink) Field(name string) validateops.Sink {
	return &childSink{
		parent: s,
		path:   name,
	}
}

func (s *Sink) Key(key any) validateops.Sink {
	return &childSink{
		parent: s,
		path:   fmt.Sprintf("[%v]", key),
	}
}

func (s *Sink) Note(note string) validateops.Sink {
	return &childSink{
		parent: s,
		path:   fmt.Sprintf("(%s)", note),
	}
}

func (s *Sink) WantMore() bool {
	return s.MaxErrors == 0 || len(s.Errors) < s.MaxErrors
}
