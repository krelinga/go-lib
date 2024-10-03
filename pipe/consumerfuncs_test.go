package pipe

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToArrayFunc(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "Empty input",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "Single element",
			input:    []int{1},
			expected: []int{1},
		},
		{
			name:     "Multiple elements",
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			in := make(chan int, len(tt.input))
			for _, v := range tt.input {
				in <- v
			}
			close(in)

			out := []int{}
			toArray := ToArrayFunc(in, &out)
			toArray()

			assert.Equal(t, tt.expected, out)
		})
	}
}

func newInt(v int) *int {
	return &v
}

func TestFirstFunc(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		input    []int
		expected *int
	}{
		{
			name:     "Empty input",
			input:    []int{},
			expected: nil,
		},
		{
			name:     "Single element",
			input:    []int{1},
			expected: newInt(1),
		},
		{
			name:     "Multiple elements",
			input:    []int{1, 2, 3, 4, 5},
			expected: newInt(1),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			in := make(chan int, len(tt.input))
			for _, v := range tt.input {
				in <- v
			}
			close(in)

			var result *int
			firstFunc := FirstFunc(in, func(v int) {
				result = &v
			})
			firstFunc()

			if tt.expected == nil {
				assert.Nil(t, result)
			} else {
				assert.NotNil(t, result)
				assert.Equal(t, *tt.expected, *result)
			}
		})
	}
}

func TestLastFunc(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		input    []int
		expected *int
	}{
		{
			name:     "Empty input",
			input:    []int{},
			expected: nil,
		},
		{
			name:     "Single element",
			input:    []int{1},
			expected: newInt(1),
		},
		{
			name:     "Multiple elements",
			input:    []int{1, 2, 3, 4, 5},
			expected: newInt(5),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			in := make(chan int, len(tt.input))
			for _, v := range tt.input {
				in <- v
			}
			close(in)

			var result *int
			lastFunc := LastFunc(in, func(v int) {
				result = &v
			})
			lastFunc()

			if tt.expected == nil {
				assert.Nil(t, result)
			} else {
				assert.NotNil(t, result)
				assert.Equal(t, *tt.expected, *result)
			}
		})
	}
}

func TestDiscardFunc(t *testing.T) {
	t.Parallel()
	empty := make(chan struct{}, 5)
	for i := 0; i < 5; i++ {
		empty <- struct{}{}
	}
	close(empty)
	DiscardFunc(empty)()
}

func TestToMapFunc(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		input    []KV[string, int]
		expected map[string]int
	}{
		{
			name:     "Empty input",
			input:    []KV[string, int]{},
			expected: map[string]int{},
		},
		{
			name: "Single element",
			input: []KV[string, int]{
				{Key: "a", Val: 1},
			},
			expected: map[string]int{"a": 1},
		},
		{
			name: "Multiple elements with unique keys",
			input: []KV[string, int]{
				{Key: "a", Val: 1},
				{Key: "b", Val: 2},
				{Key: "c", Val: 3},
			},
			expected: map[string]int{"a": 1, "b": 2, "c": 3},
		},
		{
			name: "Multiple elements with duplicate keys",
			input: []KV[string, int]{
				{Key: "a", Val: 1},
				{Key: "b", Val: 2},
				{Key: "a", Val: 3},
			},
			expected: map[string]int{"a": 3, "b": 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			in := make(chan KV[string, int], len(tt.input))
			for _, kv := range tt.input {
				in <- kv
			}
			close(in)

			out := make(map[string]int)
			toMap := ToMapFunc(in, &out)
			toMap()

			assert.Equal(t, tt.expected, out)
		})
	}
}
