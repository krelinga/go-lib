package geom

type CircleArc struct {
	tagBase
	center               *Point
	radius               float64
	rayAngle             RayAngle
	startPoint, endPoint *Point
}

func NewCircleArc(center *Point, radius float64, rayAngle RayAngle) *CircleArc {
	pt := func(r Ray) *Point {
		dx, dy := r.DxDy(radius)
		return NewPoint(center.X()+dx, center.Y()+dy)
	}
	c := &CircleArc{
		center:     center,
		radius:     radius,
		rayAngle:   rayAngle,
		startPoint: pt(rayAngle.from),
		endPoint:   pt(rayAngle.to),
	}
	c.addChildTags(center.getTagIndex())
	return c
}

func (c *CircleArc) Center() *Point {
	return c.center
}

func (c *CircleArc) Radius() float64 {
	return c.radius
}

func (c *CircleArc) RayAngle() RayAngle {
	return c.rayAngle
}

func (c *CircleArc) Endpoints() (*Point, *Point) {
	return c.startPoint, c.endPoint
}

func (c *CircleArc) BoundingBox() BoundingBox {
	minX, maxX := c.startPoint.X(), c.startPoint.X()
	minY, maxY := c.startPoint.Y(), c.startPoint.Y()
	update := func(x, y float64) {
		minX = min(minX, x)
		maxX = max(maxX, x)
		minY = min(minY, y)
		maxY = max(maxY, y)
	}
	update(c.endPoint.X(), c.endPoint.Y())
	update(c.center.X(), c.center.Y())
	if c.rayAngle.Includes(RayUp) {
		update(c.center.X(), c.center.Y()+c.radius)
	}
	if c.rayAngle.Includes(RayRight) {
		update(c.center.X()+c.radius, c.center.Y())
	}
	if c.rayAngle.Includes(RayDown) {
		update(c.center.X(), c.center.Y()-c.radius)
	}
	if c.rayAngle.Includes(RayLeft) {
		update(c.center.X()-c.radius, c.center.Y())
	}
	return BoundingBox{
		Min: NewPoint(minX, minY),
		Max: NewPoint(maxX, maxY),
	}
}

func (c *CircleArc) clone() Element {
	out := &CircleArc{
		center:     clone(c.center),
		radius:     c.radius,
		rayAngle:   c.rayAngle,
		startPoint: clone(c.startPoint),
		endPoint:   clone(c.endPoint),
	}
	out.addChildTags(out.center.getTagIndex())
	out.addChildTags(out.startPoint.getTagIndex())
	out.addChildTags(out.endPoint.getTagIndex())
	out.addSelfTags(out, c.getSelfTags())
	return out
}

func (c *CircleArc) translate(dx, dy float64) {
	c.center.translate(dx, dy)
	c.startPoint.translate(dx, dy)
	c.endPoint.translate(dx, dy)
}

func (c *CircleArc) rotate(angle Angle, dir Direction) {
	c.center.rotate(angle, dir)
	c.rayAngle = c.rayAngle.Rotate(angle, dir)
	c.startPoint.rotate(angle, dir)
	c.endPoint.rotate(angle, dir)
}

func (c *CircleArc) pathIsAClosedType() {}
