package difftestutil

import (
	"github.com/krelinga/go-lib/datapath"
	"github.com/krelinga/go-lib/diff"
)

var mapCases = []TestCase{
	testCase[map[int]string]{name: "map nil", lhs: nil, rhs: nil, want: nil},
	testCase[*map[int]string]{name: "map ptr nil", lhs: nil, rhs: nil, want: nil},
	testCase[map[int]string]{
		name: "map not equal", lhs: map[int]string{1: "1"}, rhs: map[int]string{1: "2"},
		want: []diff.Result{{
			Path: datapath.Key(1),
			Lhs:  "1", Rhs: "2", Kind: diff.Different,
		}},
	},
	testCase[map[int]string]{
		name: "map equal",
		lhs:  map[int]string{1: "1", 2: "2"},
		rhs:  map[int]string{1: "1", 2: "2"},
		want: nil,
	},
	testCase[map[int]string]{
		name: "map lhs nil", lhs: nil, rhs: map[int]string{1: "1"},
		want: []diff.Result{{
			Path: datapath.Key(1),
			Lhs:  nil, Rhs: "1", Kind: diff.Extra,
		}},
	},
	testCase[map[int]string]{
		name: "map rhs nil", lhs: map[int]string{1: "1"}, rhs: nil,
		want: []diff.Result{{
			Path: datapath.Key(1),
			Lhs:  "1", Rhs: nil, Kind: diff.Missing,
		}},
	},
	testCase[map[int]string]{
		name: "map extra",
		lhs:  map[int]string{1: "1", 2: "2"},
		rhs:  map[int]string{1: "1", 2: "2", 3: "3"},
		want: []diff.Result{{
			Path: datapath.Key(3),
			Rhs:  "3", Kind: diff.Extra,
		}},
	},
	testCase[map[int]string]{
		name: "map missing",
		lhs:  map[int]string{1: "1", 2: "2", 3: "3"},
		rhs:  map[int]string{1: "1", 2: "2"},
		want: []diff.Result{{
			Path: datapath.Key(3),
			Lhs:  "3", Kind: diff.Missing,
		}},
	},
	testCase[myMap]{
		name: "myMap not equal", lhs: myMap{1: "1"}, rhs: myMap{1: "2"},
		want: []diff.Result{{
			Path: datapath.Key(1),
			Lhs:  "1", Rhs: "2", Kind: diff.Different,
		}},
	},
	testCase[myMap]{
		name: "myMap equal", lhs: myMap{1: "1", 2: "2"}, rhs: myMap{1: "1", 2: "2"},
		want: nil,
	},
	testCase[myMap]{
		name: "myMap lhs nil", lhs: nil, rhs: myMap{1: "1"},
		want: []diff.Result{{
			Path: datapath.Key(1),
			Lhs:  nil, Rhs: "1", Kind: diff.Extra,
		}},
	},
	testCase[myMap]{
		name: "myMap rhs nil", lhs: myMap{1: "1"}, rhs: nil,
		want: []diff.Result{{
			Path: datapath.Key(1),
			Lhs:  "1", Rhs: nil, Kind: diff.Missing,
		}},
	},
	testCase[myMap]{
		name: "myMap extra",
		lhs:  myMap{1: "1", 2: "2"},
		rhs:  myMap{1: "1", 2: "2", 3: "3"},
		want: []diff.Result{{
			Path: datapath.Key(3),
			Rhs:  "3", Kind: diff.Extra,
		}},
	},
	testCase[myMap]{
		name: "myMap missing",
		lhs:  myMap{1: "1", 2: "2", 3: "3"},
		rhs:  myMap{1: "1", 2: "2"},
		want: []diff.Result{{
			Path: datapath.Key(3),
			Lhs:  "3", Kind: diff.Missing,
		}},
	},
	testCase[myMap]{name: "myMap nil", lhs: nil, rhs: nil, want: nil},
}