package diff_test

import (
	"reflect"
	"testing"

	"github.com/krelinga/go-lib/datapath"
	"github.com/krelinga/go-lib/diff"
	"github.com/stretchr/testify/assert"
)

type runner interface {
	run(t *testing.T)
}

type testDiffCase[T any] struct {
	name string
	lhs  T
	rhs  T
	want *diff.Result
}

func (c testDiffCase[T]) run(t *testing.T) {
	t.Run(c.name, func(t *testing.T) {
		got := diff.Diff(c.lhs, c.rhs)
		assert.Equalf(t, c.want, got, "Diff() = %v, want %v", got, c.want)
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

type myString string

func (m myString) Get() int {
	return len(m)
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

type mySlice []int

type myMap map[int]string

func isComparable[T any]() bool {
	return reflect.TypeFor[T]().Comparable()
}

func nilPtr[T any]() *T {
	return nil
}

func TestDiff(t *testing.T) {
	if !isComparable[compStruct]() {
		t.Fatal("compStruct is not comparable")
	}
	if isComparable[nonCompStruct]() {
		t.Fatal("nonCompStruct is comparable")
	}

	tests := []runner{
		testDiffCase[int]{
			name: "int not equal", lhs: 1, rhs: 2,
			want: &diff.Result{Lhs: int(1), Rhs: int(2), Kind: diff.Different},
		},
		testDiffCase[int]{name: "int equal", lhs: 1, rhs: 1, want: nil},
		testDiffCase[myInt]{
			name: "myInt not equal", lhs: 1, rhs: 2,
			want: &diff.Result{Lhs: myInt(1), Rhs: myInt(2), Kind: diff.Different},
		},
		testDiffCase[myInt]{name: "myInt equal", lhs: 1, rhs: 1, want: nil},
		testDiffCase[getter]{
			name: "getter not equal", lhs: myInt(1), rhs: myInt(2),
			want: &diff.Result{
				Path: datapath.TypeAssert("myInt"),
				Lhs: myInt(1), Rhs: myInt(2), Kind: diff.Different,
			},
		},
		testDiffCase[getter]{name: "getter equal", lhs: myInt(1), rhs: myInt(1), want: nil},
		testDiffCase[getter]{name: "gitter nil", lhs: nil, rhs: nil, want: nil},
		testDiffCase[getter]{
			name: "getter one nil", lhs: nil, rhs: myInt(1),
			want: &diff.Result{Lhs: nil, Rhs: myInt(1), Kind: diff.Different},
		},
		testDiffCase[getter]{
			name: "getter different underlying types", lhs: myInt(1), rhs: myString("a"),
			want: &diff.Result{Lhs: myInt(1), Rhs: myString("a"), Kind: diff.Different},
		},
		testDiffCase[ptrGetter]{
			name: "ptrGetter not equal", lhs: ptr(myInt(1)), rhs: ptr(myInt(2)),
			want: &diff.Result{
				Path: datapath.TypeAssert("*myInt").PtrDeref(),
				Lhs: myInt(1), Rhs: myInt(2), Kind: diff.Different,
			},
		},
		testDiffCase[ptrGetter]{
			name: "ptrGetter equal", lhs: ptr(myInt(1)), rhs: ptr(myInt(1)), want: nil,
		},
		testDiffCase[ptrGetter]{name: "ptrGetter nil", lhs: nil, rhs: nil, want: nil},
		testDiffCase[ptrGetter]{
			name: "ptrGetter one nil", lhs: nil, rhs: ptr(myInt(1)),
			want: &diff.Result{Lhs: nil, Rhs: ptr(myInt(1)), Kind: diff.Different},
		},
		testDiffCase[*int]{
			name: "int ptr not equal", lhs: ptr(1), rhs: ptr(2),
			want: &diff.Result{
				Path: datapath.PtrDeref(),
				Lhs: int(1), Rhs: int(2), Kind: diff.Different,
			},
		},
		testDiffCase[*int]{name: "int ptr equal", lhs: ptr(1), rhs: ptr(1), want: nil},
		testDiffCase[*int]{name: "int ptr nil", lhs: nil, rhs: nil, want: nil},
		testDiffCase[*int]{
			name: "int ptr one nil", lhs: nil, rhs: ptr(1),
			want: &diff.Result{Lhs: nilPtr[int](), Rhs: ptr(int(1)), Kind: diff.Different},
		},
		testDiffCase[**int]{
			name: "int ptr ptr not equal", lhs: ptr(ptr(1)), rhs: ptr(ptr(2)),
			want: &diff.Result{
				Path: datapath.PtrDeref().PtrDeref(),
				Lhs: int(1), Rhs: int(2), Kind: diff.Different,
			},
		},
		testDiffCase[**int]{
			name: "int ptr ptr equal", lhs: ptr(ptr(1)), rhs: ptr(ptr(1)), want: nil,
		},
		testDiffCase[string]{
			name: "string not equal", lhs: "a", rhs: "b",
			want: &diff.Result{Lhs: "a", Rhs: "b", Kind: diff.Different},
		},
		testDiffCase[string]{name: "string equal", lhs: "a", rhs: "a", want: nil},
		testDiffCase[*string]{
			name: "string ptr not equal", lhs: ptr("a"), rhs: ptr("b"),
			want: &diff.Result{
				Path: datapath.PtrDeref(),
				Lhs: "a", Rhs: "b", Kind: diff.Different,
			},
		},
		testDiffCase[*string]{name: "string ptr equal", lhs: ptr("a"), rhs: ptr("a"), want: nil},
		testDiffCase[float64]{
			name: "float64 not equal", lhs: 1.0, rhs: 2.0,
			want: &diff.Result{Lhs: float64(1.0), Rhs: float64(2.0), Kind: diff.Different},
		},
		testDiffCase[float64]{name: "float64 equal", lhs: 1.0, rhs: 1.0, want: nil},
		testDiffCase[*float64]{
			name: "float64 ptr not equal", lhs: ptr(1.0), rhs: ptr(2.0),
			want: &diff.Result{
				Path: datapath.PtrDeref(),
				Lhs: float64(1.0), Rhs: float64(2.0), Kind: diff.Different,
			},
		},
		testDiffCase[*float64]{name: "float64 ptr equal", lhs: ptr(1.0), rhs: ptr(1.0), want: nil},
		testDiffCase[compStruct]{
			name: "compStruct not equal",
			lhs:  compStruct{Str: "a", Int: 1},
			rhs:  compStruct{Str: "b", Int: 2},
			want: &diff.Result{
				Path: datapath.Field("Str"),
				Lhs: "a", Rhs: "b", Kind: diff.Different,
			},
		},
		testDiffCase[compStruct]{
			name: "compStruct equal",
			lhs:  compStruct{Str: "a", Int: 1},
			rhs:  compStruct{Str: "a", Int: 1},
			want: nil,
		},
		testDiffCase[*compStruct]{
			name: "compStruct ptr not equal",
			lhs:  ptr(compStruct{Str: "a", Int: 1}),
			rhs:  ptr(compStruct{Str: "b", Int: 2}),
			want: &diff.Result{
				Path: datapath.PtrDeref().Field("Str"),
				Lhs: "a", Rhs: "b", Kind: diff.Different,
			},
		},
		testDiffCase[*compStruct]{
			name: "compStruct ptr equal",
			lhs:  ptr(compStruct{Str: "a", Int: 1}),
			rhs:  ptr(compStruct{Str: "a", Int: 1}),
			want: nil,
		},
		testDiffCase[nonCompStruct]{
			name: "nonCompStruct not equal",
			lhs:  nonCompStruct{Slice: []int{1, 2}},
			rhs:  nonCompStruct{Slice: []int{2, 1}},
			want: &diff.Result{
				Path: datapath.Field("Slice").Index(0),
				Lhs: int(1), Rhs: int(2), Kind: diff.Different,
			},
		},
		testDiffCase[nonCompStruct]{
			name: "nonCompStruct equal",
			lhs:  nonCompStruct{Slice: []int{1, 2}},
			rhs:  nonCompStruct{Slice: []int{1, 2}},
			want: nil,
		},
		testDiffCase[nonCompStruct]{
			name: "nonCompStruct unexported field ignored",
			lhs:  nonCompStruct{Slice: []int{1, 2}, pInt: 1},
			rhs:  nonCompStruct{Slice: []int{1, 2}, pInt: 2},
			want: nil,
		},
		testDiffCase[*nonCompStruct]{
			name: "nonCompStruct ptr not equal",
			lhs:  ptr(nonCompStruct{Slice: []int{1, 2}}),
			rhs:  ptr(nonCompStruct{Slice: []int{2, 1}}),
			want: &diff.Result{
				Path: datapath.PtrDeref().Field("Slice").Index(0),
				Lhs: int(1), Rhs: int(2), Kind: diff.Different,
			},
		},
		testDiffCase[*nonCompStruct]{
			name: "nonCompStruct ptr equal",
			lhs:  ptr(nonCompStruct{Slice: []int{1, 2}}),
			rhs:  ptr(nonCompStruct{Slice: []int{1, 2}}),
			want: nil,
		},
		testDiffCase[*nonCompStruct]{
			name: "nonCompStruct ptr one nil",
			lhs:  nil,
			rhs:  ptr(nonCompStruct{Slice: []int{1, 2}}),
			want: &diff.Result{
				Lhs:  nilPtr[nonCompStruct](),
				Rhs:  ptr(nonCompStruct{Slice: []int{1, 2}}),
				Kind: diff.Different,
			},
		},
		testDiffCase[*nonCompStruct]{name: "nonCompStruct ptr nil", lhs: nil, rhs: nil, want: nil},
		testDiffCase[[]int]{name: "slice nil", lhs: nil, rhs: nil, want: nil},
		testDiffCase[*[]int]{name: "slice ptr nil", lhs: nil, rhs: nil, want: nil},
		testDiffCase[[]int]{
			name: "slice not equal", lhs: []int{1, 2}, rhs: []int{2, 1},
			want: &diff.Result{
				Path: datapath.Index(0),
				Lhs: int(1), Rhs: int(2), Kind: diff.Different,
			},
		},
		testDiffCase[[]int]{name: "slice equal", lhs: []int{1, 2}, rhs: []int{1, 2}, want: nil},
		testDiffCase[[]int]{
			name: "slice lhs nil", lhs: nil, rhs: []int{1, 2},
			want: &diff.Result{Lhs: nil, Rhs: int(1), Kind: diff.Extra},
		},
		testDiffCase[[]int]{
			name: "slice rhs nil", lhs: []int{1, 2}, rhs: nil,
			want: &diff.Result{Lhs: int(1), Rhs: nil, Kind: diff.Missing},
		},
		testDiffCase[[]int]{
			name: "slice extra", lhs: []int{1, 2}, rhs: []int{1, 2, 3},
			want: &diff.Result{Rhs: int(3), Kind: diff.Extra},
		},
		testDiffCase[[]int]{
			name: "slice missing", lhs: []int{1, 2, 3}, rhs: []int{1, 2},
			want: &diff.Result{Lhs: int(3), Kind: diff.Missing},
		},
		testDiffCase[mySlice]{
			name: "mySlice not equal", lhs: mySlice{1, 2}, rhs: mySlice{2, 1},
			want: &diff.Result{
				Path: datapath.Index(0),
				Lhs: int(1), Rhs: int(2), Kind: diff.Different,
			},
		},
		testDiffCase[mySlice]{
			name: "mySlice equal", lhs: mySlice{1, 2}, rhs: mySlice{1, 2}, want: nil,
		},
		testDiffCase[mySlice]{
			name: "mySlice lhs nil", lhs: nil, rhs: mySlice{1, 2},
			want: &diff.Result{Lhs: nil, Rhs: int(1), Kind: diff.Extra},
		},
		testDiffCase[mySlice]{
			name: "mySlice rhs nil", lhs: mySlice{1, 2}, rhs: nil,
			want: &diff.Result{Lhs: int(1), Rhs: nil, Kind: diff.Missing},
		},
		testDiffCase[mySlice]{
			name: "mySlice extra", lhs: mySlice{1, 2}, rhs: mySlice{1, 2, 3},
			want: &diff.Result{Rhs: int(3), Kind: diff.Extra},
		},
		testDiffCase[mySlice]{
			name: "mySlice missing", lhs: mySlice{1, 2, 3}, rhs: mySlice{1, 2},
			want: &diff.Result{Lhs: int(3), Kind: diff.Missing},
		},
		testDiffCase[mySlice]{name: "mySlice nil", lhs: nil, rhs: nil, want: nil},
		testDiffCase[map[int]string]{name: "map nil", lhs: nil, rhs: nil, want: nil},
		testDiffCase[*map[int]string]{name: "map ptr nil", lhs: nil, rhs: nil, want: nil},
		testDiffCase[map[int]string]{
			name: "map not equal", lhs: map[int]string{1: "1"}, rhs: map[int]string{1: "2"},
			want: &diff.Result{
				Path: datapath.Key(1),
				Lhs: "1", Rhs: "2", Kind: diff.Different,
			},
		},
		testDiffCase[map[int]string]{
			name: "map equal",
			lhs: map[int]string{1: "1", 2: "2"},
			rhs: map[int]string{1: "1", 2: "2"},
			want: nil,
		},
		testDiffCase[map[int]string]{
			name: "map lhs nil", lhs: nil, rhs: map[int]string{1: "1"},
			want: &diff.Result{Lhs: nil, Rhs: 1, Kind: diff.Extra},
		},
		testDiffCase[map[int]string]{
			name: "map rhs nil", lhs: map[int]string{1: "1"}, rhs: nil,
			want: &diff.Result{Lhs: 1, Rhs: nil, Kind: diff.Missing},
		},
		testDiffCase[map[int]string]{
			name: "map extra",
			lhs: map[int]string{1: "1", 2: "2"},
			rhs: map[int]string{1: "1", 2: "2", 3: "3"},
			want: &diff.Result{Rhs: 3, Kind: diff.Extra},
		},
		testDiffCase[map[int]string]{
			name: "map missing",
			lhs: map[int]string{1: "1", 2: "2", 3: "3"},
			rhs: map[int]string{1: "1", 2: "2"},
			want: &diff.Result{Lhs: 3, Kind: diff.Missing},
		},
		testDiffCase[myMap]{
			name: "myMap not equal",lhs: myMap{1: "1"}, rhs: myMap{1: "2"},
			want: &diff.Result{
				Path: datapath.Key(1),
				Lhs: "1", Rhs: "2", Kind: diff.Different,
			},
		},
		testDiffCase[myMap]{
			name: "myMap equal", lhs: myMap{1: "1", 2: "2"}, rhs: myMap{1: "1", 2: "2"},
			want: nil,
		},
		testDiffCase[myMap]{
			name: "myMap lhs nil", lhs: nil, rhs: myMap{1: "1"},
			want: &diff.Result{Lhs: nil, Rhs: 1, Kind: diff.Extra},
		},
		testDiffCase[myMap]{
			name: "myMap rhs nil", lhs: myMap{1: "1"}, rhs: nil,
			want: &diff.Result{Lhs: 1, Rhs: nil, Kind: diff.Missing},
		},
		testDiffCase[myMap]{
			name: "myMap extra",
			lhs: myMap{1: "1", 2: "2"},
			rhs: myMap{1: "1", 2: "2", 3: "3"},
			want: &diff.Result{Rhs: 3, Kind: diff.Extra},
		},
		testDiffCase[myMap]{
			name: "myMap missing",
			lhs: myMap{1: "1", 2: "2", 3: "3"},
			rhs: myMap{1: "1", 2: "2"},
			want: &diff.Result{Lhs: 3, Kind: diff.Missing},
		},
		testDiffCase[myMap]{name: "myMap nil", lhs: nil, rhs: nil, want: nil},
		testDiffCase[ParentStruct]{
			name: "ChildStruct not equal",
			lhs:  ParentStruct{ChildStruct: ChildStruct{Str: "a"}, Int: 1},
			rhs:  ParentStruct{ChildStruct: ChildStruct{Str: "b"}, Int: 1},
			want: &diff.Result{
				Path: datapath.Field("ChildStruct").Field("Str"),
				Lhs: "a", Rhs: "b", Kind: diff.Different,
			},
		},
		testDiffCase[ParentStruct]{
			name: "ChildStruct equal",
			lhs:  ParentStruct{ChildStruct: ChildStruct{Str: "a"}, Int: 1},
			rhs:  ParentStruct{ChildStruct: ChildStruct{Str: "a"}, Int: 1},
			want: nil,
		},
	}
	for _, tt := range tests {
		tt.run(t)
	}
}
