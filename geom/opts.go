package geom

// Add tags to specific lines.
type lineTagOpt struct {
	lineTag *LineTag
}

type TopLineTagOpt lineTagOpt

func TagTopLine(t *LineTag) TopLineTagOpt {
	return TopLineTagOpt{t}
}

type TopRightLineTagOpt lineTagOpt

func TagTopRightLine(t *LineTag) TopRightLineTagOpt {
	return TopRightLineTagOpt{t}
}

type BottomRightLineTagOpt lineTagOpt

func TagBottomRightLine(t *LineTag) BottomRightLineTagOpt {
	return BottomRightLineTagOpt{t}
}

type BottomLineTagOpt lineTagOpt

func TagBottomLine(t *LineTag) BottomLineTagOpt {
	return BottomLineTagOpt{t}
}

type BottomLeftLineTagOpt lineTagOpt

func TagBottomLeftLine(t *LineTag) BottomLeftLineTagOpt {
	return BottomLeftLineTagOpt{t}
}

type TopLeftLineTagOpt lineTagOpt

func TagTopLeftLine(t *LineTag) TopLeftLineTagOpt {
	return TopLeftLineTagOpt{t}
}

type RightLineTagOpt lineTagOpt

func TagRightLine(t *LineTag) RightLineTagOpt {
	return RightLineTagOpt{t}
}

type LeftLineTagOpt lineTagOpt

func TagLeftLine(t *LineTag) LeftLineTagOpt {
	return LeftLineTagOpt{t}
}

// Add tags to specific points.
type pointTagOpt struct {
	pointTag *PointTag
}

type TopRightPointTagOpt pointTagOpt

func TagTopRightPoint(t *PointTag) TopRightPointTagOpt {
	return TopRightPointTagOpt{t}
}

type RightPointTagOpt pointTagOpt

func TagRightPoint(t *PointTag) RightPointTagOpt {
	return RightPointTagOpt{t}
}

type BottomRightPointTagOpt pointTagOpt

func TagBottomRightPoint(t *PointTag) BottomRightPointTagOpt {
	return BottomRightPointTagOpt{t}
}

type BottomLeftPointTagOpt pointTagOpt

func TagBottomLeftPoint(t *PointTag) BottomLeftPointTagOpt {
	return BottomLeftPointTagOpt{t}
}

type LeftPointTagOpt pointTagOpt

func TagLeftPoint(t *PointTag) LeftPointTagOpt {
	return LeftPointTagOpt{t}
}

type TopLeftPointTagOpt pointTagOpt

func TagTopLeftPoint(t *PointTag) TopLeftPointTagOpt {
	return TopLeftPointTagOpt{t}
}

type CenterPointTagOpt pointTagOpt

func TagCenterPoint(t *PointTag) CenterPointTagOpt {
	return CenterPointTagOpt{t}
}

// Options for rounding corners.

type roundCornerOpt float64

type RoundAllCornersOpt roundCornerOpt

func RoundAllCorners(radius float64) RoundAllCornersOpt {
	return RoundAllCornersOpt(radius)
}

type RoundTopLeftCornerOpt roundCornerOpt

func RoundTopLeftCorner(radius float64) RoundTopLeftCornerOpt {
	return RoundTopLeftCornerOpt(radius)
}

type RoundTopRightCornerOpt roundCornerOpt

func RoundTopRightCorner(radius float64) RoundTopRightCornerOpt {
	return RoundTopRightCornerOpt(radius)
}

type RoundBottomRightCornerOpt roundCornerOpt

func RoundBottomRightCorner(radius float64) RoundBottomRightCornerOpt {
	return RoundBottomRightCornerOpt(radius)
}

type RoundBottomLeftCornerOpt roundCornerOpt

func RoundBottomLeftCorner(radius float64) RoundBottomLeftCornerOpt {
	return RoundBottomLeftCornerOpt(radius)
}
