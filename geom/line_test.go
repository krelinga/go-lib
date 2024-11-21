package geom

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLine(t *testing.T) {
	t.Run("Tags", func(t *testing.T) {
		pt := PointTag{}
		lt1 := LineTag{}
		lt2 := LineTag{}

		p1 := NewPoint(0, 0, &pt)
		p2 := NewPoint(1, 1)
		l := NewLine(p1, p2, &lt1, &lt2)

		assert.Same(t, p1, pt.Get(l))
		assert.Same(t, l, lt1.Get(l))
		assert.Same(t, l, lt2.Get(l))
	})

	t.Run("Basics", func(t *testing.T) {
		p1 := NewPoint(-1, -1)
		p2 := NewPoint(1, 1)
		l := NewLine(p1, p2)

		gotp1, gotp2 := l.Endpoints()
		assert.Same(t, p1, gotp1)
		assert.Same(t, p2, gotp2)

		bb := l.BoundingBox()
		assert.Same(t, p1, bb.Min)
		assert.Same(t, p2, bb.Max)
	})
}
