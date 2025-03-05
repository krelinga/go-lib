package ktest_test

import (
	"testing"

	"github.com/krelinga/go-lib/internal/difftestutil"
)

func TestAssertEqual(t *testing.T) {
	for _, tt := range difftestutil.TestCases {
		tt.RunAssertEqualTest(t)
	}
}
