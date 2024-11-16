package kiter

import (
	"iter"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGrouped(t *testing.T) {
	tests := []struct {
		name string
		in   []KV[int, string]
		want map[int][]string
	}{
		{
			name: "group by key",
			in: []KV[int, string]{
				{1, "a"},
				{2, "b"},
				{1, "c"},
			},
			want: map[int][]string{
				1: {"a", "c"},
				2: {"b"},
			},
		},
		{
			name: "empty input",
			in:   []KV[int, string]{},
			want: map[int][]string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			grouped := Grouped(FromKVSlice(tt.in))
			got := ToMap(Map2(grouped, func(k int, v iter.Seq[string]) (int, []string) {
				return k, ToSlice(v)
			}))

			assert.Equal(t, tt.want, got)
		})
	}
}
