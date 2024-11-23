package geom

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHexagon(t *testing.T) {
	t.Run("Creation", func(t *testing.T) {
		var top LineTag
		var topLeft, topRight PointTag
		h := Hexagon(1, TagTopLine(&top), TagTopLeftPoint(&topLeft), TagTopRightPoint(&topRight))
		gotTop := top.Get(h)
		if !assert.NotNil(t, gotTop) {
			return
		}
		gotP1, gotP2 := gotTop.Endpoints()
		assertXYEqual(t, gotP1, topLeft.Get(h))
		assertXYEqual(t, gotP2, topRight.Get(h))
	})
}
