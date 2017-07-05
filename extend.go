package cart2d

import (
	"simplex/util/math"
)

//Deflect_vector computes vector deflection given deflection angle and
// side of vector to deflect from (from_end)
func Deflect(v Pt2D, mag, deflAngle float64, fromEnd bool) (float64, float64) {
	angl := math.Pi - deflAngle
	return Extend(v, mag, angl, fromEnd)
}

//Extvect extends vector from the from end or from begin of vector
func Extend(v Pt2D, magnitude, angle float64, from_end bool) (float64, float64) {
	//from a of v back direction initiates as fwd v direction anticlockwise
	//bβ - back bearing
	//fβ - forward bearing
	bβ := Direction(v)
	if from_end {
		bβ += math.Pi
	}
	fβ := bβ + angle
	if fβ > math.Tau {
		fβ -= math.Tau
	}
	return Component(magnitude, fβ)
}
