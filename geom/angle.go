package geom

import "math"

type Angle interface {
	Degrees() float64
	Radians() float64
	Equals(Angle) bool
}

func Degrees(d float64) Angle {
	return degrees(d)
}

type degrees float64

func (d degrees) Degrees() float64 {
	return float64(d)
}

func (d degrees) Radians() float64 {
	return float64(d) * math.Pi / 180
}

func (d degrees) Equals(other Angle) bool {
	return d.Degrees() == other.Degrees()
}

func Radians(r float64) Angle {
	return radians(r)
}

type radians float64

func (r radians) Degrees() float64 {
	return float64(r) * 180 / math.Pi
}

func (r radians) Radians() float64 {
	return float64(r)
}

func (r radians) Equals(other Angle) bool {
	return r.Radians() == other.Radians()
}

type AngleFromUp struct {
	angle     Angle
	direction Direction
}

func NewAngleFromUp(angle Angle, direction Direction) AngleFromUp {
	return AngleFromUp{angle, direction}
}

func (a AngleFromUp) Angle() Angle {
	return a.angle
}

func (a AngleFromUp) Direction() Direction {
	return a.direction
}
