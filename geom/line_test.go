package geom

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMidpoint(t *testing.T) {
	tests := []struct {
		p1, p2, expected Point
	}{
		{Point{0, 0}, Point{2, 2}, Point{1, 1}},
		{Point{1, 1}, Point{3, 3}, Point{2, 2}},
		{Point{-1, -1}, Point{1, 1}, Point{0, 0}},
		{Point{0, 0}, Point{0, 0}, Point{0, 0}},
	}

	for _, test := range tests {
		result := Midpoint(test.p1, test.p2)
		assert.Equal(t, test.expected, result, "Midpoint(%v, %v)", test.p1, test.p2)
	}
}
