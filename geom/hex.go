package geom

func Hexagon(r float64) Polygon {
	return Polygon{
		Point{r, 0},
		Point{r / 2, r * Sin(Degrees(60))},
		Point{-r / 2, r * Sin(Degrees(60))},
		Point{-r, 0},
		Point{-r / 2, -r * Sin(Degrees(60))},
		Point{r / 2, -r * Sin(Degrees(60))},
	}
}