package geom

type Point struct {
	tagBase
	x, y float64
}

func NewPoint(x, y float64, tags ...*PointTag) *Point {
	p := &Point{x: x, y: y}
	addTags(&p.tagBase, p, tags...)
	return p
}

func (p *Point) X() float64 {
	return p.x
}

func (p *Point) Y() float64 {
	return p.y
}

func (p *Point) BoundingBox() BoundingBox {
	return BoundingBox{
		Min: p,
		Max: p,
	}
}

func (p *Point) Endpoints() (*Point, *Point) {
	return p, p
}

func (p *Point) pathIsAClosedType() {}
