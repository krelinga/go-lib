package diffopsmock

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSink(t *testing.T) {
	tests := []struct {
		name string
		sink *Sink
		body func(s *Sink)
		want []Diff
	}{
		{
			name: "TypeDiff",
			sink: &Sink{},
			body: func(s *Sink) {
				s.TypeDiff(1, "1")
			},
			want: []Diff{
				{TypeDiff: &Pair{1, "1"}},
			},
		},
		{
			name: "ValueDiff",
			sink: &Sink{},
			body: func(s *Sink) {
				s.ValueDiff(1, 2)
			},
			want: []Diff{
				{ValueDiff: &Pair{1, 2}},
			},
		},
		{
			name: "Extra",
			sink: &Sink{},
			body: func(s *Sink) {
				s.Extra(1)
			},
			want: []Diff{
				{Extra: 1},
			},
		},
		{
			name: "Missing",
			sink: &Sink{},
			body: func(s *Sink) {
				s.Missing(1)
			},
			want: []Diff{
				{Missing: 1},
			},
		},
		{
			name: "NestedFields",
			sink: &Sink{},
			body: func(s *Sink) {
				s.Field("a").Field("b").ValueDiff(1, 2)
				s.Field("a").Field("c").TypeDiff(3, "4")
				s.Field("d").Extra(5)
				s.Field("e").Missing(6)
			},
			want: []Diff{
				{Path: "a.b", ValueDiff: &Pair{1, 2}},
				{Path: "a.c", TypeDiff: &Pair{3, "4"}},
				{Path: "d", Extra: 5},
				{Path: "e", Missing: 6},
			},
		},
		{
			name: "NestedKeys",
			sink: &Sink{},
			body: func(s *Sink) {
				s.Key("a").Key("b").ValueDiff(1, 2)
				s.Key("a").Key("c").TypeDiff(3, "4")
				s.Key("d").Extra(5)
				s.Key("e").Missing(6)
			},
			want: []Diff{
				{Path: "[a][b]", ValueDiff: &Pair{1, 2}},
				{Path: "[a][c]", TypeDiff: &Pair{3, "4"}},
				{Path: "[d]", Extra: 5},
				{Path: "[e]", Missing: 6},
			},
		},
		// TODO: test WantMore
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.body(tt.sink)
			if !assert.Equal(t, len(tt.want), len(tt.sink.Diffs), "number of diffs") {
				return
			}
			for i, diff := range tt.sink.Diffs {
				assert.Equal(t, tt.want[i], diff, "diff at index %d", i)
			}
		})
	}
}
