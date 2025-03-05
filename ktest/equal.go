package ktest

import (
	"fmt"

	"github.com/krelinga/go-lib/diff"
)

const fmtDifferent = `
- %s: %v
+ %s: %v
`

const fmtMissing = `
- %s: %v
`

const fmtExtra = `
+ %s: %v
`

// Returns true if lhs and rhs are equal, false otherwise.
func AssertEqual[T any](t TestingT, lhs, rhs T) bool {
	t.Helper()
	diffResult := diff.Diff(lhs, rhs)
	if diffResult == nil {
		return true
	}
	switch diffResult.Kind {
	case diff.Different:
		t.Errorf(fmtDifferent, diffResult.Path.Basename("lhs"), diffResult.Lhs, diffResult.Path.Basename("rhs"), diffResult.Rhs)
	case diff.Missing:
		t.Errorf(fmtMissing, diffResult.Path.Basename("lhs"), diffResult.Lhs)
	case diff.Extra:
		t.Errorf(fmtExtra, diffResult.Path.Basename("rhs"), diffResult.Rhs)
	default:
		panic(fmt.Sprintf("unexpected diff kind: %v", diffResult.Kind))
	}
	return false
}