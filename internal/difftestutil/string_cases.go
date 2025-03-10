package difftestutil

import "github.com/krelinga/go-lib/diff"

var stringCases = []TestCase{
	testCase[string]{
		name: "string not equal", lhs: "a", rhs: "b",
		want: []diff.Result{{Lhs: "a", Rhs: "b", Kind: diff.Different}},
	},
	testCase[string]{name: "string equal", lhs: "a", rhs: "a", want: nil},
}