package geom

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHexagon(t *testing.T) {
	tests := []struct {
		r        float64
		expected Polygon
	}{
		{
			r: 1,
			expected: Polygon{
				Point{1, 0},
				Point{0.5, Sin(Degrees(60))},
				Point{-0.5, Sin(Degrees(60))},
				Point{-1, 0},
				Point{-0.5, -Sin(Degrees(60))},
				Point{0.5, -Sin(Degrees(60))},
			},
		},
		{
			r: 2,
			expected: Polygon{
				Point{2, 0},
				Point{1, 2 * Sin(Degrees(60))},
				Point{-1, 2 * Sin(Degrees(60))},
				Point{-2, 0},
				Point{-1, -2 * Sin(Degrees(60))},
				Point{1, -2 * Sin(Degrees(60))},
			},
		},
	}

	for _, test := range tests {
		result := Hexagon(test.r)
		for i, point := range result {
			assert.Equal(t, test.expected[i], point, "Hexagon(%v)[%d] = %v; want %v", test.r, i, point, test.expected[i])
		}
	}
}
