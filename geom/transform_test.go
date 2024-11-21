package geom

import (
	"testing"
)

func TestTransform(t *testing.T) {
	t.Run("Translate", func(t *testing.T) {
		t.Run("Point", func(t *testing.T) {
			var inTag PointTag
			in := NewPoint(1, 2, &inTag)
			out := Transform(in, Translate(3, 4))
			assertXYEqual(t, NewPoint(4, 6), out)

			outFromTag := inTag.Get(out)
			assertXYEqual(t, NewPoint(4, 6), outFromTag)
			assertXYEqual(t, NewPoint(1, 2), in)
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
					assertXYEqual(t, tt.wantP1, p1)
					assertXYEqual(t, tt.wantP2, p2)

					fromTag := lt.Get(out)
					p1, p2 = fromTag.Endpoints()
					assertXYEqual(t, tt.wantP1, p1)
					assertXYEqual(t, tt.wantP2, p2)

					p1, p2 = in.Endpoints()
					assertXYEqual(t, tt.p1, p1)
					assertXYEqual(t, tt.p2, p2)

					fromTag = lt.Get(in)
					p1, p2 = fromTag.Endpoints()
					assertXYEqual(t, tt.p1, p1)
					assertXYEqual(t, tt.p2, p2)
				})
			}
		})
	})
}
