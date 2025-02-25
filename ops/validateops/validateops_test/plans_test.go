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
				in: ErrWrapper{},
				want: nil,
			},
			{
				name: "error",
				in: ErrWrapper{Err: e1},
				want: []validateopsmock.Entry{
					{Err: e1},
				},
			},
			{
				name: "error when no more wanted",
				in: ErrWrapper{Err: e2},
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
				in: &PtrErrWrapper{Err: e1},
				want: []validateopsmock.Entry{
					{Err: e1},
				},
			},
			{
				name: "error with nil pointer receiver",
				in: (*PtrErrWrapper)(nil),
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
				in: 1,
				want: nil,
			},
			{
				name: "error",
				in: 0,
				want: []validateopsmock.Entry{
					{Err: validateops.ErrWantNonZero},
				},
			},
			{
				name: "error when no more wanted",
				in: 0,
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
				in: 0,
				want: nil,
			},
			{
				name: "error",
				in: 1,
				want: []validateopsmock.Entry{
					{Err: validateops.ErrWantZero},
				},
			},
			{
				name: "error when no more wanted",
				in: 1,
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
}
