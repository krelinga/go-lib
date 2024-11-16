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
			got := ToSliceAppend(tt.input, seq)

			assert.Equal(t, tt.want, got)
		})
	}
}
func TestFromKVSlice(t *testing.T) {
	tests := []struct {
		name      string
		input     []KV[int, string]
		want      []KV[int, string]
		stopEarly bool
		stopKey   int
	}{
		{
			name:  "empty slice",
			input: []KV[int, string]{},
			want:  []KV[int, string]{},
		},
		{
			name:  "single element",
			input: []KV[int, string]{{K: 1, V: "one"}},
			want:  []KV[int, string]{{K: 1, V: "one"}},
		},
		{
			name:  "multiple elements",
			input: []KV[int, string]{{K: 1, V: "one"}, {K: 2, V: "two"}, {K: 3, V: "three"}},
			want:  []KV[int, string]{{K: 1, V: "one"}, {K: 2, V: "two"}, {K: 3, V: "three"}},
		},
		{
			name:      "stop early",
			input:     []KV[int, string]{{K: 1, V: "one"}, {K: 2, V: "two"}, {K: 3, V: "three"}},
			want:      []KV[int, string]{{K: 1, V: "one"}, {K: 2, V: "two"}},
			stopEarly: true,
			stopKey:   2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := []KV[int, string]{}
			seq := FromKVSlice(tt.input)
			for k, v := range seq {
				got = append(got, KV[int, string]{K: k, V: v})
				if tt.stopEarly && k == tt.stopKey {
					break
				}
			}

			assert.Equal(t, tt.want, got)
		})
	}
}
func TestToKVSlice(t *testing.T) {
	tests := []struct {
		name  string
		input []KV[int, string]
		want  []KV[int, string]
	}{
		{
			name:  "empty sequence",
			input: []KV[int, string]{},
			want:  []KV[int, string]{},
		},
		{
			name:  "single element",
			input: []KV[int, string]{{K: 1, V: "one"}},
			want:  []KV[int, string]{{K: 1, V: "one"}},
		},
		{
			name:  "multiple elements",
			input: []KV[int, string]{{K: 1, V: "one"}, {K: 2, V: "two"}, {K: 3, V: "three"}},
			want:  []KV[int, string]{{K: 1, V: "one"}, {K: 2, V: "two"}, {K: 3, V: "three"}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			seq := FromKVSlice(tt.input)
			got := ToKVSlice(seq)

			assert.Equal(t, tt.want, got)
		})
	}
}
func TestAppendToKVSlice(t *testing.T) {
	tests := []struct {
		name  string
		input []KV[int, string]
		seq   []KV[int, string]
		want  []KV[int, string]
	}{
		{
			name:  "append to empty slice",
			input: []KV[int, string]{},
			seq:   []KV[int, string]{{K: 1, V: "one"}, {K: 2, V: "two"}},
			want:  []KV[int, string]{{K: 1, V: "one"}, {K: 2, V: "two"}},
		},
		{
			name:  "append to non-empty slice",
			input: []KV[int, string]{{K: 1, V: "one"}},
			seq:   []KV[int, string]{{K: 2, V: "two"}, {K: 3, V: "three"}},
			want:  []KV[int, string]{{K: 1, V: "one"}, {K: 2, V: "two"}, {K: 3, V: "three"}},
		},
		{
			name:  "append empty sequence",
			input: []KV[int, string]{{K: 1, V: "one"}},
			seq:   []KV[int, string]{},
			want:  []KV[int, string]{{K: 1, V: "one"}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			seq := FromKVSlice(tt.seq)
			got := AppendToKVSlice(tt.input, seq)

			assert.Equal(t, tt.want, got)
		})
	}
}
