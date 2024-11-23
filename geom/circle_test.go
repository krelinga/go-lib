package geom

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCircle(t *testing.T) {
	t.Run("Basics", func(t *testing.T) {
		const radius = 5.0
		circle := NewCircle(radius)

		assertXYEqual(t, NewPoint(0, 0), circle.Center())
		assert.Equal(t, radius, circle.Radius())
	})

	t.Run("Tags", func(t *testing.T) {
		var pt PointTag
		const radius = 5.0
		circle := NewCircle(radius, TagCenterPoint(&pt))

		tCircle := Transform(circle, Translate(1, 2))
		tCenter := pt.Get(tCircle)
		assertXYEqual(t, NewPoint(1, 2), tCenter)
		assert.Equal(t, radius, tCircle.Radius())
	})
}
