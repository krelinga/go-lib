package geom

import "math"

func Sin(a Angle) float64 {
	return math.Sin(a.Radians())
}

func Cos(a Angle) float64 {
	return math.Cos(a.Radians())
}

func Tan(a Angle) float64 {
	return math.Tan(a.Radians())
}

func Asin(x float64) Angle {
	return Radians(math.Asin(x))
}

func Acos(x float64) Angle {
	return Radians(math.Acos(x))
}

func Atan(x float64) Angle {
	return Radians(math.Atan(x))
}