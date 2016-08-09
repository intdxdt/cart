package cart2d

import (
    "math"
    ."simplex/util/math"
)

//Dir computes direction in radians - counter clockwise from x-axis.
func Direction(v Cart2D) float64 {
	return DirectionXY(v.X(), v.Y())
}

//Dir computes direction in radians - counter clockwise from x-axis.
func DirectionXY(x , y float64) float64 {
	d := math.Atan2(y, x)
	if d < 0 {
		d += Tau
	}
	return d
}

//Revdir computes the reversed direction from a foward direction
func ReverseDirection(d float64) float64 {
	if d < Pi {
		return d + Pi
	}
	return d - Pi
}


func DeflectionAngle(bearing1, bearing2 float64) float64 {
	a := bearing2 - ReverseDirection(bearing1)
	if a < 0.0 {
		a = a + Tau
	}
	return Pi - a
}