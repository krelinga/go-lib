package geom

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const Delta = 1e-9

func TestTransform(t *testing.T) {
	t.Run("Translate", func(t *testing.T) {
		t.Run("Point", func(t *testing.T) {
			var inTag PointTag
			in := NewPoint(1, 2, &inTag)
			out := Transform(in, Translate(3, 4))

			assert.Equal(t, 4.0, out.X())
			assert.Equal(t, 6.0, out.Y())

			outFromTag := inTag.Get(out)
			assert.Equal(t, 4.0, outFromTag.X())
			assert.Equal(t, 6.0, outFromTag.Y())

			assert.Equal(t, 1.0, in.X())
			assert.Equal(t, 2.0, in.Y())
		})
	})

	t.Run("Rotate", func(t *testing.T) {
		t.Run("Line", func(t *testing.T) {
			tests := []struct {
				name           string
				p1, p2         *Point
				angle          Angle
				direction      Direction
				wantP1, wantP2 *Point
			}{
				{
					name:      "90 degrees clockwise",
					p1:        NewPoint(0, 1),
					p2:        NewPoint(1, 0),
					angle:     Degrees(90),
					direction: Clockwise,
					wantP1:    NewPoint(1, 0),
					wantP2:    NewPoint(0, -1),
				},
				{
					name:      "90 degrees counter-clockwise",
					p1:        NewPoint(0, 1),
					p2:        NewPoint(1, 0),
					angle:     Degrees(90),
					direction: CounterClockwise,
					wantP1:    NewPoint(-1, 0),
					wantP2:    NewPoint(0, 1),
				},
			}
			for _, tt := range tests {
				t.Run(tt.name, func(t *testing.T) {
					var lt LineTag
					in := NewLine(tt.p1, tt.p2, &lt)
					out := Transform(in, Rotate(tt.angle, tt.direction))

					p1, p2 := out.Endpoints()
					assert.InDelta(t, tt.wantP1.X(), p1.X(), Delta)
					assert.InDelta(t, tt.wantP1.Y(), p1.Y(), Delta)
					assert.InDelta(t, tt.wantP2.X(), p2.X(), Delta)
					assert.InDelta(t, tt.wantP2.Y(), p2.Y(), Delta)

					fromTag := lt.Get(out)
					p1, p2 = fromTag.Endpoints()
					assert.InDelta(t, tt.wantP1.X(), p1.X(), Delta)
					assert.InDelta(t, tt.wantP1.Y(), p1.Y(), Delta)
					assert.InDelta(t, tt.wantP2.X(), p2.X(), Delta)
					assert.InDelta(t, tt.wantP2.Y(), p2.Y(), Delta)

					p1, p2 = in.Endpoints()
					assert.InDelta(t, tt.p1.X(), p1.X(), Delta)
					assert.InDelta(t, tt.p1.Y(), p1.Y(), Delta)
					assert.InDelta(t, tt.p2.X(), p2.X(), Delta)
					assert.InDelta(t, tt.p2.Y(), p2.Y(), Delta)

					fromTag = lt.Get(in)
					p1, p2 = fromTag.Endpoints()
					assert.InDelta(t, tt.p1.X(), p1.X(), Delta)
					assert.InDelta(t, tt.p1.Y(), p1.Y(), Delta)
					assert.InDelta(t, tt.p2.X(), p2.X(), Delta)
					assert.InDelta(t, tt.p2.Y(), p2.Y(), Delta)
				})
			}
		})
	})
}
