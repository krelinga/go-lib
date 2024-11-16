package geom

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const Delta = 1e-9

func TestRotatePolygon(t *testing.T) {
	tests := []struct {
		name     string
		polygon  Polygon
		angle    Angle
		expected Polygon
	}{
		{
			name: "rotate square 90 degrees",
			polygon: Polygon{
				{0, 0},
				{1, 0},
				{1, 1},
				{0, 1},
			},
			angle: Degrees(90),
			expected: Polygon{
				{0, 0},
				{0, 1},
				{-1, 1},
				{-1, 0},
			},
		},
		{
			name: "rotate triangle 180 degrees",
			polygon: Polygon{
				{0, 0},
				{1, 0},
				{0.5, 1},
			},
			angle: Degrees(180),
			expected: Polygon{
				{0, 0},
				{-1, 0},
				{-0.5, -1},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := RotatePolygon(tt.polygon, tt.angle)
			for i, p := range result {
				assert.InDelta(t, tt.expected[i].X, p.X, Delta)
				assert.InDelta(t, tt.expected[i].Y, p.Y, Delta)
			}
		})
	}
}
func TestRotate(t *testing.T) {
	tests := []struct {
		name     string
		point    Point
		angle    Angle
		expected Point
	}{
		{
			name:     "rotate point 90 degrees",
			point:    Point{1, 0},
			angle:    Degrees(90),
			expected: Point{0, 1},
		},
		{
			name:     "rotate point 180 degrees",
			point:    Point{1, 0},
			angle:    Degrees(180),
			expected: Point{-1, 0},
		},
		{
			name:     "rotate point 270 degrees",
			point:    Point{1, 0},
			angle:    Degrees(270),
			expected: Point{0, -1},
		},
		{
			name:     "rotate point 360 degrees",
			point:    Point{1, 0},
			angle:    Degrees(360),
			expected: Point{1, 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Rotate(tt.point, tt.angle)
			assert.InDelta(t, tt.expected.X, result.X, Delta)
			assert.InDelta(t, tt.expected.Y, result.Y, Delta)
		})
	}
}
