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
			name: "single root error",
			setup: func(s *validateopsmock.Sink) {
				s.Error(wantError)
			},
			want: []validateopsmock.Entry{
				{Err: wantError},
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