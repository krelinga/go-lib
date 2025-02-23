package diffops_test

import (
	"testing"

	"github.com/krelinga/go-lib/ops/diffops"
	"github.com/krelinga/go-lib/ops/diffops/diffopsmock"
	"github.com/stretchr/testify/assert"
)

func TestCast(t *testing.T) {
	tests := []struct {
		name string
		lhs int
		rhs any
		wantDiffs []diffopsmock.Diff
	}{
		{
			name: "compatible types, different values",
			lhs: 1,
			rhs: 2,
			wantDiffs: []diffopsmock.Diff{
				{ValueDiff: &diffopsmock.Pair{Lhs: 1, Rhs: 2}},
			},
		},
		{
			name: "compatible types, same values",
			lhs: 1,
			rhs: 1,
			wantDiffs: []diffopsmock.Diff{},
		},
		{
			name: "incompatible types",
			lhs: 1,
			rhs: "1",
			wantDiffs: []diffopsmock.Diff{
				{TypeDiff: &diffopsmock.Pair{Lhs: 1, Rhs: "1"}},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sink := &diffopsmock.Sink{}
			diffops.CastRhs(tt.lhs, tt.rhs, sink, func(rhs int) {
				if tt.lhs != rhs {
					sink.ValueDiff(tt.lhs, rhs)
				}
			})
			if !assert.Equal(t, len(tt.wantDiffs), len(sink.Diffs)) {
				return
			}
			for i := range tt.wantDiffs {
				assert.Equal(t, tt.wantDiffs[i], sink.Diffs[i])
			}
		})
	}
}