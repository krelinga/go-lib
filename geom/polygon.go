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