package geom

import "github.com/krelinga/go-lib/kiter"

func Rotate(p Point, a Angle) Point {
	return Point{
		p.X*Cos(a) - p.Y*Sin(a),
		p.X*Sin(a) + p.Y*Cos(a),
	}
}

func RotatePolygon(p Polygon, a Angle) Polygon {
	return kiter.SliceMap(p, func(p Point) Point {
		return Rotate(p, a)
	})
}
