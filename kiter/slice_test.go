package kiter

import (
	"testing"
)

func TestFromSlice(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  []int
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
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []int
			seq := FromSlice(tt.input)
			for v := range seq {
				got = append(got, v)
			}

			if len(got) != len(tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}

			for i := range got {
				if got[i] != tt.want[i] {
					t.Errorf("got %v, want %v", got, tt.want)
					break
				}
			}

			// early exit
			for v := range FromSlice(tt.input) {
				if v == 1 {
					break
				}
			}
		})
	}
}
