package kslice_test

import (
	"reflect"
	"testing"

	"github.com/krelinga/go-lib/kslice"
)

func FlattenTest(t *testing.T) {
	cases := []struct {
		name string
		in   [][]int
		want []int
	}{
		{
			name: "2 slices",
			in: [][]int{{1, 2}, {3, 4}},
			want: []int{1, 2, 3, 4},
		},
		{
			name: "3 slices",
			in: [][]int{{1, 2}, {3, 4}, {5, 6}},
			want: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name: "empty slices",
			in: [][]int{{}, {}, {}},
			want: []int{},
		},
		{
			name: "mixed slices",
			in: [][]int{{1, 2}, {}, {3, 4}},
			want: []int{1, 2, 3, 4},
		},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			got := kslice.Flatten(tt.in...)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}