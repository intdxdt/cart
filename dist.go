package cart

import (
	"github.com/intdxdt/math"
)

// Minimum distance to vector from a point
// if points outside the range of the vector the minimum distance
// is not perperndicular to the vector
// modified @Ref: http://www.mappinghacks.com/code/PolyLineReduction/
func DistanceToPoint(a, b, pnt Coord2D) float64 {

	var vx, vy = b.X()-a.X(), b.Y()-a.Y()
	var ux, uy = pnt.X()-a.X(), pnt.Y()-a.Y()
	var dist_uv = ProjectXY(ux, uy, vx, vy)
	var rstate = false
	var result = 0.0

	if dist_uv < 0 {
		// if negative
		result = MagnitudeXY(ux, uy)
		rstate = true
	} else {
		var d = ProjectXY(ux-vx, uy-vy, -vx, -vy)
		if d < 0.0 {
			result = MagnitudeXY(ux-vx, uy-vy)
			rstate = true
		}
	}

	if rstate == false {
		// avoid floating point imprecision
		var h = math.Round(math.Abs(MagnitudeXY(ux, uy)), math.PRECISION)
		var a = math.Round(math.Abs(dist_uv), math.PRECISION)

		if math.FloatEqual(h, 0.0) && math.FloatEqual(a, 0.0) {
			result = 0.0
		} else {
			var r = math.Round(a/h, math.PRECISION)
			// to avoid numeric overflow
			result = h * math.Sqrt(1-r*r)
		}
	}
	//opposite distance to hypotenus
	return result
}
