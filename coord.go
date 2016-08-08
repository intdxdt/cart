package cart2d

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