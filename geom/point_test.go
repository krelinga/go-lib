package geom

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTranslate(t *testing.T) {
	tests := []struct {
		name     string
		point    Point
		dx, dy   float64
		expected Point
	}{
		{
			name:     "translate by positive values",
			point:    Point{X: 1, Y: 2},
			dx:       3,
			dy:       4,
			expected: Point{X: 4, Y: 6},
		},
		{
			name:     "translate by negative values",
			point:    Point{X: 1, Y: 2},
			dx:       -1,
			dy:       -2,
			expected: Point{X: 0, Y: 0},
		},
		{
			name:     "translate by zero",
			point:    Point{X: 1, Y: 2},
			dx:       0,
			dy:       0,
			expected: Point{X: 1, Y: 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Translate(tt.point, tt.dx, tt.dy)
			assert.Equal(t, tt.expected, result, "Translate(%v, %v, %v)", tt.point, tt.dx, tt.dy)
		})
	}
}
