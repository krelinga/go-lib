package geom

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCircleArc(t *testing.T) {
	tests := []struct {
		name                         string
		center                       *Point
		radius                       float64
		rayAngle                     RayAngle
		wantStartPoint, wantEndPoint *Point
		wantMinBB, wantMaxBB         *Point
	}{
		{
			name:           "0 to 90 degrees",
			center:         NewPoint(0, 0),
			radius:         1,
			rayAngle:       NewRayAngle(RayUp, RayRight, Clockwise),
			wantStartPoint: NewPoint(0, 1),
			wantEndPoint:   NewPoint(1, 0),
			wantMinBB:      NewPoint(0, 0),
			wantMaxBB:      NewPoint(1, 1),
		},
		{
			name:           "90 to 180 degrees",
			center:         NewPoint(0, 0),
			radius:         1,
			rayAngle:       NewRayAngle(RayRight, RayDown, Clockwise),
			wantStartPoint: NewPoint(1, 0),
			wantEndPoint:   NewPoint(0, -1),
			wantMinBB:      NewPoint(0, -1),
			wantMaxBB:      NewPoint(1, 0),
		},
		{
			name:           "top Arc",
			center:         NewPoint(0, 0),
			radius:         1,
			rayAngle:       NewRayAngle(RayUp.Offset(Degrees(45), CounterClockwise), RayUp.Offset(Degrees(45), Clockwise), Clockwise),
			wantStartPoint: NewPoint(-math.Sqrt(0.5), math.Sqrt(0.5)),
			wantEndPoint:   NewPoint(math.Sqrt(0.5), math.Sqrt(0.5)),
			wantMinBB:      NewPoint(-math.Sqrt(0.5), 0),
			wantMaxBB:      NewPoint(math.Sqrt(0.5), math.Sqrt(0.5)),
		},
		{
			name:           "bottom Arc",
			center:         NewPoint(0, 0),
			radius:         1,
			rayAngle:       NewRayAngle(RayUp.Offset(Degrees(45), CounterClockwise), RayUp.Offset(Degrees(45), Clockwise), CounterClockwise),
			wantStartPoint: NewPoint(-math.Sqrt(0.5), math.Sqrt(0.5)),
			wantEndPoint:   NewPoint(math.Sqrt(0.5), math.Sqrt(0.5)),
			wantMinBB:      NewPoint(-1, -1),
			wantMaxBB:      NewPoint(1, math.Sqrt(0.5)),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			arc := NewCircleArc(tt.center, tt.radius, tt.rayAngle)
			startPoint, endPoint := arc.Endpoints()
			assertXYEqual(t, tt.wantStartPoint, startPoint)
			assertXYEqual(t, tt.wantEndPoint, endPoint)
			assertXYEqual(t, tt.wantMinBB, arc.BoundingBox().BottomLeft())
			assertXYEqual(t, tt.wantMaxBB, arc.BoundingBox().TopRight())
		})
	}

	assert.Implements(t, (*Path)(nil), (*CircleArc)(nil))
}
