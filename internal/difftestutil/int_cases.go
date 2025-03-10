package difftestutil

import (
	"github.com/krelinga/go-lib/datapath"
	"github.com/krelinga/go-lib/diff"
)

var intCases = []TestCase{
	testCase[int]{
		name: "int not equal", lhs: 1, rhs: 2,
		want: []diff.Result{{Lhs: int(1), Rhs: int(2), Kind: diff.Different}},
	},
	testCase[int]{name: "int equal", lhs: 1, rhs: 1, want: nil},
	testCase[myInt]{
		name: "myInt not equal", lhs: 1, rhs: 2,
		want: []diff.Result{
			{
				Path: datapath.Method("Get"),
				Lhs:  int(1), Rhs: int(2), Kind: diff.Different,
			},
			{
				Lhs: myInt(1), Rhs: myInt(2), Kind: diff.Different,
			},
		},
	},
	testCase[myInt]{name: "myInt equal", lhs: 1, rhs: 1, want: nil},
}