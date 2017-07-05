package cart2d

import (
	"simplex/util/math"
)

const precision = 12

//Dist2Pt computes distance from a point to Vect
// Minimum distance to vector from a point
// compute the minimum distance between point and vector
// if points outside the range of the vector the minimum distance
// is not perperndicular to the vector
// modified @Ref: http://www.mappinghacks.com/code/PolyLineReduction/
func DistanceToPoint(a, b, pnt Pt2D) float64 {
	//vect = &Options{A: vect.a, B : pnt, }
	vx, vy := b.X()-a.X(), b.Y()-a.Y()
	ux, uy := pnt.X()-a.X(), pnt.Y()-a.Y()

	dist_uv := ProjectXY(ux, uy, vx, vy)

	rstate := false
	result := 0.0

	if dist_uv < 0 {
		// if negative
		result = MagnitudeXY(ux, uy)
		rstate = true
	} else {
		d := ProjectXY(ux-vx, uy-vy, -vx, -vy)
		if d < 0.0 {
			result = MagnitudeXY(ux-vx, uy-vy)
			rstate = true
		}
	}

	if rstate == false {
		// avoid floating point imprecision
		h := math.Round(math.Abs(MagnitudeXY(ux, uy)), precision)
		a := math.Round(math.Abs(dist_uv), precision)

		if math.FloatEqual(h, 0.0) && math.FloatEqual(a, 0.0) {
			result = 0.0
		} else {
			r := math.Round(a/h, precision)
			// to avoid numeric overflow
			result = h * math.Sqrt(1-r*r)
		}
	}
	//opposite distance to hypotenus
	return result
}
