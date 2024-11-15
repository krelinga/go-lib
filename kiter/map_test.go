package kiter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMap(t *testing.T) {
	tests := []struct {
		name      string
		input     []int
		expected  []int
		stopEarly bool
		stopValue int
	}{
		{
			name:     "Normal case",
			input:    []int{1, 2, 3, 4},
			expected: []int{2, 4, 6, 8},
		},
		{
			name:     "Empty input",
			input:    []int{},
			expected: []int{},
		},
		{
			name:      "Stop early",
			input:     []int{1, 2, 3, 4},
			expected:  []int{2, 4},
			stopEarly: true,
			stopValue: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			in := FromSlice(tt.input)
			fn := func(x int) int { return x * 2 }
			out := Map(in, fn)
			result := []int{}
			for v := range out {
				result = append(result, v)
				if tt.stopEarly && v == tt.stopValue {
					break
				}
			}
			assert.Equal(t, tt.expected, result)
		})
	}
}
