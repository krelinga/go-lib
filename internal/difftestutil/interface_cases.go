package difftestutil

import (
	"github.com/krelinga/go-lib/datapath"
	"github.com/krelinga/go-lib/diff"
)

var interfaceCases = []TestCase{
	testCase[getter]{
		name: "getter not equal", lhs: myInt(1), rhs: myInt(2),
		want: []diff.Result{
			{
				Path: datapath.Method("Get"),
				Lhs:  int(1), Rhs: int(2), Kind: diff.Different,
			},
			// TODO: It doesn't seem desirable to have this method show up twice...
			{
				Path: datapath.TypeAssert("myInt").Method("Get"),
				Lhs:  int(1), Rhs: int(2), Kind: diff.Different,
			},
			{
				Path: datapath.TypeAssert("myInt"),
				Lhs:  myInt(1), Rhs: myInt(2), Kind: diff.Different,
			},
		},
	},
	testCase[getter]{name: "getter equal", lhs: myInt(1), rhs: myInt(1), want: nil},
	testCase[getter]{name: "gitter nil", lhs: nil, rhs: nil, want: nil},
	testCase[getter]{
		name: "getter one nil", lhs: nil, rhs: myInt(1),
		want: []diff.Result{{Lhs: nil, Rhs: myInt(1), Kind: diff.Different}},
	},
	testCase[getter]{
		name: "getter different underlying types", lhs: myInt(1), rhs: myString("a"),
		want: []diff.Result{{Lhs: myInt(1), Rhs: myString("a"), Kind: diff.Different}},
	},
	testCase[ptrGetter]{
		name: "ptrGetter not equal", lhs: ptr(myInt(1)), rhs: ptr(myInt(2)),
		want: []diff.Result{
			{
				Path: datapath.Method("GetPtr"),
				Lhs:  int(1), Rhs: int(2), Kind: diff.Different,
			},
			// TODO: it might be nice to disentangle the test types a bit to avoid this Get()
			// showing up in a test that is nominally about GetPtr()
			{
				Path: datapath.TypeAssert("*myInt").PtrDeref().Method("Get"),
				Lhs:  int(1), Rhs: int(2), Kind: diff.Different,
			},
			{
				Path: datapath.TypeAssert("*myInt").PtrDeref(),
				Lhs:  myInt(1), Rhs: myInt(2), Kind: diff.Different,
			},
		},
	},
	testCase[ptrGetter]{
		name: "ptrGetter equal", lhs: ptr(myInt(1)), rhs: ptr(myInt(1)), want: nil,
	},
	testCase[ptrGetter]{name: "ptrGetter nil", lhs: nil, rhs: nil, want: nil},
	testCase[ptrGetter]{
		name: "ptrGetter one nil", lhs: nil, rhs: ptr(myInt(1)),
		want: []diff.Result{{Lhs: nil, Rhs: ptr(myInt(1)), Kind: diff.Different}},
	},
}