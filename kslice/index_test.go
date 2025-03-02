package kslice_test

import (
	"testing"

	"github.com/krelinga/go-lib/kslice"
)

func TestIndex(t *testing.T) {
	tests := []struct {
		name string
		index kslice.Index
		want error
	}{
		{
			name: "valid index",
			index: 0,
			want: nil,
		},
		{
			name: "negative index",
			index: -1,
			want: kslice.ErrNegativeIndex,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.index.Validate(); got != tt.want {
				t.Errorf("Index.Validate() = %v, want %v", got, tt.want)
			}
		})
	}
}