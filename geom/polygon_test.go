package geom

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPolygon_Lines(t *testing.T) {
	tests := []struct {
		name     string
		polygon  Polygon
		expected [][2]Point
	}{
		{
			name: "triangle",
			polygon: Polygon{
				{0, 0},
				{1, 0},
				{0, 1},
			},
			expected: [][2]Point{
				{{0, 0}, {1, 0}},
				{{1, 0}, {0, 1}},
				{{0, 1}, {0, 0}},
			},
		},
		{
			name: "square",
			polygon: Polygon{
				{0, 0},
				{1, 0},
				{1, 1},
				{0, 1},
			},
			expected: [][2]Point{
				{{0, 0}, {1, 0}},
				{{1, 0}, {1, 1}},
				{{1, 1}, {0, 1}},
				{{0, 1}, {0, 0}},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result [][2]Point
			for p1, p2 := range tt.polygon.Lines() {
				result = append(result, [2]Point{p1, p2})
			}

			require.Equal(t, len(tt.expected), len(result), "expected %d lines, got %d", len(tt.expected), len(result))

			for i, line := range result {
				assert.Equal(t, tt.expected[i], line, "expected line %v, got %v", tt.expected[i], line)
			}
		})
	}
}
