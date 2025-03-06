package ktest

import (
	"fmt"
	"strings"

	"github.com/krelinga/go-lib/diff"
)

const fmtDifferent = `
- %s: %s
+ %s: %s`

const fmtMissing = `
- %s: %s`

const fmtExtra = `
+ %s: %s`

func format(v any) string {
	raw := fmt.Sprintf("%v", v)
	if strings.HasPrefix(raw, "0x") {
		return "<non-nil>"
	}
	return raw
}

// Returns true if lhs and rhs are equal, false otherwise.
func AssertEqual[T any](t TestingT, lhs, rhs T) bool {
	t.Helper()
	results := diff.Diff(lhs, rhs)
	if results == nil {
		return true
	}
	for i, r := range results {
		switch r.Kind {
		case diff.Different:
			t.Errorf(fmtDifferent, r.Path.Basename("lhs"), format(r.Lhs), r.Path.Basename("rhs"), format(r.Rhs))
		case diff.Missing:
			t.Errorf(fmtMissing, r.Path.Basename("lhs"), format(r.Lhs))
		case diff.Extra:
			t.Errorf(fmtExtra, r.Path.Basename("rhs"), format(r.Rhs))
		default:
			panic(fmt.Sprintf("unexpected diff kind: %v", r.Kind))
		}
		if i == len(results)-1 {
			t.Errorf("\n")
		}
	}

	return false
}
