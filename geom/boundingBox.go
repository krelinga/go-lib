package geom

type BoundingBox struct {
	Min, Max *Point
}

func (bb BoundingBox) Width() float64 {
	return bb.Max.X() - bb.Min.X()
}

func (bb BoundingBox) Height() float64 {
	return bb.Max.Y() - bb.Min.Y()
}

func (bb BoundingBox) Center() *Point {
	return NewPoint(
		(bb.Min.X() + bb.Max.X()) / 2,
		(bb.Min.Y() + bb.Max.Y()) / 2,
	)
}

func (bb BoundingBox) TopLeft() *Point {
	return NewPoint(bb.Min.X(), bb.Max.Y())
}

func (bb BoundingBox) TopRight() *Point {
	return bb.Max
}

func (bb BoundingBox) BottomLeft() *Point {
	return bb.Min
}

func (bb BoundingBox) BottomRight() *Point {
	return NewPoint(bb.Max.X(), bb.Min.Y())
}
