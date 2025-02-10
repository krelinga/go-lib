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
		wantMissing    []basetest.ReprtedMissing
		wantExtra      []basetest.ReportedExtra
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
				mdr.ValueDiff(1, 2)
			},
			wantValueDiffs: []basetest.ReportedValueDiff{
				{A: 1, B: 2},
			},
		},
		{
			name: "Missing",
			init: func(mdr *basetest.MockDiffReporter) {
				mdr.Missing(1)
			},
			wantMissing: []basetest.ReprtedMissing{
				{A: 1},
			},
		},
		{
			name: "Extra",
			init: func(mdr *basetest.MockDiffReporter) {
				mdr.Extra(1)
			},
			wantExtra: []basetest.ReportedExtra{
				{B: 1},
			},
		},
		{
			name: "ChildField",
			init: func(mdr *basetest.MockDiffReporter) {
				child := mdr.ChildField("child")
				child.TypeDiff(1, "1")
				child.ValueDiff(1, 2)

				grandchild := child.ChildField("grandchild")
				grandchild.TypeDiff(10, "10")
				grandchild.ValueDiff(10, 20)
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
		{
			name: "ChildKey",
			init: func(mdr *basetest.MockDiffReporter) {
				child := mdr.ChildKey(1)
				child.TypeDiff(2, "2")
				child.ValueDiff(3, 2)
				child.Missing(4)
				child.Extra(5)

				grandchild := child.ChildKey(10)
				grandchild.TypeDiff(20, "20")
				grandchild.ValueDiff(30, 20)
				grandchild.Missing(40)
				grandchild.Extra(50)
			},
			wantTypeDiffs: []basetest.ReportedTypeDiff{
				{Path: "[1]", A: 2, B: "2"},
				{Path: "[1][10]", A: 20, B: "20"},
			},
			wantValueDiffs: []basetest.ReportedValueDiff{
				{Path: "[1]", A: 3, B: 2},
				{Path: "[1][10]", A: 30, B: 20},
			},
			wantMissing: []basetest.ReprtedMissing{
				{Path: "[1]", A: 4},
				{Path: "[1][10]", A: 40},
			},
			wantExtra: []basetest.ReportedExtra{
				{Path: "[1]", B: 5},
				{Path: "[1][10]", B: 50},
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
