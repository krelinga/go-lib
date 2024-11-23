package geom

import "math"

type RectangleOpt interface {
	Rectangle(*rectangleOpts)
}

type rectangleOpts struct {
	topLineTags, bottomLineTags, rightLineTags, leftLineTags                                   []*LineTag
	topLeftPointTags, topRightPointTags, bottomRightPointTags, bottomLeftPointTags             []*PointTag
	topLeftCornerRadius, topRightCornerRadius, bottomRightCornerRadius, bottomLeftCornerRadius float64
}

func NewRectangle(width, height float64, inOpts ...RectangleOpt) *PathFigure {
	// TODO: bounds checking for corner radii vs. width/height.
	// TOOD: figure out how to handle corner point tags for rounded corners.
	opts := &rectangleOpts{}
	for _, opt := range inOpts {
		opt.Rectangle(opts)
	}
	xRad := width / 2
	yRad := height / 2
	corner := func(x, y, rad float64, start, end Ray) *CircleArc {
		if rad == 0 {
			return nil
		}
		c := NewPoint(x-math.Copysign(rad, x), y-math.Copysign(rad, y))
		return NewCircleArc(c, rad, NewRayAngle(start, end, Clockwise))
	}
	tl := corner(-xRad, yRad, opts.topLeftCornerRadius, RayLeft, RayUp)
	tr := corner(xRad, yRad, opts.topRightCornerRadius, RayUp, RayRight)
	br := corner(xRad, -yRad, opts.bottomRightCornerRadius, RayRight, RayDown)
	bl := corner(-xRad, -yRad, opts.bottomLeftCornerRadius, RayDown, RayLeft)
	side := func(c1, c2 *CircleArc, p1, p2 *Point, lineTags []*LineTag) *Line {
		var start, stop *Point
		if c1 != nil {
			_, start = c1.Endpoints()
		} else {
			start = p1
		}
		if c2 != nil {
			stop, _ = c2.Endpoints()
		} else {
			stop = p2
		}
		return NewLine(start, stop, lineTags...)
	}
	t := side(tl, tr, NewPoint(-xRad, yRad, opts.topLeftPointTags...), NewPoint(xRad, yRad, opts.topRightPointTags...), opts.topLineTags)
	r := side(tr, br, NewPoint(xRad, yRad, opts.topRightPointTags...), NewPoint(xRad, -yRad, opts.bottomRightPointTags...), opts.rightLineTags)
	b := side(br, bl, NewPoint(xRad, -yRad, opts.bottomRightPointTags...), NewPoint(-xRad, -yRad, opts.bottomLeftPointTags...), opts.bottomLineTags)
	l := side(bl, tl, NewPoint(-xRad, -yRad, opts.bottomLeftPointTags...), NewPoint(-xRad, yRad, opts.topLeftPointTags...), opts.leftLineTags)

	args := make([]Path, 0, 8)
	addArg := func(p Path) {
		var isNil bool
		switch typedP := p.(type) {
		case *Line:
			isNil = typedP == nil
		case *CircleArc:
			isNil = typedP == nil
		}
		if isNil {
			return
		}
		args = append(args, p)
	}
	addArg(tl)
	addArg(t)
	addArg(tr)
	addArg(r)
	addArg(br)
	addArg(b)
	addArg(bl)
	addArg(l)

	return NewPathFigure(args...)
}

func (opt TopLineTagOpt) Rectangle(r *rectangleOpts) {
	r.topLineTags = append(r.topLineTags, opt.lineTag)
}

func (opt BottomLineTagOpt) Rectangle(r *rectangleOpts) {
	r.bottomLineTags = append(r.bottomLineTags, opt.lineTag)
}

func (opt LeftLineTagOpt) Rectangle(r *rectangleOpts) {
	r.leftLineTags = append(r.leftLineTags, opt.lineTag)
}

func (opt RightLineTagOpt) Rectangle(r *rectangleOpts) {
	r.rightLineTags = append(r.rightLineTags, opt.lineTag)
}

func (opt TopLeftPointTagOpt) Rectangle(r *rectangleOpts) {
	r.topLeftPointTags = append(r.topLeftPointTags, opt.pointTag)
}

func (opt TopRightPointTagOpt) Rectangle(r *rectangleOpts) {
	r.topRightPointTags = append(r.topRightPointTags, opt.pointTag)
}

func (opt BottomRightPointTagOpt) Rectangle(r *rectangleOpts) {
	r.bottomRightPointTags = append(r.bottomRightPointTags, opt.pointTag)
}

func (opt BottomLeftPointTagOpt) Rectangle(r *rectangleOpts) {
	r.bottomLeftPointTags = append(r.bottomLeftPointTags, opt.pointTag)
}

func (opt RoundAllCornersOpt) Rectangle(r *rectangleOpts) {
	r.topLeftCornerRadius = float64(opt)
	r.topRightCornerRadius = float64(opt)
	r.bottomRightCornerRadius = float64(opt)
	r.bottomLeftCornerRadius = float64(opt)
}

func (opt RoundTopRightCornerOpt) Rectangle(r *rectangleOpts) {
	r.topRightCornerRadius = float64(opt)
}

func (opt RoundBottomRightCornerOpt) Rectangle(r *rectangleOpts) {
	r.bottomRightCornerRadius = float64(opt)
}

func (opt RoundBottomLeftCornerOpt) Rectangle(r *rectangleOpts) {
	r.bottomLeftCornerRadius = float64(opt)
}

func (opt RoundTopLeftCornerOpt) Rectangle(r *rectangleOpts) {
	r.topLeftCornerRadius = float64(opt)
}
