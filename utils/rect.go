package utils

import (
	"fmt"
	"math"
)

type Rect struct {
	X, Y, W, H int
}

func (r Rect) Left() int    { return r.X }
func (r Rect) Top() int     { return r.Y }
func (r Rect) Right() int   { return r.X + r.W }
func (r Rect) Bottom() int  { return r.Y + r.H }
func (r Rect) CenterX() int { return r.X + r.W/2 }
func (r Rect) CenterY() int { return r.Y + r.H/2 }

func (r Rect) MoveTo(x, y int) Rect {
	r.X = x
	r.Y = y
	return r
}

func (r Rect) ResizeTo(w, h int) Rect {
	r.W = w
	r.H = h
	return r
}

func (r Rect) ShrinkByInsets(i Insets) Rect {
	return Rect{
		r.X + i.Left,
		r.Y + i.Top,
		r.W - i.Right - i.Left,
		r.H - i.Bottom - i.Top,
	}
}

func (r Rect) ShrinkByInt(i int) Rect {
	return Rect{
		r.X + i,
		r.Y + i,
		r.W - i*2,
		r.H - i*2,
	}
}

func (r Rect) CenterIn(o Rect) Rect {
	hW := (o.W - r.W) / 2
	hH := (o.H - r.H) / 2
	return Rect{
		r.X + hW,
		r.Y + hH,
		r.W + hW*2,
		r.H + hH*2,
	}
}

func (r Rect) ContainsPoint(pointX, pointY int) bool {
	return containsPoint(r.X, r.Y, r.W, r.H, pointX, pointY)
}

func (r Rect) ContainsRect(other Rect) bool {
	return containsRect(r.X, r.Y, r.W, r.H, other.X, other.Y, other.W, other.H)
}

func (r Rect) IntersectsRect(other Rect) bool {
	return intersectsRect(r.X, r.Y, r.W, r.H, other.X, other.Y, other.W, other.H)
}

func containsPoint(x, y, w, h, pointX, pointY int) bool {
	if w <= 0 || h <= 0 {
		return false
	}
	if pointX < x || pointY < y {
		return false
	}
	w += x
	h += y
	//    overflow || intersect
	return (w < x || w >= pointX) && (h < y || h >= pointY)
}

func containsRect(x, y, w, h, rx, ry, rw, rh int) bool {
	if (rw | rh | w | h) <= 0 {
		return false
	}
	// Note: if any dimension is zero, tests below must return false...
	if rx < x || ry < y {
		return false
	}
	w += x
	rw += rx
	if rw <= rx {
		// rx+rw overflowed or rw was zero, return false if...
		// either original w or rw was zero or
		// x+w did not overflow or
		// the overflowed x+w is smaller than the overflowed rx+rw
		if w >= x || rw > w {
			return false
		}
	} else {
		// rx+rw did not overflow and rw was not zero, return false if...
		// original w was zero or
		// x+w did not overflow and x+w is smaller than rx+rw
		if w >= x && rw > w {
			return false
		}
	}
	h += y
	rh += ry
	if rh <= ry {
		if h >= y || rh > h {
			return false
		}
	} else {
		if h >= y && rh > h {
			return false
		}
	}
	return true
}

func intersectsRect(x, y, w, h, rx, ry, rw, rh int) bool {
	if (rw | rh | w | h) <= 0 {
		return false
	}
	rw += rx
	rh += ry
	w += x
	h += y
	//      overflow || intersect
	return (rw < rx || rw > x) &&
		(rh < ry || rh > y) &&
		(w < x || w > rx) &&
		(h < y || h > ry)
}

func (r Rect) IntersectionWith(other Rect) Rect {
	tx1 := r.X
	ty1 := r.Y
	rx1 := other.X
	ry1 := other.Y
	tx2 := tx1
	tx2 += r.W
	ty2 := ty1
	ty2 += r.H
	rx2 := rx1
	rx2 += other.W
	ry2 := ry1
	ry2 += other.H
	if tx1 < rx1 {
		tx1 = rx1
	}
	if ty1 < ry1 {
		ty1 = ry1
	}
	if tx2 > rx2 {
		tx2 = rx2
	}
	if ty2 > ry2 {
		ty2 = ry2
	}
	tx2 -= tx1
	ty2 -= ty1
	// tx2,ty2 will never overflow (they will never be
	// larger than the smallest of the two source w,h)
	// they might underflow, though...
	if tx2 < math.MinInt32 {
		tx2 = math.MinInt32
	}
	if ty2 < math.MinInt32 {
		ty2 = math.MinInt32
	}
	return Rect{tx1, ty1, tx2, ty2}
}

func (r Rect) EqualsTo(other Rect) bool {
	return r.X == other.X && r.Y == other.Y && r.W == other.W && r.H == other.H
}

func (r Rect) ToString() string {
	return fmt.Sprintf("{x:%d,y:%d,w:%d,h:%d}", r.X, r.Y, r.W, r.H)
}
