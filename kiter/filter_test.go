package kiter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilter(t *testing.T) {
	tests := []struct {
		name        string
		input       []int
		filterFunc  func(int) bool
		expectedOut []int
	}{
		{
			name:        "All pass",
			input:       []int{1, 2, 3, 4, 5},
			filterFunc:  func(v int) bool { return v%2 == 0 || v%2 != 0 },
			expectedOut: []int{1, 2, 3, 4, 5},
		},
		{
			name:        "None pass",
			input:       []int{1, 3, 5},
			filterFunc:  func(v int) bool { return v%2 == 0 },
			expectedOut: []int{},
		},
		{
			name:        "Some pass",
			input:       []int{1, 2, 3, 4, 5},
			filterFunc:  func(v int) bool { return v%2 == 0 },
			expectedOut: []int{2, 4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			in := FromSlice(tt.input)
			out := Filter(in, tt.filterFunc)
			result := ToSlice(out)
			assert.Equal(t, tt.expectedOut, result)
		})
	}
}
func TestFilter2(t *testing.T) {
	tests := []struct {
		name        string
		input       []KV[int, string]
		filterFunc  func(int, string) bool
		expectedOut []KV[int, string]
	}{
		{
			name: "All pass",
			input: []KV[int, string]{
				{1, "one"},
				{2, "two"},
				{3, "three"},
			},
			filterFunc: func(k int, v string) bool { return true },
			expectedOut: []KV[int, string]{
				{1, "one"},
				{2, "two"},
				{3, "three"},
			},
		},
		{
			name: "None pass",
			input: []KV[int, string]{
				{1, "one"},
				{2, "two"},
				{3, "three"},
			},
			filterFunc: func(k int, v string) bool { return false },
			expectedOut: []KV[int, string]{},
		},
		{
			name: "Some pass",
			input: []KV[int, string]{
				{1, "one"},
				{2, "two"},
				{3, "three"},
			},
			filterFunc: func(k int, v string) bool { return k%2 == 0 },
			expectedOut: []KV[int, string]{
				{2, "two"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			in := FromKVSlice(tt.input)
			out := Filter2(in, tt.filterFunc)
			result := ToKVSlice(out)
			assert.Equal(t, tt.expectedOut, result)
		})
	}
}

