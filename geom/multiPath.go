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
		mp.addChildTags(path.getTagIndex())
	}
	return mp
}

func (mp *MultiPath) BoundingBox() BoundingBox {
	minX := mp.paths[0].BoundingBox().BottomLeft().X()
	minY := mp.paths[0].BoundingBox().BottomLeft().Y()
	maxX := mp.paths[0].BoundingBox().TopRight().X()
	maxY := mp.paths[0].BoundingBox().TopRight().Y()
	for _, path := range mp.paths[1:] {
		bb := path.BoundingBox()
		minX = min(minX, bb.BottomLeft().X())
		minY = min(minY, bb.BottomLeft().Y())
		maxX = max(maxX, bb.TopRight().X())
		maxY = max(maxY, bb.TopRight().Y())
	}
	return BoundingBox{Min: NewPoint(minX, minY), Max: NewPoint(maxX, maxY)}
}

func (mp *MultiPath) Endpoints() (*Point, *Point) {
	start, _ := mp.paths[0].Endpoints()
	_, end := mp.paths[len(mp.paths)-1].Endpoints()
	return start, end
}

func (mu *MultiPath) Paths() iter.Seq[Path] {
	return slices.Values(mu.paths)
}

func (mp *MultiPath) Extend(paths ...Path) *MultiPath {
	parts := make([]Path, 0, len(mp.paths)+len(paths))
	parts = append(parts, mp.paths...)
	parts = append(parts, paths...)
	return NewMultiPath(parts...)
}

func (mp *MultiPath) clone() Element {
	out := &MultiPath{
		paths: make([]Path, len(mp.paths)),
	}
	for i, path := range mp.paths {
		out.paths[i] = clone(path)
		out.addChildTags(out.paths[i].getTagIndex())
	}
	return out
}

func (mp *MultiPath) translate(dx, dy float64) {
	for _, path := range mp.paths {
		path.translate(dx, dy)
	}
}

func (mp *MultiPath) rotate(angle Angle, dir Direction) {
	for _, path := range mp.paths {
		path.rotate(angle, dir)
	}
}

func (mp *MultiPath) pathIsAClosedType() {}
