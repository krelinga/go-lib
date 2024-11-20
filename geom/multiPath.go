package geom

import (
	"iter"
	"slices"
)

type MultiPath struct {
	tagBase
	paths []Path
}

func NewMultiPath(paths ...Path) *MultiPath {
	// TODO: translate paths so that the starting point of path n-1
	// is the same as the ending point of path n.
	mp := &MultiPath{paths: paths}
	for _, path := range paths {
		mp.addAllTags(path.getTagIndex())
	}
	return mp
}

func (mp *MultiPath) BoundingBox() BoundingBox {
	// TODO: implement this
	return BoundingBox{}
}

func (mp *MultiPath) Endpoints() (*Point, *Point) {
	start, _ := mp.paths[0].Endpoints()
	_, end := mp.paths[len(mp.paths)-1].Endpoints()
	return start, end
}

func (mu MultiPath) Paths() iter.Seq[Path] {
	return slices.Values(mu.paths)
}

func (mp MultiPath) Extend(paths ...Path) *MultiPath {
	parts := make([]Path, 0, len(mp.paths)+len(paths))
	parts = append(parts, mp.paths...)
	parts = append(parts, paths...)
	return NewMultiPath(parts...)
}

func (mp MultiPath) clone() Element {
	out := &MultiPath{
		tagBase: mp.tagBase,
		paths:   make([]Path, len(mp.paths)),
	}
	for i, path := range mp.paths {
		out.paths[i] = clone(path)
	}
	return out
}

func (mp MultiPath) translate(dx, dy float64) {
	for _, path := range mp.paths {
		path.translate(dx, dy)
	}
}

func (mp MultiPath) pathIsAClosedType() {}
