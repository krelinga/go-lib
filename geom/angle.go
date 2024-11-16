package geom

import "math"

type Angle interface {
	Degrees() float64
	Radians() float64
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