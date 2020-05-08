package ngin2d

import (
	"testing"

	"github.com/c2fo/testify/require"
)

func Test_Distance(t *testing.T) {
	req := require.New(t)

	testCases := []struct {
		x1   float64
		x2   float64
		y1   float64
		y2   float64
		dist float64
	}{
		{
			x1:   0,
			y1:   0,
			x2:   0,
			y2:   0,
			dist: 0,
		},
		{
			x1:   0,
			y1:   0,
			x2:   3,
			y2:   4,
			dist: 5,
		},
		{
			x1:   -2.5,
			y1:   0,
			x2:   3.5,
			y2:   8,
			dist: 10,
		},
	}

	for _, tCase := range testCases {
		req.Equal(tCase.dist, distance(tCase.x1, tCase.y1, tCase.x2, tCase.y2), "Distance should be equal")
	}
}
