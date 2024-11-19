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
