package kiter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToPairs(t *testing.T) {
	tests := []struct {
		name     string
		input    map[int]string
		expected []KV[int, string]
	}{
		{
			name: "three elements",
			input: map[int]string{
				1: "one",
				2: "two",
				3: "three",
			},
			expected: []KV[int, string]{
				{1, "one"},
				{2, "two"},
				{3, "three"},
			},
		},
		{
			name:     "empty map",
			input:    map[int]string{},
			expected: []KV[int, string]{},
		},
		{
			name: "single element",
			input: map[int]string{
				1: "one",
			},
			expected: []KV[int, string]{
				{1, "one"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := FromMap(tt.input)
			result := ToSlice(ToPairs(input))
			assert.ElementsMatch(t, tt.expected, result)
		})
	}
}
func TestFromPairs(t *testing.T) {
	tests := []struct {
		name     string
		input    []KV[int, string]
		expected map[int]string
	}{
		{
			name: "three elements",
			input: []KV[int, string]{
				{1, "one"},
				{2, "two"},
				{3, "three"},
			},
			expected: map[int]string{
				1: "one",
				2: "two",
				3: "three",
			},
		},
		{
			name:     "empty slice",
			input:    []KV[int, string]{},
			expected: map[int]string{},
		},
		{
			name: "single element",
			input: []KV[int, string]{
				{1, "one"},
			},
			expected: map[int]string{
				1: "one",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := FromSlice(tt.input)
			result := ToMap(FromPairs(input))
			assert.Equal(t, tt.expected, result)
		})
	}
}
