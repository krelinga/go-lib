package kiter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFromSlice(t *testing.T) {
	tests := []struct {
		name      string
		input     []int
		want      []int
		stopEarly bool
		stopValue int
	}{
		{
			name:  "empty slice",
			input: []int{},
			want:  []int{},
		},
		{
			name:  "single element",
			input: []int{1},
			want:  []int{1},
		},
		{
			name:  "multiple elements",
			input: []int{1, 2, 3, 4, 5},
			want:  []int{1, 2, 3, 4, 5},
		},
		{
			name:      "stop early",
			input:     []int{1, 2, 3, 4},
			want:      []int{1, 2},
			stopEarly: true,
			stopValue: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := []int{}
			seq := FromSlice(tt.input)
			for v := range seq {
				got = append(got, v)
				if tt.stopEarly && v == tt.stopValue {
					break
				}
			}

			assert.Equal(t, tt.want, got)
		})
	}
}

func TestToSlice(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  []int
	}{
		{
			name:  "empty sequence",
			input: []int{},
			want:  []int{},
		},
		{
			name:  "single element",
			input: []int{1},
			want:  []int{1},
		},
		{
			name:  "multiple elements",
			input: []int{1, 2, 3, 4, 5},
			want:  []int{1, 2, 3, 4, 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			seq := FromSlice(tt.input)
			got := ToSlice(seq)

			assert.Equal(t, tt.want, got)
		})
	}
}
func TestAppendToSlice(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		seq   []int
		want  []int
	}{
		{
			name:  "append to empty slice",
			input: []int{},
			seq:   []int{1, 2, 3},
			want:  []int{1, 2, 3},
		},
		{
			name:  "append to non-empty slice",
			input: []int{1, 2},
			seq:   []int{3, 4, 5},
			want:  []int{1, 2, 3, 4, 5},
		},
		{
			name:  "append empty sequence",
			input: []int{1, 2, 3},
			seq:   []int{},
			want:  []int{1, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			seq := FromSlice(tt.seq)
			got := AppendToSlice(tt.input, seq)

			assert.Equal(t, tt.want, got)
		})
	}
}
