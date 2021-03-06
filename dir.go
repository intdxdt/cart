package cart

import (
	"github.com/intdxdt/math"
)

//Dir computes direction in radians - counter clockwise from x-axis.
func Direction(x, y float64) float64 {
	d := math.Atan2(y, x)
	if d < 0 {
		d += math.Tau
	}
	return d
}

//Revdir computes the reversed direction from a foward direction
func ReverseDirection(d float64) float64 {
	if d < math.Pi {
		return d + math.Pi
	}
	return d - math.Pi
}

func DeflectionAngle(bearing1, bearing2 float64) float64 {
	a := bearing2 - ReverseDirection(bearing1)
	if a < 0.0 {
		a = a + math.Tau
	}
	return math.Pi - a
}
