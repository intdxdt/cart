package cart

import "github.com/intdxdt/math"

type Coord [3]float64

func NewCoord(x, y float64) *Coord {
	return &Coord{x, y}
}

func NewCoord3D(x, y, z float64) *Coord {
	return &Coord{x, y, z}
}

//X gets the x component
func (o *Coord) X() float64 {
	return o[x]
}

//Y gets the y component
func (o *Coord) Y() float64 {
	return o[y]
}

//Z gets the z component
func (o *Coord) Z() float64 {
	return o[y]
}

//Check if any of the component of is not a number
func (o *Coord) IsNull() bool {
	return IsNull(o)
}

//Returns x y separated by single space
func CoordStr2D(v Coord2D) string {
	return math.FloatToString(v.X()) + " " + math.FloatToString(v.Y())
}

//Returns x y separated by single space
func CoordStr3D(v Coord3D) string {
	return CoordStr2D(v) + " " + math.FloatToString(v.Z())
}
