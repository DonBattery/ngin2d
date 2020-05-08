package ngin2d

type Rect struct {
	pos *Vector
	h   float64
	w   float64
}

func NewRect(x, y, w, h float64) *Rect {
	return &Rect{
		pos: NewVector(x, y),
		h:   h,
		w:   w,
	}
}
