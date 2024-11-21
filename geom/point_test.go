package geom

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPoint(t *testing.T) {
	tests := []struct {
		x, y float64
		tags []PointTag
	}{
		{0, 0, nil},
		{1.5, -2.3, nil},
		{3.4, 5.6, make([]PointTag, 2)},
	}

	for _, tt := range tests {
		tagAddrs := make([]*PointTag, len(tt.tags))
		for i := range tt.tags {
			tagAddrs[i] = &tt.tags[i]
		}
		p := NewPoint(tt.x, tt.y, tagAddrs...)
		assert.Equal(t, tt.x, p.X())
		assert.Equal(t, tt.y, p.Y())
		for _, tag := range tt.tags {
			got := tag.Get(p)
			assert.Same(t, p, got)
		}

		var otherTag PointTag
		got := otherTag.Get(p)
		assert.Nil(t, got)
	}
}
