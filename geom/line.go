package geom

func Midpoint(p1, p2 Point) Point {
	return Point{
		(p1.X + p2.X) / 2,
		(p1.Y + p2.Y) / 2,
	}
}
