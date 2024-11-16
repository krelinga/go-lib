package geom

type Point struct {
	X, Y float64
}

func Translate(p Point, dx, dy float64) Point {
	return Point{p.X + dx, p.Y + dy}
}
