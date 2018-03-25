package cart

import (
	"github.com/intdxdt/math"
)

//Deflect_vector computes vector deflection given deflection angle and
// side of vector to deflect from (from_end)
func Deflect(v Coord2D, mag, deflAngle float64, fromEnd bool) (float64, float64) {
	angl := math.Pi - deflAngle
	return Extend(v, mag, angl, fromEnd)
}


//Extvect extends vector from the from end or from begin of vector
func Extend(v Coord2D, magnitude, angle float64, fromEnd bool) (float64, float64) {
	//from a of v back direction initiates as fwd v direction anticlockwise
	//bb - back bearing
	//fb - forward bearing
	bb := Direction(v)
	if fromEnd {
		bb += math.Pi
	}
	fb := bb + angle
	if fb > math.Tau {
		fb -= math.Tau
	}
	return Component(magnitude, fb)
}
