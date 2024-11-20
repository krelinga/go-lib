package geom

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTransform(t *testing.T) {
	t.Run("Translate", func(t *testing.T) {
		t.Run("Point", func(t *testing.T) {
			in := NewPoint(1, 2)
			out := Transform(in, Translate(3, 4))

			assert.Equal(t, 4.0, out.X())
			assert.Equal(t, 6.0, out.Y())

			assert.Equal(t, 1.0, in.X())
			assert.Equal(t, 2.0, in.Y())
		})
	})
}
