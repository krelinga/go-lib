package difftestutil

import (
	"github.com/krelinga/go-lib/datapath"
	"github.com/krelinga/go-lib/diff"
)

var pointerCases = []TestCase{
	testCase[*int]{
		name: "int ptr not equal", lhs: ptr(1), rhs: ptr(2),
		want: []diff.Result{{
			Path: datapath.PtrDeref(),
			Lhs:  int(1), Rhs: int(2), Kind: diff.Different,
		}},
	},
	testCase[*int]{name: "int ptr equal", lhs: ptr(1), rhs: ptr(1), want: nil},
	testCase[*int]{name: "int ptr nil", lhs: nil, rhs: nil, want: nil},
	testCase[*int]{
		name: "int ptr one nil", lhs: nil, rhs: ptr(1),
		want: []diff.Result{{Lhs: nilPtr[int](), Rhs: ptr(int(1)), Kind: diff.Different}},
	},
	testCase[**int]{
		name: "int ptr ptr not equal", lhs: ptr(ptr(1)), rhs: ptr(ptr(2)),
		want: []diff.Result{{
			Path: datapath.PtrDeref().PtrDeref(),
			Lhs:  int(1), Rhs: int(2), Kind: diff.Different,
		}},
	},
	testCase[**int]{
		name: "int ptr ptr equal", lhs: ptr(ptr(1)), rhs: ptr(ptr(1)), want: nil,
	},
	testCase[*string]{
		name: "string ptr not equal", lhs: ptr("a"), rhs: ptr("b"),
		want: []diff.Result{{
			Path: datapath.PtrDeref(),
			Lhs:  "a", Rhs: "b", Kind: diff.Different,
		}},
	},
	testCase[*string]{name: "string ptr equal", lhs: ptr("a"), rhs: ptr("a"), want: nil},
	testCase[*float64]{
		name: "float64 ptr not equal", lhs: ptr(1.0), rhs: ptr(2.0),
		want: []diff.Result{{
			Path: datapath.PtrDeref(),
			Lhs:  float64(1.0), Rhs: float64(2.0), Kind: diff.Different,
		}},
	},
	testCase[*float64]{name: "float64 ptr equal", lhs: ptr(1.0), rhs: ptr(1.0), want: nil},
}
