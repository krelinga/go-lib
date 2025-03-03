package diff

import (
	"reflect"
	"testing"
)

type runner interface {
	run(t *testing.T)
}

type testDiffCase[T any] struct {
	name    string
	lhs     T
	rhs     T
	want    bool
}

func (c testDiffCase[T]) run(t *testing.T) {
	t.Run(c.name, func(t *testing.T) {
		got := Diff(c.lhs, c.rhs)
		if got != c.want {
			t.Errorf("Diff() = %v, want %v", got, c.want)
		}
	})
}

func ptr[T any](v T) *T {
	return &v
}

type getter interface {
	Get() int
}

type ptrGetter interface {
	GetPtr() int
}

type myInt int

func (m myInt) Get() int {
	return int(m)
}

func (m *myInt) GetPtr() int {
	return int(*m)
}

type compStruct struct {
	Str string
	Int int
}

type nonCompStruct struct {
	Slice []int
	pInt  int
}

type ChildStruct struct {
	Str string
}

type ParentStruct struct {
	ChildStruct
	Int int
}

func isComparable[T any]() bool {
	return reflect.TypeFor[T]().Comparable()
}

func TestDiff(t *testing.T) {
	if !isComparable[compStruct]() {
		t.Fatal("compStruct is not comparable")
	}
	if isComparable[nonCompStruct]() {
		t.Fatal("nonCompStruct is comparable")
	}

	tests := []runner{
		testDiffCase[int]{name: "int not equal", lhs: 1, rhs: 2, want: true},
		testDiffCase[int]{name: "int equal", lhs: 1, rhs: 1, want: false},
		testDiffCase[myInt]{name: "myInt not equal", lhs: 1, rhs: 2, want: true},
		testDiffCase[myInt]{name: "myInt equal", lhs: 1, rhs: 1, want: false},
		testDiffCase[getter]{name: "getter not equal", lhs: myInt(1), rhs: myInt(2), want: true},
		testDiffCase[getter]{name: "getter equal", lhs: myInt(1), rhs: myInt(1), want: false},
		testDiffCase[getter]{name: "gitter nil", lhs: nil, rhs: nil, want: false},
		testDiffCase[getter]{name: "getter one nil", lhs: nil, rhs: myInt(1), want: true},
		testDiffCase[ptrGetter]{name: "ptrGetter not equal", lhs: ptr(myInt(1)), rhs: ptr(myInt(2)), want: true},
		testDiffCase[ptrGetter]{name: "ptrGetter equal", lhs: ptr(myInt(1)), rhs: ptr(myInt(1)), want: false},
		testDiffCase[ptrGetter]{name: "ptrGetter nil", lhs: nil, rhs: nil, want: false},
		testDiffCase[ptrGetter]{name: "ptrGetter one nil", lhs: nil, rhs: ptr(myInt(1)), want: true},
		testDiffCase[*int]{name: "int ptr not equal", lhs: ptr(1), rhs: ptr(2), want: true},
		testDiffCase[*int]{name: "int ptr equal", lhs: ptr(1), rhs: ptr(1), want: false},
		testDiffCase[*int]{name: "int ptr nil", lhs: nil, rhs: nil, want: false},
		testDiffCase[*int]{name: "int ptr one nil", lhs: nil, rhs: ptr(1), want: true},
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
		testDiffCase[compStruct]{name: "compStruct not equal", lhs: compStruct{Str: "a", Int: 1}, rhs: compStruct{Str: "b", Int: 2}, want: true},
		testDiffCase[compStruct]{name: "compStruct equal", lhs: compStruct{Str: "a", Int: 1}, rhs: compStruct{Str: "a", Int: 1}, want: false},
		testDiffCase[*compStruct]{name: "compStruct ptr not equal", lhs: ptr(compStruct{Str: "a", Int: 1}), rhs: ptr(compStruct{Str: "b", Int: 2}), want: true},
		testDiffCase[*compStruct]{name: "compStruct ptr equal", lhs: ptr(compStruct{Str: "a", Int: 1}), rhs: ptr(compStruct{Str: "a", Int: 1}), want: false},
		testDiffCase[nonCompStruct]{name: "nonCompStruct not equal", lhs: nonCompStruct{Slice: []int{1, 2}}, rhs: nonCompStruct{Slice: []int{2, 1}}, want: true},
		testDiffCase[nonCompStruct]{name: "nonCompStruct equal", lhs: nonCompStruct{Slice: []int{1, 2}}, rhs: nonCompStruct{Slice: []int{1, 2}}, want: false},
		testDiffCase[nonCompStruct]{name: "nonCompStruct unexported field ignored", lhs: nonCompStruct{Slice: []int{1, 2}, pInt: 1}, rhs: nonCompStruct{Slice: []int{1, 2}, pInt: 2}, want: false},
		testDiffCase[*nonCompStruct]{name: "nonCompStruct ptr not equal", lhs: ptr(nonCompStruct{Slice: []int{1, 2}}), rhs: ptr(nonCompStruct{Slice: []int{2, 1}}), want: true},
		testDiffCase[*nonCompStruct]{name: "nonCompStruct ptr equal", lhs: ptr(nonCompStruct{Slice: []int{1, 2}}), rhs: ptr(nonCompStruct{Slice: []int{1, 2}}), want: false},
		testDiffCase[*nonCompStruct]{name: "nonCompStruct ptr one nil", lhs: nil, rhs: ptr(nonCompStruct{Slice: []int{1, 2}}), want: true},
		testDiffCase[*nonCompStruct]{name: "nonCompStruct ptr nil", lhs: nil, rhs: nil, want: false},
		testDiffCase[[]int]{name: "slice nil", lhs: nil, rhs: nil, want: false},
		testDiffCase[*[]int]{name: "slice ptr nil", lhs: nil, rhs: nil, want: false},
		testDiffCase[[]int]{name: "slice not equal", lhs: []int{1, 2}, rhs: []int{2, 1}, want: true},
		testDiffCase[[]int]{name: "slice equal", lhs: []int{1, 2}, rhs: []int{1, 2}, want: false},
		testDiffCase[[]int]{name: "slice one nil", lhs: nil, rhs: []int{1, 2}, want: true},
		testDiffCase[[]int]{name: "slice differnt lengths", lhs: []int{1, 2}, rhs: []int{1, 2, 3}, want: true},
		testDiffCase[map[int]int]{name: "map nil", lhs: nil, rhs: nil, want: false},
		testDiffCase[*map[int]int]{name: "map ptr nil", lhs: nil, rhs: nil, want: false},
		testDiffCase[map[int]int]{name: "map not equal", lhs: map[int]int{1: 1, 2: 2}, rhs: map[int]int{2: 1, 1: 2}, want: true},
		testDiffCase[map[int]int]{name: "map equal", lhs: map[int]int{1: 1, 2: 2}, rhs: map[int]int{1: 1, 2: 2}, want: false},
		testDiffCase[map[int]int]{name: "map one nil", lhs: nil, rhs: map[int]int{1: 1, 2: 2}, want: true},
		testDiffCase[map[int]int]{name: "map different lengths", lhs: map[int]int{1: 1, 2: 2}, rhs: map[int]int{1: 1, 2: 2, 3: 3}, want: true},
		testDiffCase[ParentStruct]{
			name: "ChildStruct not equal",
			lhs:  ParentStruct{ChildStruct: ChildStruct{Str: "a"}, Int: 1},
			rhs:  ParentStruct{ChildStruct: ChildStruct{Str: "b"}, Int: 1},
			want: true,
		},
		testDiffCase[ParentStruct]{
			name: "ChildStruct equal",
			lhs:  ParentStruct{ChildStruct: ChildStruct{Str: "a"}, Int: 1},
			rhs:  ParentStruct{ChildStruct: ChildStruct{Str: "a"}, Int: 1},
			want: false,
		},
	}
	for _, tt := range tests {
		tt.run(t)
	}
}
