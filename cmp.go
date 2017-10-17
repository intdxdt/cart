package cart

import "github.com/intdxdt/math"

//2D compare x | y ordering
func Compare(self, other Coord2D) int {
	d  := self.X() - other.X()
	if math.FloatEqual(d, 0.0) {
		d = self.Y() - other.Y()
	}
	if math.FloatEqual(d, 0.0) {
		return 0
	} else if d < 0 {
		return -1
	}
	return 1
}
