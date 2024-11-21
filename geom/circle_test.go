package geom

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCircle(t *testing.T) {
	t.Run("Basics", func(t *testing.T) {
		center := NewPoint(0, 0)
		const radius = 5.0
		circle := NewCircle(center, radius)

		assertXYEqual(t, center, circle.Center())
		assert.Equal(t, radius, circle.Radius())
	})

	t.Run("Tags", func(t *testing.T) {
		var pt PointTag
		center := NewPoint(0, 0, &pt)
		const radius = 5.0
		var ct CircleTag
		circle := NewCircle(center, radius, &ct)

		tCircle := Transform(circle, Translate(1, 2))
		tCenter := pt.Get(tCircle)
		assertXYEqual(t, NewPoint(1, 2), tCenter)
		assert.Equal(t, radius, tCircle.Radius())
		assert.Same(t, ct.Get(tCircle), tCircle)
	})
}
