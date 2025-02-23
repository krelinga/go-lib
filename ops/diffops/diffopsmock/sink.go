package diffopsmock

import (
	"fmt"

	"github.com/krelinga/go-lib/ops/diffops"
)

type Pair struct {
	Lhs, Rhs any
}

type Diff struct {
	Path string

	TypeDiff  *Pair
	ValueDiff *Pair
	Extra     any
	Missing   any
}

type Sink struct {
	MaxDiffs int
	Diffs    []Diff
}

type subSink struct {
	parent *Sink
	path   string
}

func (s *Sink) TypeDiff(lhs, rhs any) {
	s.Diffs = append(s.Diffs, Diff{
		TypeDiff: &Pair{lhs, rhs},
	})
}

func (s *Sink) ValueDiff(lhs, rhs any) {
	s.Diffs = append(s.Diffs, Diff{
		ValueDiff: &Pair{lhs, rhs},
	})
}

func (s *Sink) Extra(lhs any) {
	s.Diffs = append(s.Diffs, Diff{
		Extra: lhs,
	})
}

func (s *Sink) Missing(lhs any) {
	s.Diffs = append(s.Diffs, Diff{
		Missing: lhs,
	})
}

func (s *Sink) Field(name string) diffops.Sink {
	return &subSink{
		parent: s,
		path:   name,
	}
}

func (s *Sink) Key(key any) diffops.Sink {
	return &subSink{
		parent: s,
		path:   fmt.Sprintf("[%v]", key),
	}
}

func (s *Sink) WantMore() bool {
	return s.MaxDiffs == 0 || len(s.Diffs) < s.MaxDiffs
}

func (ss *subSink) TypeDiff(lhs, rhs any) {
	ss.parent.Diffs = append(ss.parent.Diffs, Diff{
		Path:     ss.path,
		TypeDiff: &Pair{lhs, rhs},
	})
}

func (ss *subSink) ValueDiff(lhs, rhs any) {
	ss.parent.Diffs = append(ss.parent.Diffs, Diff{
		Path:      ss.path,
		ValueDiff: &Pair{lhs, rhs},
	})
}

func (ss *subSink) Extra(rhs any) {
	ss.parent.Diffs = append(ss.parent.Diffs, Diff{
		Path:  ss.path,
		Extra: rhs,
	})
}

func (ss *subSink) Missing(lhs any) {
	ss.parent.Diffs = append(ss.parent.Diffs, Diff{
		Path:    ss.path,
		Missing: lhs,
	})
}

func (ss *subSink) Field(name string) diffops.Sink {
	return &subSink{
		parent: ss.parent,
		path:   fmt.Sprintf("%s.%s", ss.path, name),
	}
}

func (ss *subSink) Key(key any) diffops.Sink {
	return &subSink{
		parent: ss.parent,
		path:   fmt.Sprintf("%s[%v]", ss.path, key),
	}
}

func (ss *subSink) WantMore() bool {
	return ss.parent.WantMore()
}
