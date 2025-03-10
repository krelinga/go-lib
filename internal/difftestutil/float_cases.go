package difftestutil

import "github.com/krelinga/go-lib/diff"

var floatCases = []TestCase{
	testCase[float64]{
		name: "float64 not equal", lhs: 1.0, rhs: 2.0,
		want: []diff.Result{{Lhs: float64(1.0), Rhs: float64(2.0), Kind: diff.Different}},
	},
	testCase[float64]{name: "float64 equal", lhs: 1.0, rhs: 1.0, want: nil},
}