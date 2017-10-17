package cart

import "github.com/intdxdt/math"

type Coord [2]float64

func NewCoord(x, y float64) *Coord {
	return &Coord{x, y}
}

//X gets the x component
func (o *Coord) X() float64 {
	return o[x]
}

//Y gets the y component
func (o *Coord) Y() float64 {
	return o[y]
}

//Check if any of the component of is not a number
func (o *Coord) IsNull() bool {
	return IsNull(o)
}

//Returns x y separated by single space
func CoordString(v Pt2D) string {
	return math.FloatToString(v.X()) + " " + math.FloatToString(v.Y())
}
