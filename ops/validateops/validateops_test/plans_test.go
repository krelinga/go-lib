package validateops_test

import (
	"errors"
	"testing"

	"github.com/krelinga/go-lib/ops/validateops"
	"github.com/krelinga/go-lib/ops/validateops/validateopsmock"
	"github.com/stretchr/testify/assert"
)

func TestPlans(t *testing.T) {
	e1 := errors.New("e1")
	e2 := errors.New("e2")
	t.Run("ByMethod", func(t *testing.T) {
		tests := []struct {
			name     string
			in       validateops.ValidateOper
			sinkInit func(*validateopsmock.Sink)
			want     []validateopsmock.Entry
		}{
			{
				name: "no error",
				in:   ErrWrapper{},
				want: nil,
			},
			{
				name: "error",
				in:   ErrWrapper{Err: e1},
				want: []validateopsmock.Entry{
					{Err: e1},
				},
			},
			{
				name: "error when no more wanted",
				in:   ErrWrapper{Err: e2},
				sinkInit: func(s *validateopsmock.Sink) {
					s.MaxErrors = 1
					s.Error(e1)
				},
				want: []validateopsmock.Entry{
					{Err: e1},
				},
			},
			{
				name: "error with pointer receiver",
				in:   &PtrErrWrapper{Err: e1},
				want: []validateopsmock.Entry{
					{Err: e1},
				},
			},
			{
				name: "error with nil pointer receiver",
				in:   (*PtrErrWrapper)(nil),
				want: nil,
			},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				s := &validateopsmock.Sink{}
				if tt.sinkInit != nil {
					tt.sinkInit(s)
				}
				validateops.ByMethod[validateops.ValidateOper]()(tt.in, s)
				assert.Equal(t, tt.want, s.Errors)
			})
		}
	})

	t.Run("NonZero", func(t *testing.T) {
		tests := []struct {
			name     string
			in       int
			sinkInit func(*validateopsmock.Sink)
			want     []validateopsmock.Entry
		}{
			{
				name: "no error",
				in:   1,
				want: nil,
			},
			{
				name: "error",
				in:   0,
				want: []validateopsmock.Entry{
					{Err: validateops.ErrWantNonZero},
				},
			},
			{
				name: "error when no more wanted",
				in:   0,
				sinkInit: func(s *validateopsmock.Sink) {
					s.MaxErrors = 1
					s.Error(e1)
				},
				want: []validateopsmock.Entry{
					{Err: e1},
				},
			},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				s := &validateopsmock.Sink{}
				if tt.sinkInit != nil {
					tt.sinkInit(s)
				}
				validateops.NonZero[int]()(tt.in, s)
				assert.Equal(t, tt.want, s.Errors)
			})
		}
	})

	t.Run("Zero", func(t *testing.T) {
		tests := []struct {
			name     string
			in       int
			sinkInit func(*validateopsmock.Sink)
			want     []validateopsmock.Entry
		}{
			{
				name: "no error",
				in:   0,
				want: nil,
			},
			{
				name: "error",
				in:   1,
				want: []validateopsmock.Entry{
					{Err: validateops.ErrWantZero},
				},
			},
			{
				name: "error when no more wanted",
				in:   1,
				sinkInit: func(s *validateopsmock.Sink) {
					s.MaxErrors = 1
					s.Error(e1)
				},
				want: []validateopsmock.Entry{
					{Err: e1},
				},
			},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				s := &validateopsmock.Sink{}
				if tt.sinkInit != nil {
					tt.sinkInit(s)
				}
				validateops.Zero[int]()(tt.in, s)
				assert.Equal(t, tt.want, s.Errors)
			})
		}
	})

	t.Run("AllOf", func(t *testing.T) {
		greaterThanTen := func(in int, sink validateops.Sink) {
			if in <= 10 {
				sink.Error(e1)
			}
		}
		greaterThanTwenty := func(in int, sink validateops.Sink) {
			if in <= 20 {
				sink.Error(e2)
			}
		}
		tests := []struct {
			name     string
			in       int
			plans    []validateops.Plan[int]
			sinkInit func(*validateopsmock.Sink)
			want     []validateopsmock.Entry
		}{
			{
				name:  "no error",
				in:    100,
				plans: []validateops.Plan[int]{greaterThanTen, greaterThanTwenty},
				want:  nil,
			},
			{
				name:  "one error",
				in:    11,
				plans: []validateops.Plan[int]{greaterThanTen, greaterThanTwenty},
				want: []validateopsmock.Entry{
					{Err: e2},
				},
			},
			{
				name:  "two errors",
				in:    0,
				plans: []validateops.Plan[int]{greaterThanTen, greaterThanTwenty},
				want: []validateopsmock.Entry{
					{Err: e1},
					{Err: e2},
				},
			},
			{
				name:  "only one error recorded when no more wanted",
				in:    0,
				plans: []validateops.Plan[int]{greaterThanTen, greaterThanTwenty},
				sinkInit: func(s *validateopsmock.Sink) {
					s.MaxErrors = 1
				},
				want: []validateopsmock.Entry{
					{Err: e1},
				},
			},
			{
				name:  "no error when no plans",
				in:    0,
				plans: nil,
				want:  nil,
			},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				s := &validateopsmock.Sink{}
				if tt.sinkInit != nil {
					tt.sinkInit(s)
				}
				validateops.AllOf[int](tt.plans...)(tt.in, s)
				assert.Equal(t, tt.want, s.Errors)
			})
		}
	})

	t.Run("SliceOf", func(t *testing.T) {
		valueGreaterThanTen := func(in validateops.KV[int, int], sink validateops.Sink) {
			if in.V <= 10 {
				sink.Error(e1)
			}
		}
		tests := []struct {
			name     string
			in       []int
			sinkInit func(*validateopsmock.Sink)
			want     []validateopsmock.Entry
		}{
			{
				name: "no error",
				in:   []int{11, 12},
				want: nil,
			},
			{
				name: "one error",
				in:   []int{11, 10},
				want: []validateopsmock.Entry{
					{Context: "[1]", Err: e1},
				},
			},
			{
				name: "two errors",
				in:   []int{11, 10, 9},
				want: []validateopsmock.Entry{
					{Context: "[1]", Err: e1},
					{Context: "[2]", Err: e1},
				},
			},
			{
				name: "only one error recorded when no more wanted",
				in:   []int{11, 10, 9},
				sinkInit: func(s *validateopsmock.Sink) {
					s.MaxErrors = 1
				},
				want: []validateopsmock.Entry{
					{Context: "[1]", Err: e1},
				},
			},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				s := &validateopsmock.Sink{}
				if tt.sinkInit != nil {
					tt.sinkInit(s)
				}
				validateops.SliceOf[int](valueGreaterThanTen)(tt.in, s)
				assert.Equal(t, tt.want, s.Errors)
			})
		}
	})

	t.Run("MapOf", func(t *testing.T) {
		valueGreaterThanTen := func(in validateops.KV[int, int], sink validateops.Sink) {
			if in.V <= 10 {
				sink.Error(e1)
			}
		}
		tests := []struct {
			name     string
			in       map[int]int
			sinkInit func(*validateopsmock.Sink)
			want     []validateopsmock.Entry
		}{
			{
				name: "no error",
				in:   map[int]int{1: 11, 2: 12},
				want: nil,
			},
			{
				name: "one error",
				in:   map[int]int{1: 11, 2: 10},
				want: []validateopsmock.Entry{
					{Context: "[2]", Err: e1},
				},
			},
			{
				name: "two errors",
				in:   map[int]int{1: 11, 2: 10, 3: 9},
				want: []validateopsmock.Entry{
					{Context: "[2]", Err: e1},
					{Context: "[3]", Err: e1},
				},
			},
			{
				name: "only one error recorded when no more wanted",
				in:   map[int]int{1: 11, 2: 10, 3: 9},
				sinkInit: func(s *validateopsmock.Sink) {
					// Bypassing the input altogether because we can't control
					// map itration order.
					s.MaxErrors = 1
					s.Error(e2)
				},
				want: []validateopsmock.Entry{
					{Err: e2},
				},
			},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				s := &validateopsmock.Sink{}
				if tt.sinkInit != nil {
					tt.sinkInit(s)
				}
				validateops.MapOf[int, int](valueGreaterThanTen)(tt.in, s)
				assert.Equal(t, tt.want, s.Errors)
			})
		}
	})

	t.Run("Keys", func(t *testing.T) {
		tests := []struct {
			name     string
			in       []int
			sinkInit func(*validateopsmock.Sink)
			want     []validateopsmock.Entry
		}{
			{
				name: "no error",
				in:   []int{1},
				want: nil,
			},
			{
				name: "error",
				in:   []int{0, 1},
				want: []validateopsmock.Entry{
					{Context: "[1](key)", Err: validateops.ErrWantZero},
				},
			},
			{
				name: "error when no more wanted",
				in:   []int{0, 1, 2},
				sinkInit: func(s *validateopsmock.Sink) {
					s.MaxErrors = 1
				},
				want: []validateopsmock.Entry{
					{Context: "[1](key)", Err: validateops.ErrWantZero},
				},
			},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				s := &validateopsmock.Sink{}
				if tt.sinkInit != nil {
					tt.sinkInit(s)
				}
				validateops.SliceOf[int](
					validateops.Keys[int, int](
						validateops.Zero[int]()))(tt.in, s)
				assert.Equal(t, tt.want, s.Errors)
			})
		}
	})
}
