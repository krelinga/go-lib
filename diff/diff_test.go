package diff

import "testing"

type runner interface {
	run(t *testing.T)
}

type testDiffCase[T comparable] struct {
	name string
	lhs  T
	rhs  T
	want bool
}

func (c testDiffCase[T]) run(t *testing.T) {
	t.Helper()
	if got := Diff(c.lhs, c.rhs); got != c.want {
		t.Errorf("Diff() = %v, want %v", got, c.want)
	}
}

func TestDiff(t *testing.T) {
	tests := []runner{
		testDiffCase[int]{name: "int not equal", lhs: 1, rhs: 2, want: true},
		testDiffCase[int]{name: "int equal", lhs: 1, rhs: 1, want: false},
		testDiffCase[string]{name: "string not equal", lhs: "a", rhs: "b", want: true},
		testDiffCase[string]{name: "string equal", lhs: "a", rhs: "a", want: false},
		testDiffCase[float64]{name: "float64 not equal", lhs: 1.0, rhs: 2.0, want: true},
		testDiffCase[float64]{name: "float64 equal", lhs: 1.0, rhs: 1.0, want: false},
	}
	for _, tt := range tests {
		tt.run(t)
	}
}
