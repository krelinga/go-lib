package base_test

import (
	"testing"

	"github.com/krelinga/go-lib/base"
	"github.com/krelinga/go-lib/base/basetest"
	"github.com/stretchr/testify/assert"
)

type testDifferString struct {
	data string
}

func (tds testDifferString) Diff(other interface{}, reporter base.DiffReporter) {
	typedOther, ok := other.(testDifferString)
	if !ok {
		reporter.ReportTypeDiff(tds, other)
		return
	}

	if tds.data != typedOther.data {
		reporter.ChildField("data").ReportDiffValues(tds.data, typedOther.data)
	}
}

type testDifferInt struct {
	data int
}

func (tdi testDifferInt) Diff(other interface{}, reporter base.DiffReporter) {
	typedOther, ok := other.(testDifferInt)
	if !ok {
		reporter.ReportTypeDiff(tdi, other)
		return
	}

	if tdi.data != typedOther.data {
		reporter.ChildField("data").ReportDiffValues(tdi.data, typedOther.data)
	}
}

func TestSameAndDiff(t *testing.T) {
	tests := []struct {
		name            string
		a, b            base.Differ
		wantSame        bool
		wantTypeDiffs   []basetest.ReportedTypeDiff
		wantValueDiffs  []basetest.ReportedValueDiff
		wantStringDiffs []basetest.ReportedStringDiff
	}{
		{
			name:     "Same String",
			a:        testDifferString{data: "a"},
			b:        testDifferString{data: "a"},
			wantSame: true,
		},
		{
			name: "Different String",
			a:    testDifferString{data: "a"},
			b:    testDifferString{data: "b"},
			wantValueDiffs: []basetest.ReportedValueDiff{
				{Path: "data", A: "a", B: "b"},
			},
		},
		{
			name:     "Same Int",
			a:        testDifferInt{data: 1},
			b:        testDifferInt{data: 1},
			wantSame: true,
		},
		{
			name: "Different Int",
			a:    testDifferInt{data: 1},
			b:    testDifferInt{data: 2},
			wantValueDiffs: []basetest.ReportedValueDiff{
				{Path: "data", A: 1, B: 2},
			},
		},
		{
			name: "String vs Int",
			a:    testDifferString{data: "a"},
			b:    testDifferInt{data: 1},
			wantTypeDiffs: []basetest.ReportedTypeDiff{
				{A: testDifferString{data: "a"}, B: testDifferInt{data: 1}},
			},
		},
		{
			name: "Int vs String",
			a:    testDifferInt{data: 1},
			b:    testDifferString{data: "a"},
			wantTypeDiffs: []basetest.ReportedTypeDiff{
				{A: testDifferInt{data: 1}, B: testDifferString{data: "a"}},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := base.Same(tt.a, tt.b)
			assert.Equal(t, tt.wantSame, got)
			mdr := &basetest.MockDiffReporter{}
			base.Diff(tt.a, tt.b, mdr)
			assert.Equal(t, tt.wantTypeDiffs, mdr.ReportedTypeDiffs, "ReportedTypeDiffs mismatch")
			assert.Equal(t, tt.wantValueDiffs, mdr.ReportedValueDiffs, "ReportedValueDiffs mismatch")
			assert.Equal(t, tt.wantStringDiffs, mdr.ReportedStringDiffs, "ReportedStringDiffs mismatch")
		})
	}
}
