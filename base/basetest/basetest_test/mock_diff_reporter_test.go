package basetest_test

import (
	"testing"

	"github.com/krelinga/go-lib/base/basetest"
	"github.com/stretchr/testify/assert"
)

func TestMockDiffReporter(t *testing.T) {
	tests := []struct {
		name           string
		init           func(mdr *basetest.MockDiffReporter)
		wantTypeDiffs  []basetest.ReportedTypeDiff
		wantValueDiffs []basetest.ReportedValueDiff
	}{
		{
			name: "Empty",
		},
		{
			name: "TypeDiff",
			init: func(mdr *basetest.MockDiffReporter) {
				mdr.TypeDiff(1, "1")
			},
			wantTypeDiffs: []basetest.ReportedTypeDiff{
				{A: 1, B: "1"},
			},
		},
		{
			name: "ValueDiff",
			init: func(mdr *basetest.MockDiffReporter) {
				mdr.ReportDiffValues(1, 2)
			},
			wantValueDiffs: []basetest.ReportedValueDiff{
				{A: 1, B: 2},
			},
		},
		{
			name: "Child",
			init: func(mdr *basetest.MockDiffReporter) {
				child := mdr.ChildField("child")
				child.TypeDiff(1, "1")
				child.ReportDiffValues(1, 2)

				grandchild := child.ChildField("grandchild")
				grandchild.TypeDiff(10, "10")
				grandchild.ReportDiffValues(10, 20)
			},
			wantTypeDiffs: []basetest.ReportedTypeDiff{
				{Path: "child", A: 1, B: "1"},
				{Path: "child.grandchild", A: 10, B: "10"},
			},
			wantValueDiffs: []basetest.ReportedValueDiff{
				{Path: "child", A: 1, B: 2},
				{Path: "child.grandchild", A: 10, B: 20},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mdr := &basetest.MockDiffReporter{}
			if tt.init != nil {
				tt.init(mdr)
			}

			assert.Equal(t, tt.wantTypeDiffs, mdr.ReportedTypeDiffs, "ReportedTypeDiffs mismatch")
			assert.Equal(t, tt.wantValueDiffs, mdr.ReportedValueDiffs, "ReportedValueDiffs mismatch")
		})
	}
}
