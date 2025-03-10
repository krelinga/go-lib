package difftestutil

import (
	"github.com/krelinga/go-lib/datapath"
	"github.com/krelinga/go-lib/diff"
)

var structCases = []TestCase{
	testCase[compStruct]{
		name: "compStruct not equal",
		lhs:  compStruct{Str: "a", Int: 1},
		rhs:  compStruct{Str: "b", Int: 2},
		want: []diff.Result{
			{
				Path: datapath.Field("Str"),
				Lhs:  "a", Rhs: "b", Kind: diff.Different,
			},
			{
				Path: datapath.Field("Int"),
				Lhs:  int(1), Rhs: int(2), Kind: diff.Different,
			},
		},
	},
	testCase[compStruct]{
		name: "compStruct equal",
		lhs:  compStruct{Str: "a", Int: 1},
		rhs:  compStruct{Str: "a", Int: 1},
		want: nil,
	},
	testCase[*compStruct]{
		name: "compStruct ptr not equal",
		lhs:  ptr(compStruct{Str: "a", Int: 1}),
		rhs:  ptr(compStruct{Str: "b", Int: 2}),
		want: []diff.Result{{
			Path: datapath.PtrDeref().Field("Str"),
			Lhs:  "a", Rhs: "b", Kind: diff.Different,
		},
			{
				Path: datapath.PtrDeref().Field("Int"),
				Lhs:  int(1), Rhs: int(2), Kind: diff.Different,
			}},
	},
	testCase[*compStruct]{
		name: "compStruct ptr equal",
		lhs:  ptr(compStruct{Str: "a", Int: 1}),
		rhs:  ptr(compStruct{Str: "a", Int: 1}),
		want: nil,
	},
	testCase[nonCompStruct]{
		name: "nonCompStruct not equal",
		lhs:  nonCompStruct{Slice: []int{1, 2}},
		rhs:  nonCompStruct{Slice: []int{2, 1}},
		want: []diff.Result{
			{
				Path: datapath.Field("Slice").Index(0),
				Lhs:  int(1), Rhs: int(2), Kind: diff.Different,
			},
			{
				Path: datapath.Field("Slice").Index(1),
				Lhs:  int(2), Rhs: int(1), Kind: diff.Different,
			},
		},
	},
	testCase[nonCompStruct]{
		name: "nonCompStruct equal",
		lhs:  nonCompStruct{Slice: []int{1, 2}},
		rhs:  nonCompStruct{Slice: []int{1, 2}},
		want: nil,
	},
	testCase[nonCompStruct]{
		name: "nonCompStruct unexported field ignored",
		lhs:  nonCompStruct{Slice: []int{1, 2}, pInt: 1},
		rhs:  nonCompStruct{Slice: []int{1, 2}, pInt: 2},
		want: nil,
	},
	testCase[*nonCompStruct]{
		name: "nonCompStruct ptr not equal",
		lhs:  ptr(nonCompStruct{Slice: []int{1, 2}}),
		rhs:  ptr(nonCompStruct{Slice: []int{2, 1}}),
		want: []diff.Result{
			{
				Path: datapath.PtrDeref().Field("Slice").Index(0),
				Lhs:  int(1), Rhs: int(2), Kind: diff.Different,
			},
			{
				Path: datapath.PtrDeref().Field("Slice").Index(1),
				Lhs:  int(2), Rhs: int(1), Kind: diff.Different,
			},
		},
	},
	testCase[*nonCompStruct]{
		name: "nonCompStruct ptr equal",
		lhs:  ptr(nonCompStruct{Slice: []int{1, 2}}),
		rhs:  ptr(nonCompStruct{Slice: []int{1, 2}}),
		want: nil,
	},
	testCase[*nonCompStruct]{
		name: "nonCompStruct ptr one nil",
		lhs:  nil,
		rhs:  ptr(nonCompStruct{Slice: []int{1, 2}}),
		want: []diff.Result{{
			Lhs:  nilPtr[nonCompStruct](),
			Rhs:  ptr(nonCompStruct{Slice: []int{1, 2}}),
			Kind: diff.Different,
		}},
	},
	testCase[*nonCompStruct]{name: "nonCompStruct ptr nil", lhs: nil, rhs: nil, want: nil},
	testCase[ParentStruct]{
		name: "ChildStruct not equal",
		lhs:  ParentStruct{ChildStruct: ChildStruct{Str: "a"}, Int: 1},
		rhs:  ParentStruct{ChildStruct: ChildStruct{Str: "b"}, Int: 1},
		want: []diff.Result{{
			Path: datapath.Field("ChildStruct").Field("Str"),
			Lhs:  "a", Rhs: "b", Kind: diff.Different,
		}},
	},
	testCase[ParentStruct]{
		name: "ChildStruct equal",
		lhs:  ParentStruct{ChildStruct: ChildStruct{Str: "a"}, Int: 1},
		rhs:  ParentStruct{ChildStruct: ChildStruct{Str: "a"}, Int: 1},
		want: nil,
	},
}