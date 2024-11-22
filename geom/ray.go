package geom

import (
	"math"
)

type Ray struct {
	// Meaured in the clockwise direction from straight up.
	// Valid values are in the range [0, 360).
	degrees float64
}

var (
	RayUp    = Ray{0}
	RayRight = Ray{90}
	RayDown  = Ray{180}
	RayLeft  = Ray{270}
)

func (r Ray) Offset(a Angle, d Direction) Ray {
	a, d = normalize(a, d)
	deg := r.degrees
	if d == Clockwise {
		deg += a.Degrees()
	} else {
		deg -= a.Degrees()
	}
	deg = normalizeDegrees(deg)

	return Ray{deg}
}

// The returned degrees will always be in the range [0, 360).
func normalizeDegrees(deg float64) float64 {
	mult := math.Floor(deg / 360)
	return deg - mult*360
}

func (r Ray) Angle(other Ray, d Direction) Angle {
	var diffDeg float64
	if d == Clockwise {
		diffDeg = other.degrees - r.degrees
	} else {
		diffDeg = r.degrees - other.degrees
	}
	diffDeg = normalizeDegrees(diffDeg)
	return Degrees(diffDeg)
}

func (r Ray) Equals(other Ray) bool {
	return r.degrees == other.degrees
}

func (r Ray) DxDy(distance float64) (dx, dy float64) {
	return distance * Sin(Degrees(r.degrees)), distance * Cos(Degrees(r.degrees))
}

// the returned angle will always be in the range [0, 360) degrees.
func normalize(a Angle, d Direction) (Angle, Direction) {
	if a.Degrees() < 0 {
		d = !d
		a = Degrees(-a.Degrees())
	}
	mult := a.Degrees() / 360
	if mult >= 1 {
		a = Degrees(a.Degrees() - 360*mult)
	}
	return a, d
}

type RayAngle struct {
	from, to  Ray
	angle     Angle
	direction Direction
}

func NewRayAngle(from, to Ray, d Direction) RayAngle {
	a := from.Angle(to, d)
	a, d = normalize(a, d)
	return RayAngle{from, to, a, d}

}

func (r RayAngle) From() Ray {
	return r.from
}

func (r RayAngle) To() Ray {
	return r.to
}

func (r RayAngle) Angle() Angle {
	return r.angle
}

func (r RayAngle) Direction() Direction {
	return r.direction
}

func (r RayAngle) Reverse() RayAngle {
	a := Degrees(360 - r.angle.Degrees())
	d := !r.direction
	a, d = normalize(a, d)
	return RayAngle{r.to, r.from, a, d}
}

func (r RayAngle) Includes(ray Ray) bool {
	if r.direction == Clockwise {
		return r.from.degrees <= ray.degrees && ray.degrees <= r.to.degrees
	}
	return r.to.degrees <= ray.degrees && ray.degrees <= r.from.degrees
}

func (r RayAngle) Rotate(angle Angle, d Direction) RayAngle {
	return RayAngle{
		from:      r.from.Offset(angle, d),
		to:        r.to.Offset(angle, d),
		angle:     r.angle,
		direction: r.direction,
	}
}
