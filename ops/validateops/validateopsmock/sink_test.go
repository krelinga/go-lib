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
		wantWantMore bool
	}{
		{
			name: "no errors",
			setup: func(s *validateopsmock.Sink) {},
			want: nil,
			wantWantMore: true,
		},
		{
			name: "single root error",
			setup: func(s *validateopsmock.Sink) {
				s.Error(wantError)
			},
			want: []validateopsmock.Entry{
				{Err: wantError},
			},
			wantWantMore: true,
		},
		{
			name: "single field error",
			setup: func(s *validateopsmock.Sink) {
				s.Field("foo").Error(wantError)
			},
			want: []validateopsmock.Entry{
				{Context: "foo", Err: wantError},
			},
			wantWantMore: true,
		},
		{
			name: "single key error",
			setup: func(s *validateopsmock.Sink) {
				s.Key(10).Error(wantError)
			},
			want: []validateopsmock.Entry{
				{Context: "[10]", Err: wantError},
			},
			wantWantMore: true,
		},
		{
			name: "single note error",
			setup: func(s *validateopsmock.Sink) {
				s.Note("note").Error(wantError)
			},
			want: []validateopsmock.Entry{
				{Context: "(note)", Err: wantError},
			},
			wantWantMore: true,
		},
		{
			name: "errors above max are ignored",
			setup: func(s *validateopsmock.Sink) {
				s.MaxErrors = 1
				s.Error(wantError)
				s.Error(wantError)
			},
			want: []validateopsmock.Entry{
				{Err: wantError},
			},
			wantWantMore: false,
		},
		{
			name: "WantMore() becomes false after limit",
			setup: func(s *validateopsmock.Sink) {
				s.MaxErrors = 1
				s.Error(wantError)
			},
			want: []validateopsmock.Entry{
				{Err: wantError},
			},
			wantWantMore: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &validateopsmock.Sink{}
			tt.setup(s)
			assert.Equal(t, tt.want, s.Errors, "unexpected errors")
			assert.Equal(t, tt.wantWantMore, s.WantMore(), "unexpected WantMore()")
		})
	}
}