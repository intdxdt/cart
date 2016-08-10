package cart2d

//Extvect extends vector from the from end or from begin of vector
func   Extvect(v Cart2D, magnitude, angle float64, from_end bool) (float64, float64) {
	//from a of v back direction initiates as fwd v direction anticlockwise
	//bβ - back bearing
	//fβ - forward bearing
	bβ := Direction(v)
	a := v.a
	if from_end {
		bβ +=  Pi
		a = v.b
	}
	fβ := bβ + angle
	if fβ > Tau {
		fβ -= Tau
	}

	opts := &Options{
		A : a,
		M : &magnitude,
		D : &fβ,
	}
	return NewVect(opts)
}


