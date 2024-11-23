package geom

type Circle struct {
	tagBase
	center *Point
	radius float64
}

type CircleOpt interface {
	Circle(*circleOpts)
}

type circleOpts struct {
	centerPointTags []*PointTag
}

func NewCircle(radius float64, inOpts... CircleOpt) *Circle {
	opts := &circleOpts{}
	for _, o := range inOpts {
		o.Circle(opts)
	}
	center := NewPoint(0, 0, opts.centerPointTags...)
	c := &Circle{center: center, radius: radius}
	c.addChildTags(center.getTagIndex())
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

func (t CenterPointTagOpt) Circle(opts *circleOpts) {
	opts.centerPointTags = append(opts.centerPointTags, t.pointTag)
}