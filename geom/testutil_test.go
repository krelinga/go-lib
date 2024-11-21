package geom

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func assertXYEqual(t *testing.T, expected, actual *Point) {
	t.Helper()
	const delta = 1e-9
	xFmt := "Max difference in X coordinate of (%v, %v) and (%v, %v) is %v, but difference was %v"
	xMsg := fmt.Sprintf(xFmt, expected.X(), expected.Y(), actual.X(), actual.Y(), delta, expected.X()-actual.X())
	assert.InDelta(t, expected.X(), actual.X(), delta, xMsg)
	yFmt := "Max difference in Y coordinate of (%v, %v) and (%v, %v) is %v, but difference was %v"
	yMsg := fmt.Sprintf(yFmt, expected.X(), expected.Y(), actual.X(), actual.Y(), delta, expected.Y()-actual.Y())
	assert.InDelta(t, expected.Y(), actual.Y(), delta, yMsg)
}
