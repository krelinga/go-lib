package geom

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRayOffsets(t *testing.T) {
	tests := []struct {
		name      string
		got, want Ray
	}{
		{
			name: "Zero Clockwise Offset From Up",
			got:  RayUp.Offset(Degrees(0), Clockwise),
			want: RayUp,
		},
		{
			name: "Zero CounterClockwise Offset From Up",
			got:  RayUp.Offset(Degrees(0), CounterClockwise),
			want: RayUp,
		},
		{
			name: "360 degrees clockwise offset from up",
			got:  RayUp.Offset(Degrees(360), Clockwise),
			want: RayUp,
		},
		{
			name: "360 degrees counterclockwise offset from up",
			got:  RayUp.Offset(Degrees(360), CounterClockwise),
			want: RayUp,
		},
		{
			name: "90 degrees clockwise offset from up and return",
			got:  RayUp.Offset(Degrees(90), Clockwise).Offset(Degrees(90), CounterClockwise),
			want: RayUp,
		},
		{
			name: "90 degrees counterclockwise offset from up and return",
			got:  RayUp.Offset(Degrees(90), CounterClockwise).Offset(Degrees(90), Clockwise),
			want: RayUp,
		},
		{
			name: "45 + 45 - 90",
			got:  RayUp.Offset(Degrees(45), Clockwise).Offset(Degrees(45), Clockwise).Offset(Degrees(90), CounterClockwise),
			want: RayUp,
		},
		{
			name: "45 + 45 - 90 (with direction negation)",
			got:  RayUp.Offset(Degrees(45), Clockwise).Offset(Degrees(45), Clockwise).Offset(Degrees(-90), Clockwise),
			want: RayUp,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.True(t, tt.got.Equals(tt.want))
			assert.True(t, tt.want.Equals(tt.got))
		})
	}
}

func TestRayAngle(t *testing.T) {
	tests := []struct {
		name     string
		from, to Ray
		dir      Direction

		wantAngle        Angle
		wantReverseAngle Angle
	}{
		{
			name:             "90 degrees clockwise from up",
			from:             RayUp,
			to:               RayRight,
			dir:              Clockwise,
			wantAngle:        Degrees(90),
			wantReverseAngle: Degrees(270),
		},
		{
			name:             "90 degrees counterclockwise from up",
			from:             RayUp,
			to:               RayRight,
			dir:              CounterClockwise,
			wantAngle:        Degrees(270),
			wantReverseAngle: Degrees(90),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ra := NewRayAngle(tt.from, tt.to, tt.dir)
			assert.Equal(t, tt.wantAngle, ra.Angle())
			assert.Equal(t, tt.dir, ra.Direction())
			reversed := ra.Reverse()
			assert.True(t, tt.wantReverseAngle.Equals(reversed.Angle()))
			assert.True(t, reversed.From().Equals(ra.To()))
			assert.True(t, reversed.To().Equals(ra.From()))
			assert.Equal(t, !tt.dir, reversed.Direction())
		})
	}
}
