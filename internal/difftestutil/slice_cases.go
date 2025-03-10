package difftestutil

import (
	"github.com/krelinga/go-lib/datapath"
	"github.com/krelinga/go-lib/diff"
)

var sliceCases = []TestCase{
	testCase[[]int]{name: "slice nil", lhs: nil, rhs: nil, want: nil},
	testCase[*[]int]{name: "slice ptr nil", lhs: nil, rhs: nil, want: nil},
	testCase[[]int]{
		name: "slice not equal", lhs: []int{1, 2}, rhs: []int{2, 1},
		want: []diff.Result{
			{
				Path: datapath.Index(0),
				Lhs:  int(1), Rhs: int(2), Kind: diff.Different,
			},
			{
				Path: datapath.Index(1),
				Lhs:  int(2), Rhs: int(1), Kind: diff.Different,
			},
		},
	},
	testCase[[]int]{name: "slice equal", lhs: []int{1, 2}, rhs: []int{1, 2}, want: nil},
	testCase[[]int]{
		name: "slice lhs nil", lhs: nil, rhs: []int{1, 2},
		want: []diff.Result{
			{
				Path: datapath.Index(0),
				Lhs:  nil, Rhs: int(1), Kind: diff.Extra,
			},
			{
				Path: datapath.Index(1),
				Lhs:  nil, Rhs: int(2), Kind: diff.Extra,
			},
		},
	},
	testCase[[]int]{
		name: "slice rhs nil", lhs: []int{1, 2}, rhs: nil,
		want: []diff.Result{
			{
				Path: datapath.Index(0),
				Lhs:  int(1), Rhs: nil, Kind: diff.Missing,
			},
			{
				Path: datapath.Index(1),
				Lhs:  int(2), Rhs: nil, Kind: diff.Missing,
			},
		},
	},
	testCase[[]int]{
		name: "slice extra", lhs: []int{1, 2}, rhs: []int{1, 2, 3},
		want: []diff.Result{
			{
				Path: datapath.Index(2),
				Rhs:  int(3), Kind: diff.Extra,
			},
		},
	},
	testCase[[]int]{
		name: "slice missing", lhs: []int{1, 2, 3}, rhs: []int{1, 2},
		want: []diff.Result{
			{
				Path: datapath.Index(2),
				Lhs:  int(3), Kind: diff.Missing,
			},
		},
	},
	testCase[mySlice]{
		name: "mySlice not equal", lhs: mySlice{1, 2}, rhs: mySlice{2, 1},
		want: []diff.Result{
			{
				Path: datapath.Index(0),
				Lhs:  int(1), Rhs: int(2), Kind: diff.Different,
			},
			{
				Path: datapath.Index(1),
				Lhs:  int(2), Rhs: int(1), Kind: diff.Different,
			},
		},
	},
	testCase[mySlice]{
		name: "mySlice equal", lhs: mySlice{1, 2}, rhs: mySlice{1, 2}, want: nil,
	},
	testCase[mySlice]{
		name: "mySlice lhs nil", lhs: nil, rhs: mySlice{1, 2},
		want: []diff.Result{
			{
				Path: datapath.Index(0),
				Lhs:  nil, Rhs: int(1), Kind: diff.Extra,
			},
			{
				Path: datapath.Index(1),
				Lhs:  nil, Rhs: int(2), Kind: diff.Extra,
			},
		},
	},
	testCase[mySlice]{
		name: "mySlice rhs nil", lhs: mySlice{1, 2}, rhs: nil,
		want: []diff.Result{
			{
				Path: datapath.Index(0),
				Lhs:  int(1), Rhs: nil, Kind: diff.Missing,
			},
			{
				Path: datapath.Index(1),
				Lhs:  int(2), Rhs: nil, Kind: diff.Missing,
			},
		},
	},
	testCase[mySlice]{
		name: "mySlice extra", lhs: mySlice{1, 2}, rhs: mySlice{1, 2, 3},
		want: []diff.Result{
			{
				Path: datapath.Index(2),
				Rhs:  int(3), Kind: diff.Extra,
			},
		},
	},
	testCase[mySlice]{
		name: "mySlice missing", lhs: mySlice{1, 2, 3}, rhs: mySlice{1, 2},
		want: []diff.Result{
			{
				Path: datapath.Index(2),
				Lhs:  int(3), Kind: diff.Missing,
			},
		},
	},
	testCase[mySlice]{name: "mySlice nil", lhs: nil, rhs: nil, want: nil},
}