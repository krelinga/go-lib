package kiter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFromMap(t *testing.T) {
	tests := []struct {
		name      string
		input     map[string]int
		expected  map[string]int
		stopEarly bool
		stopKey   string
	}{
		{
			name:     "empty map",
			input:    map[string]int{},
			expected: map[string]int{},
		},
		{
			name:     "single element",
			input:    map[string]int{"a": 1},
			expected: map[string]int{"a": 1},
		},
		{
			name:     "multiple elements",
			input:    map[string]int{"a": 1, "b": 2, "c": 3},
			expected: map[string]int{"a": 1, "b": 2, "c": 3},
		},
		{
			name:      "stop early",
			input:     map[string]int{"a": 1, "b": 2, "c": 3},
			expected:  map[string]int{"a": 1, "b": 2},
			stopEarly: true,
			stopKey:   "b",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			seq := Sorted2(FromMap(tt.input))
			result := make(map[string]int)
			for k, v := range seq {
				result[k] = v
				if tt.stopEarly && k == tt.stopKey {
					break
				}
			}

			assert.Equal(t, tt.expected, result)
		})
	}
}
func TestToMap(t *testing.T) {
	tests := []struct {
		name     string
		input    map[string]int
		expected map[string]int
	}{
		{
			name:     "empty map",
			input:    map[string]int{},
			expected: map[string]int{},
		},
		{
			name:     "single element",
			input:    map[string]int{"a": 1},
			expected: map[string]int{"a": 1},
		},
		{
			name:     "multiple elements",
			input:    map[string]int{"a": 1, "b": 2, "c": 3},
			expected: map[string]int{"a": 1, "b": 2, "c": 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			seq := FromMap(tt.input)
			result := ToMap(seq)
			assert.Equal(t, tt.expected, result)
		})
	}
}
func TestInsertIntoMap(t *testing.T) {
	tests := []struct {
		name     string
		initial  map[string]int
		input    map[string]int
		expected map[string]int
	}{
		{
			name:     "empty initial and input map",
			initial:  map[string]int{},
			input:    map[string]int{},
			expected: map[string]int{},
		},
		{
			name:     "empty initial map",
			initial:  map[string]int{},
			input:    map[string]int{"a": 1},
			expected: map[string]int{"a": 1},
		},
		{
			name:     "empty input map",
			initial:  map[string]int{"a": 1},
			input:    map[string]int{},
			expected: map[string]int{"a": 1},
		},
		{
			name:     "non-empty initial and input map",
			initial:  map[string]int{"a": 1},
			input:    map[string]int{"b": 2},
			expected: map[string]int{"a": 1, "b": 2},
		},
		{
			name:     "overlapping keys",
			initial:  map[string]int{"a": 1},
			input:    map[string]int{"a": 2, "b": 3},
			expected: map[string]int{"a": 2, "b": 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			seq := FromMap(tt.input)
			InsertIntoMap(tt.initial, seq)
			assert.Equal(t, tt.expected, tt.initial)
		})
	}
}
func TestFromMapKeys(t *testing.T) {
	tests := []struct {
		name      string
		input     map[string]int
		expected  []string
		stopEarly bool
		stopKey   string
	}{
		{
			name:     "empty map",
			input:    map[string]int{},
			expected: []string{},
		},
		{
			name:     "single element",
			input:    map[string]int{"a": 1},
			expected: []string{"a"},
		},
		{
			name:     "multiple elements",
			input:    map[string]int{"a": 1, "b": 2, "c": 3},
			expected: []string{"a", "b", "c"},
		},
		{
			name:      "stop early",
			input:     map[string]int{"a": 1, "b": 2, "c": 3},
			expected:  []string{"a", "b"},
			stopEarly: true,
			stopKey:   "b",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			seq := Sorted(FromMapKeys(tt.input))
			result := []string{}
			for k := range seq {
				result = append(result, k)
				if tt.stopEarly && k == tt.stopKey {
					break
				}
			}

			assert.ElementsMatch(t, tt.expected, result)
		})
	}
}
