package ngin2d

import (
	"testing"

	"github.com/c2fo/testify/require"
)

func Test_Vector_Creation_XY(t *testing.T) {
	req := require.New(t)

	v := NewVector(0, 0)
	req.NotNil(v, "The Vector should have been created")
	req.True(v.IsNull(), "v.IsNull should be true for a new vector")
	req.Equal(float64(0), v.GetX(), "v.x should be 0 of the new vector")
	req.Equal(float64(0), v.GetY(), "v.y should be 0 of the new vector")
	x, y := v.GetXY()
	req.Equal(float64(0), x, "v.x should be 0 of the new vector")
	req.Equal(float64(0), y, "v.y should be 0 of the new vector")

	v.SetXY(44, -12.786)
	x = v.GetX()
	y = v.GetY()
	req.Equal(float64(44), x, "v.x should be 44 after v.SetXY(44, -12.786)")
	req.Equal(float64(-12.786), y, "v.y should be -12.786 after v.SetXY(44, -12.786)")
	v.SetX(0.00001)
	v.SetY(123)
	x, y = v.GetXY()
	req.Equal(float64(0.00001), x, "v.x should be 0.00001 after v.SetX(0.00001)")
	req.Equal(float64(123), y, "v.y should be 123 after v.SetY(123)")
}

// func Test_Vector_Creation_XY(t *testing.T) {
// 	req := require.New(t)

// 	v := NewVector(0, 0)
// }
