package geom

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRectangle(t *testing.T) {
	t.Run("Square Corners", func(t *testing.T) {
		var tl, tr, br, bl PointTag
		var top, right, bottom, left LineTag
		args := []RectangleOpt{
			TagTopLine(&top),
			TagRightLine(&right),
			TagBottomLine(&bottom),
			TagLeftLine(&left),
			TagTopRightPoint(&tr),
			TagTopLeftPoint(&tl),
			TagBottomRightPoint(&br),
			TagBottomLeftPoint(&bl),
		}
		r := NewRectangle(10, 20, args...)

		// Lines
		checkLine := func(lt LineTag, ep1, ep2 *Point) {
			if !assert.NotNil(t, lt.Get(r)) {
				return
			}
			p1, p2 := lt.Get(r).Endpoints()
			assertXYEqual(t, ep1, p1)
			assertXYEqual(t, ep2, p2)
		}
		checkLine(top, NewPoint(-5, 10), NewPoint(5, 10))
		checkLine(right, NewPoint(5, 10), NewPoint(5, -10))
		checkLine(bottom, NewPoint(5, -10), NewPoint(-5, -10))
		checkLine(left, NewPoint(-5, -10), NewPoint(-5, 10))

		checkPoint := func(pt PointTag, ep *Point) {
			if !assert.NotNil(t, pt.Get(r)) {
				return
			}
			assertXYEqual(t, ep, pt.Get(r))
		}
		checkPoint(tl, NewPoint(-5, 10))
		checkPoint(tr, NewPoint(5, 10))
		checkPoint(br, NewPoint(5, -10))
		checkPoint(bl, NewPoint(-5, -10))
	})
	t.Run("Rounded Corners", func(t *testing.T) {
		var top, right, bottom, left LineTag
		args := []RectangleOpt{
			TagTopLine(&top),
			TagRightLine(&right),
			TagBottomLine(&bottom),
			TagLeftLine(&left),
			RoundTopLeftCorner(1),
			RoundTopRightCorner(2),
			RoundBottomRightCorner(3),
			RoundBottomLeftCorner(4),
		}
		r := NewRectangle(10, 20, args...)

		checkLine := func(lt LineTag, ep1, ep2 *Point) {
			if !assert.NotNil(t, lt.Get(r)) {
				return
			}
			p1, p2 := lt.Get(r).Endpoints()
			assertXYEqual(t, ep1, p1)
			assertXYEqual(t, ep2, p2)
		}
		checkLine(top, NewPoint(-4, 10), NewPoint(3, 10))
		checkLine(right, NewPoint(5, 8), NewPoint(5, -7))
		checkLine(bottom, NewPoint(2, -10), NewPoint(-1, -10))
		checkLine(left, NewPoint(-5, -6), NewPoint(-5, 9))
	})

	t.Run("Round All Corners", func(t *testing.T) {
		var top, right, bottom, left LineTag
		args := []RectangleOpt{
			TagTopLine(&top),
			TagRightLine(&right),
			TagBottomLine(&bottom),
			TagLeftLine(&left),
			RoundAllCorners(2),
		}
		r := NewRectangle(10, 20, args...)

		checkLine := func(lt LineTag, ep1, ep2 *Point) {
			if !assert.NotNil(t, lt.Get(r)) {
				return
			}
			p1, p2 := lt.Get(r).Endpoints()
			assertXYEqual(t, ep1, p1)
			assertXYEqual(t, ep2, p2)
		}
		checkLine(top, NewPoint(-3, 10), NewPoint(3, 10))
		checkLine(right, NewPoint(5, 8), NewPoint(5, -8))
		checkLine(bottom, NewPoint(3, -10), NewPoint(-3, -10))
		checkLine(left, NewPoint(-5, -8), NewPoint(-5, 8))
	})
}
