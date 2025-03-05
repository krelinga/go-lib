package diff_test

import (
	"testing"

	"github.com/krelinga/go-lib/internal/difftestutil"
)

func TestDiff(t *testing.T) {
	for _, tt := range difftestutil.TestCases {
		tt.RunDiffTest(t)
	}
}
