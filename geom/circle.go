package geom

type Circle struct {
	tagBase
	center *Point
	radius float64
}

func NewCircle(center *Point, radius float64, tags ...*CircleTag) *Circle {
	c := &Circle{center: center, radius: radius}
	c.addChildTags(center.getTagIndex())
	c.addPublicTags(c, toPublicTagArray(tags))
	return c
}

func (c *Circle) Center() *Point {
	return c.center
}

func (c *Circle) Radius() float64 {
	return c.radius
}

func (c *Circle) BoundingBox() BoundingBox {
	return BoundingBox{
		Min: NewPoint(c.center.X()-c.radius, c.center.Y()-c.radius),
		Max: NewPoint(c.center.X()+c.radius, c.center.Y()+c.radius),
	}
}

func (c *Circle) clone() Element {
	out := &Circle{
		center: clone(c.center),
		radius: c.radius,
	}
	out.addChildTags(out.center.getTagIndex())
	out.addSelfTags(out, c.getSelfTags())
	return out
}

func (c *Circle) translate(dx, dy float64) {
	c.center.translate(dx, dy)
}

func (c *Circle) rotate(angle Angle, dir Direction) {
	c.center.rotate(angle, dir)
}

func (c *Circle) figureIsAClosedType() {}