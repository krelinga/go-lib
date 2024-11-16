package geom

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDegrees(t *testing.T) {
	tests := []struct {
		input    float64
		expected float64
	}{
		{0, 0},
		{90, math.Pi / 2},
		{180, math.Pi},
		{270, 3 * math.Pi / 2},
		{360, 2 * math.Pi},
	}

	for _, test := range tests {
		angle := Degrees(test.input)
		assert.Equal(t, test.input, angle.Degrees(), "Degrees(%f).Degrees() should be %f", test.input, test.input)
		assert.Equal(t, test.expected, angle.Radians(), "Degrees(%f).Radians() should be %f", test.input, test.expected)
	}
}
func TestRadians(t *testing.T) {
	tests := []struct {
		input    float64
		expected float64
	}{
		{0, 0},
		{math.Pi / 2, 90},
		{math.Pi, 180},
		{3 * math.Pi / 2, 270},
		{2 * math.Pi, 360},
	}

	for _, test := range tests {
		angle := Radians(test.input)
		assert.Equal(t, test.input, angle.Radians(), "Radians(%f).Radians() should be %f", test.input, test.input)
		assert.Equal(t, test.expected, angle.Degrees(), "Radians(%f).Degrees() should be %f", test.input, test.expected)
	}
}

