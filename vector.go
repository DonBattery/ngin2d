package ngin2d

import "math"

type Vector struct {
	x float64
	y float64
}

func NewVector(x, y float64) *Vector {
	return &Vector{
		x: x,
		y: y,
	}
}

func (v *Vector) IsNull() bool {
	return v.x == 0 && v.y == 0
}

func (v *Vector) Len() float64 {
	return math.Sqrt(math.Pow(v.x, 2) + math.Pow(v.y, 2))
}

func (v *Vector) SetX(x float64) {
	v.x = x
}

func (v *Vector) SetY(y float64) {
	v.y = y
}

func (v *Vector) SetXY(x, y float64) {
	v.x = x
	v.y = y
}

func (v *Vector) GetX() float64 {
	return v.x
}

func (v *Vector) GetY() float64 {
	return v.y
}

func (v *Vector) GetXY() (float64, float64) {
	return v.x, v.y
}

func (v *Vector) Copy() *Vector {
	return NewVector(v.x, v.y)
}

func (v *Vector) AddX(x float64) {
	v.x += x
}

func (v *Vector) AddedX(x float64) *Vector {
	v2 := v.Copy()
	v2.AddX(x)
	return v2
}

func (v *Vector) AddY(y float64) {
	v.y += y
}

func (v *Vector) AddedY(y float64) *Vector {
	v2 := v.Copy()
	v2.AddY(y)
	return v2
}

func (v *Vector) Add(other *Vector) {
	v.x += other.x
	v.y += other.y
}

func (v *Vector) Added(other *Vector) *Vector {
	v2 := v.Copy()
	v2.Add(other)
	return v2
}

func (v *Vector) Sub(other *Vector) {
	v.x -= other.x
	v.y -= other.y
}

func (v *Vector) Subbed(other *Vector) *Vector {
	v2 := v.Copy()
	v2.Sub(other)
	return v2
}

func (v *Vector) Reverse() {
	v.x *= -1
	v.y *= -1
}

func (v *Vector) Reversed() *Vector {
	v2 := v.Copy()
	v2.Reverse()
	return v2
}

func (v *Vector) Normal() {
	len := v.Len()
	v.x /= len
	v.y /= len
}

func (v *Vector) Normalized() *Vector {
	v2 := v.Copy()
	v2.Normal()
	return v2
}

func (v *Vector) SetLen(len float64) {
	v.Scale(len / v.Len())
}

func (v *Vector) SettedLen(len float64) *Vector {
	v2 := v.Copy()
	v2.SetLen(len)
	return v2
}

func (v *Vector) ScaleToX(scaleTo float64) {
	v.y = v.y * scaleTo / v.x
	v.x = scaleTo
}

func (v *Vector) ScaledToX(scaleTo float64) *Vector {
	v2 := v.Copy()
	v2.ScaleToX(scaleTo)
	return v2
}

func (v *Vector) ScaleToY(scaleTo float64) {
	v.x = v.x * scaleTo / v.y
	v.y = scaleTo
}

func (v *Vector) ScaledToY(scaleTo float64) *Vector {
	v2 := v.Copy()
	v2.ScaleToY(scaleTo)
	return v2
}

func (v *Vector) Scale(size float64) {
	v.x *= size
	v.y *= size
}

func (v *Vector) Scaled(size float64) *Vector {
	v2 := v.Copy()
	v2.Scale(size)
	return v2
}
