package diff

import "testing"

type runner interface {
	run(t *testing.T)
}

type testDiffCase[T any] struct {
	name    string
	lhs     T
	rhs     T
	want    bool
	wantErr error
}

func (c testDiffCase[T]) run(t *testing.T) {
	t.Run(c.name, func(t *testing.T) {
		got, gotErr := Diff(c.lhs, c.rhs)
		if gotErr != c.wantErr {
			t.Errorf("Diff() error = %v, want %v", gotErr, c.wantErr)
		} else if got != c.want {
			t.Errorf("Diff() = %v, want %v", got, c.want)
		}
	})
}

func ptr[T any](v T) *T {
	return &v
}

func TestDiff(t *testing.T) {
	tests := []runner{
		testDiffCase[int]{name: "int not equal", lhs: 1, rhs: 2, want: true},
		testDiffCase[int]{name: "int equal", lhs: 1, rhs: 1, want: false},
		testDiffCase[*int]{name: "int ptr not equal", lhs: ptr(1), rhs: ptr(2), want: true},
		testDiffCase[*int]{name: "int ptr equal", lhs: ptr(1), rhs: ptr(1), want: false},
		testDiffCase[**int]{name: "int ptr ptr not equal", lhs: ptr(ptr(1)), rhs: ptr(ptr(2)), want: true},
		testDiffCase[**int]{name: "int ptr ptr equal", lhs: ptr(ptr(1)), rhs: ptr(ptr(1)), want: false},
		testDiffCase[string]{name: "string not equal", lhs: "a", rhs: "b", want: true},
		testDiffCase[string]{name: "string equal", lhs: "a", rhs: "a", want: false},
		testDiffCase[*string]{name: "string ptr not equal", lhs: ptr("a"), rhs: ptr("b"), want: true},
		testDiffCase[*string]{name: "string ptr equal", lhs: ptr("a"), rhs: ptr("a"), want: false},
		testDiffCase[float64]{name: "float64 not equal", lhs: 1.0, rhs: 2.0, want: true},
		testDiffCase[float64]{name: "float64 equal", lhs: 1.0, rhs: 1.0, want: false},
		testDiffCase[*float64]{name: "float64 ptr not equal", lhs: ptr(1.0), rhs: ptr(2.0), want: true},
		testDiffCase[*float64]{name: "float64 ptr equal", lhs: ptr(1.0), rhs: ptr(1.0), want: false},
		testDiffCase[map[int]int]{name: "map not supported", lhs: nil, rhs: nil, wantErr: ErrUnsupportedType},
		testDiffCase[[]int]{name: "slice not supported", lhs: nil, rhs: nil, wantErr: ErrUnsupportedType},
	}
	for _, tt := range tests {
		tt.run(t)
	}
}
