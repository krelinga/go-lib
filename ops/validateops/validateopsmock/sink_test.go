package validateopsmock_test

import (
	"errors"
	"testing"

	"github.com/krelinga/go-lib/ops/validateops/validateopsmock"
	"github.com/stretchr/testify/assert"
)

func TestSink(t *testing.T) {
	wantError := errors.New("want error")
	tests := []struct {
		name string
		setup func(*validateopsmock.Sink)
		want []validateopsmock.Entry
	}{
		{
			name: "no errors",
			setup: func(s *validateopsmock.Sink) {},
			want: nil,
		},
		{
			name: "single root error",
			setup: func(s *validateopsmock.Sink) {
				s.Error(wantError)
			},
			want: []validateopsmock.Entry{
				{Err: wantError},
			},
		},
		{
			name: "single field error",
			setup: func(s *validateopsmock.Sink) {
				s.Field("foo").Error(wantError)
			},
			want: []validateopsmock.Entry{
				{Context: "foo", Err: wantError},
			},
		},
		{
			name: "single key error",
			setup: func(s *validateopsmock.Sink) {
				s.Key(10).Error(wantError)
			},
			want: []validateopsmock.Entry{
				{Context: "[10]", Err: wantError},
			},
		},
		{
			name: "single note error",
			setup: func(s *validateopsmock.Sink) {
				s.Note("note").Error(wantError)
			},
			want: []validateopsmock.Entry{
				{Context: "(note)", Err: wantError},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &validateopsmock.Sink{}
			tt.setup(s)
			assert.Equal(t, tt.want, s.Errors, "unexpected errors")
		})
	}
}