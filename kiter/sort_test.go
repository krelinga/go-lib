package kiter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSorted2(t *testing.T) {
	tests := []struct {
		name string
		in   []KV[int, string]
		want []KV[int, string]
	}{
		{
			name: "sorted input",
			in:   []KV[int, string]{{1, "a"}, {2, "b"}, {3, "c"}},
			want: []KV[int, string]{{1, "a"}, {2, "b"}, {3, "c"}},
		},
		{
			name: "unsorted input",
			in:   []KV[int, string]{{3, "c"}, {1, "a"}, {2, "b"}},
			want: []KV[int, string]{{1, "a"}, {2, "b"}, {3, "c"}},
		},
		{
			name: "empty input",
			in:   []KV[int, string]{},
			want: []KV[int, string]{},
		},
		{
			name: "single element",
			in:   []KV[int, string]{{1, "a"}},
			want: []KV[int, string]{{1, "a"}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			in := FromSlice(tt.in)
			got := ToSlice(ToPairs(Sorted2(FromPairs(in))))
			assert.Equal(t, tt.want, got)
		})
	}
}
