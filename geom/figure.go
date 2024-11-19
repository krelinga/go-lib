package geom

import "iter"

// A figure is a collection of paths that form a closed shape.
type Figure struct {
	*tagBase
	mp *MultiPath
}

// NewFigure() creates a new figure from the given paths.
// Returns nil if the paths do not form a closed shape.
func NewFigure(paths ...Path) *Figure {
	mp := NewMultiPath(paths...)
	f := &Figure{
		tagBase: &mp.tagBase,
		mp:      mp,
	}

	// TODO: check that the figure is closed before returning.

	return f
}

func (f *Figure) BoundingBox() BoundingBox {
	return f.mp.BoundingBox()
}

func (f *Figure) Paths() iter.Seq[Path] {
	return f.mp.Paths()
}
