package geom

type HexagonOpt interface {
	Hexagon(*hexagonTags)
}

func Hexagon(r float64, opts ...HexagonOpt) *Figure {
	ht := hexagonTags{}
	for _, opt := range opts {
		opt.Hexagon(&ht)
	}

	rightPoint := NewPoint(r, 0, ht.rightPointTags...)
	bottomRightPoint := NewPoint(r/2, -r*Sin(Degrees(60)), ht.bottomRightPointTags...)
	bottomLeftPoint := NewPoint(-r/2, -r*Sin(Degrees(60)), ht.bottomLeftPointTags...)
	leftPoint := NewPoint(-r, 0, ht.leftPointTags...)
	topLeftPoint := NewPoint(-r/2, r*Sin(Degrees(60)), ht.topLeftPointTags...)
	topRightPoint := NewPoint(r/2, r*Sin(Degrees(60)), ht.topRightPointTags...)

	topLine := NewLine(topLeftPoint, topRightPoint, ht.topLineTags...)
	topRightLine := NewLine(topRightPoint, rightPoint, ht.topRightLineTags...)
	bottomRightLine := NewLine(rightPoint, bottomRightPoint, ht.bottomRightLineTags...)
	bottomLine := NewLine(bottomRightPoint, bottomLeftPoint, ht.bottomLineTags...)
	bottomLeftLine := NewLine(bottomLeftPoint, leftPoint, ht.bottomLeftLineTags...)
	topLeftLine := NewLine(leftPoint, topLeftPoint, ht.topLeftLineTags...)

	f := NewFigure(topLine, topRightLine, bottomRightLine, bottomLine, bottomLeftLine, topLeftLine)
	if f == nil {
		panic("Hexagon should be closed")
	}
	return f
}

type hexagonTags struct {
	topLineTags, topRightLineTags, bottomRightLineTags, bottomLineTags, bottomLeftLineTags, topLeftLineTags       []*LineTag
	topRightPointTags, bottomRightPointTags, bottomLeftPointTags, topLeftPointTags, rightPointTags, leftPointTags []*PointTag
}

func (opt TopLineTagOpt) Hexagon(h *hexagonTags) {
	h.topLineTags = append(h.topLineTags, opt.lineTag)
}

func (opt TopRightLineTagOpt) Hexagon(h *hexagonTags) {
	h.topRightLineTags = append(h.topRightLineTags, opt.lineTag)
}

func (opt BottomRightLineTagOpt) Hexagon(h *hexagonTags) {
	h.bottomRightLineTags = append(h.bottomRightLineTags, opt.lineTag)
}

func (opt BottomLineTagOpt) Hexagon(h *hexagonTags) {
	h.bottomLineTags = append(h.bottomLineTags, opt.lineTag)
}

func (opt BottomLeftLineTagOpt) Hexagon(h *hexagonTags) {
	h.bottomLeftLineTags = append(h.bottomLeftLineTags, opt.lineTag)
}

func (opt TopLeftLineTagOpt) Hexagon(h *hexagonTags) {
	h.topLeftLineTags = append(h.topLeftLineTags, opt.lineTag)
}

func (opt TopRightPointTagOpt) Hexagon(h *hexagonTags) {
	h.topRightPointTags = append(h.topRightPointTags, opt.pointTag)
}

func (opt BottomRightPointTagOpt) Hexagon(h *hexagonTags) {
	h.bottomRightPointTags = append(h.bottomRightPointTags, opt.pointTag)
}

func (opt BottomLeftPointTagOpt) Hexagon(h *hexagonTags) {
	h.bottomLeftPointTags = append(h.bottomLeftPointTags, opt.pointTag)
}

func (opt TopLeftPointTagOpt) Hexagon(h *hexagonTags) {
	h.topLeftPointTags = append(h.topLeftPointTags, opt.pointTag)
}

func (opt RightPointTagOpt) Hexagon(h *hexagonTags) {
	h.rightPointTags = append(h.rightPointTags, opt.pointTag)
}

func (opt LeftPointTagOpt) Hexagon(h *hexagonTags) {
	h.leftPointTags = append(h.leftPointTags, opt.pointTag)
}
