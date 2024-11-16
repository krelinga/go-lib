package kiter

import (
	"strings"
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
func TestMap2(t *testing.T) {
	tests := []struct {
		name      string
		input     map[int]string
		expected  map[int]string
		stopEarly bool
		stopKey   int
	}{
		{
			name:     "Normal case",
			input:    map[int]string{1: "a", 2: "b", 3: "c"},
			expected: map[int]string{2: "A", 4: "B", 6: "C"},
		},
		{
			name:     "Empty input",
			input:    map[int]string{},
			expected: map[int]string{},
		},
		{
			name:      "Stop early",
			input:     map[int]string{1: "a", 2: "b", 3: "c"},
			expected:  map[int]string{2: "A", 4: "B"},
			stopEarly: true,
			stopKey:   4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			in := Sorted2(FromMap(tt.input))
			fn := func(k int, v string) (int, string) { return k * 2, strings.ToUpper(v) }
			out := Map2(in, fn)
			result := map[int]string{}
			for k, v := range out {
				result[k] = v
				if tt.stopEarly && k == tt.stopKey {
					break
				}
			}
			assert.Equal(t, tt.expected, result)
		})
	}
}

