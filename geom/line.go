package geom

type Line struct {
	tagBase
	p1, p2 *Point
}

func NewLine(p1, p2 *Point, tags ...*LineTag) *Line {
	l := &Line{p1: p1, p2: p2}
	l.addChildTags(p1.getTagIndex())
	l.addChildTags(p2.getTagIndex())
	l.addPublicTags(l, toPublicTagArray(tags))
	return l
}

func (l *Line) BoundingBox() BoundingBox {
	minX := min(l.p1.X(), l.p2.X())
	minY := min(l.p1.Y(), l.p2.Y())
	maxX := max(l.p1.X(), l.p2.X())
	maxY := max(l.p1.Y(), l.p2.Y())

	var minPoint, maxPoint *Point
	if minX == l.p1.X() && minY == l.p1.Y() {
		minPoint = l.p1
	} else if minX == l.p2.X() && minY == l.p2.Y() {
		minPoint = l.p2
	} else {
		minPoint = NewPoint(minX, minY)
	}
	if maxX == l.p1.X() && maxY == l.p1.Y() {
		maxPoint = l.p1
	} else if maxX == l.p2.X() && maxY == l.p2.Y() {
		maxPoint = l.p2
	} else {
		maxPoint = NewPoint(maxX, maxY)
	}

	return BoundingBox{Min: minPoint, Max: maxPoint}
}

func (l *Line) Endpoints() (*Point, *Point) {
	return l.p1, l.p2
}

func (l *Line) clone() Element {
	out := &Line{}
	out.p1 = clone(l.p1)
	out.addChildTags(out.p1.getTagIndex())
	out.p2 = clone(l.p2)
	out.addChildTags(out.p2.getTagIndex())
	out.addSelfTags(out, l.getSelfTags())
	return out
}

func (l *Line) translate(dx, dy float64) {
	l.p1.translate(dx, dy)
	l.p2.translate(dx, dy)
}

func (l *Line) pathIsAClosedType() {}
