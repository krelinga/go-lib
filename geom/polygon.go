package geom

import "iter"

type Polygon []Point

func (s Polygon) Lines() iter.Seq2[Point, Point] {
	return func(yield func(Point, Point) bool) {
		for i := 0; i < len(s)-1; i++ {
			if !yield(s[i], s[i+1]) {
				return
			}
		}
		yield(s[len(s)-1], s[0])
	}
}

func (s Polygon) Width() float64 {
	minx, maxx := s[0].X, s[0].X
	for _, p := range s[1:] {
		if p.X < minx {
			minx = p.X
		}
		if p.X > maxx {
			maxx = p.X
		}
	}
	return maxx - minx
}

func (s Polygon) Height() float64 {
	miny, maxy := s[0].Y, s[0].Y
	for _, p := range s[1:] {
		if p.Y < miny {
			miny = p.Y
		}
		if p.Y > maxy {
			maxy = p.Y
		}
	}
	return maxy - miny
}
