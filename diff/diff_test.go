package diff

import "testing"

func TestDiff(t *testing.T) {
	tests := []struct {
		name string
		lhs  int
		rhs  int
		want bool
	}{
		{
			name: "lhs is different than rhs",
			lhs:  1,
			rhs:  2,
			want: true,
		},
		{
			name: "lhs is equal to rhs",
			lhs:  1,
			rhs:  1,
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Diff(tt.lhs, tt.rhs); got != tt.want {
				t.Errorf("Diff() = %v, want %v", got, tt.want)
			}
		})
	}
}