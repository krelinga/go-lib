package kslice_test

import (
	"testing"

	"github.com/krelinga/go-lib/kslice"
	"github.com/stretchr/testify/assert"
)

func TestIndex(t *testing.T) {
	tests := []struct {
		name       string
		index      kslice.Index
		wantErr       error
		wantString string
	}{
		{
			name:  "zero index",
			index: 0,
			wantErr:  nil,
			wantString: "0",
		},
		{
			name:  "positive index",
			index: 1,
			wantErr:  nil,
			wantString: "1",
		},
		{
			name:  "negative index",
			index: -1,
			wantErr:  kslice.ErrNegativeIndex,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.index.Validate()
			if got != tt.wantErr {
				assert.ErrorIs(t, got, tt.wantErr)
			}
			if tt.wantErr == nil {
				assert.Equal(t, tt.wantString, tt.index.String())
			} else {
				assert.PanicsWithError(t, tt.wantErr.Error(), func() {
					_ = tt.index.String()
				})
			}
		})
	}
}
