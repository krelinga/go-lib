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

		assert.Equal(t, center.X(), circle.Center().X())
		assert.Equal(t, center.Y(), circle.Center().Y())

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
		assert.Equal(t, 1.0, tCenter.X())
		assert.Equal(t, 2.0, tCenter.Y())
		assert.Equal(t, radius, tCircle.Radius())
		assert.Equal(t, ct.Get(tCircle), tCircle)
	})
}
