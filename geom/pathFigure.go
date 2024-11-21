package geom

import "iter"

// A figure is a collection of paths that form a closed shape.
type PathFigure struct {
	*tagBase
	mp *MultiPath
}

// NewPathFigure() creates a new figure from the given paths.
// Returns nil if the paths do not form a closed shape.
func NewPathFigure(paths ...Path) *PathFigure {
	mp := NewMultiPath(paths...)
	f := &PathFigure{
		tagBase: &mp.tagBase,
		mp:      mp,
	}

	// TODO: check that the figure is closed before returning.

	return f
}

func (f *PathFigure) BoundingBox() BoundingBox {
	return f.mp.BoundingBox()
}

func (f *PathFigure) Paths() iter.Seq[Path] {
	return f.mp.Paths()
}

func (f *PathFigure) clone() Element {
	return &PathFigure{
		tagBase: f.tagBase,
		mp:      clone(f.mp),
	}
}

func (f *PathFigure) translate(dx, dy float64) {
	f.mp.translate(dx, dy)
}

func (f *PathFigure) rotate(angle Angle, dir Direction) {
	f.mp.rotate(angle, dir)
}

func (f* PathFigure) figureIsAClosedType() {}