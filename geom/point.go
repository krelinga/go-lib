package geom

type Point struct {
	tagBase
	x, y float64
}

func NewPoint(x, y float64, tags ...*PointTag) *Point {
	p := &Point{x: x, y: y}
	p.addPublicTags(p, toPublicTagArray(tags))
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

func (p *Point) clone() Element {
	out := &Point{
		x: p.x,
		y: p.y,
	}
	out.addSelfTags(out, p.getSelfTags())
	return out
}

func (p *Point) translate(dx, dy float64) {
	p.x += dx
	p.y += dy
}

func (p *Point) rotate(angle Angle, dir Direction) {
	cos := Cos(angle)
	sin := Sin(angle)

	if dir == Clockwise {
		sin = -sin
		cos = -cos
	}

	newX := p.x*cos - p.y*sin
	newY := p.x*sin + p.y*cos

	p.x = newX
	p.y = newY
}

func (p *Point) pathIsAClosedType() {}
