package cart

import (
	"github.com/intdxdt/math"
	"github.com/intdxdt/robust"
)

const (
	x = iota
	y
	z
)

type Coord2D interface {
	X() float64
	Y() float64
	IsNull() bool
}

type Coord3D interface {
	Coord2D
	Z() float64
}

//Component vector
func Component(m, d float64) (float64, float64) {
	return m * math.Cos(d), m * math.Sin(d)
}

//Equals evaluates whether two points are the same
func Equals(v, o Coord2D) bool {
	return math.FloatEqual(v.X(), o.X()) && math.FloatEqual(v.Y(), o.Y())
}

//Computes the addition of x and y components
func Add(v, o Coord2D) (float64, float64) {
	return v.X() + o.X(), v.Y() + o.Y()
}

//Computes the difference between x , y components
func Sub(v, o Coord2D) (float64, float64) {
	return v.X() - o.X(), v.Y() - o.Y()
}

//KProduct scales x and y components by constant  k
func KProduct(v Coord2D, k float64) (float64, float64) {
	return k * v.X(), k * v.Y()
}

//Negates components x and y
func Neg(v Coord2D) (float64, float64) {
	return KProduct(v, -1.0)
}

//Dot Product of two points as vectors
func DotProduct(v, o Coord2D) float64 {
	return DotProductXY(v.X(), v.Y(), o.X(), o.Y())
}

//Dot Product of two points as vectors
func DotProductXY(vx, vy, ox, oy float64) float64 {
	return (vx * ox) + (vy * oy)
}

//Unit vector of point
func Unit(v Coord2D) (float64, float64) {
	return UnitXY(v.X(), v.Y())
}

//Unit vector of point
func UnitXY(x, y float64) (float64, float64) {
	m := MagnitudeXY(x, y)
	if math.FloatEqual(m, 0.0) {
		m = math.EPSILON
	}
	return x / m, y / m
}

//Projects  u on to v
func Project(u, onv Coord2D) float64 {
	return ProjectXY(u.X(), u.Y(), onv.X(), onv.Y())
}

//Projects  u on to v, using x and y compoents of u and v
func ProjectXY(ux, uy, onvX, onvY float64) float64 {
	cx, cy := UnitXY(onvX, onvY)
	return DotProductXY(ux, uy, cx, cy)
}

//2D cross product of AB and AC vectors given A, B, and C as points,
//i.e. z-component of their 3D cross product.
//Returns a positive value, if ABC makes a counter-clockwise turn,
//negative for clockwise turn, and zero if the points are collinear.
func Orientation2D(a, b, c Coord2D) float64 {
	return robust.Orientation2D(
		[]float64{a.X(), a.Y()},
		[]float64{b.X(), b.Y()},
		[]float64{c.X(), c.Y()},
	)
}

//2D cross product of AB and AC vectors,
//i.e. z-component of their 3D cross product.
//negative cw and positive if ccw
func CrossProduct(ab, ac Coord2D) float64 {
	return (ab.X() * ac.Y()) - (ab.Y() * ac.X())
}

//Computes the square vector magnitude of pt as vector: x , y as components
//This has a potential overflow problem based on coordinates of pt x^2 + y^2
func SquareMagnitude(v Coord2D, other ...Coord2D) float64 {
	var dx, dy float64
	if len(other) == 0 {
		dx, dy = v.X(), v.Y()
	} else {
		o := other[0]
		dx, dy = o.X()-v.X(), o.Y()-v.Y()
	}
	return (dx * dx) + (dy * dy)
}

//Computes vector magnitude of pt as vector: x , y as components
func Magnitude(v Coord2D, other ...Coord2D) float64 {
	var dx, dy float64
	if len(other) == 0 {
		dx, dy = v.X(), v.Y()
	} else {
		o := other[0]
		dx, dy = o.X()-v.X(), o.Y()-v.Y()
	}
	return MagnitudeXY(dx, dy)
}

//Computes vector magnitude given x an dy component
func MagnitudeXY(dx, dy float64) float64 {
	return math.Hypot(dx, dy)
}

//Checks if catesian coordinate is null ( has NaN )
func IsNull(v Coord2D) bool {
	return math.IsNaN(v.X()) || math.IsNaN(v.Y())
}

//Checks if x and y components are zero
func IsZero(v Coord2D) bool {
	return math.FloatEqual(v.X(), 0.0) && math.FloatEqual(v.Y(), 0.0)
}

//Checks equality in 2d
func Equals2D(a, b Coord2D) bool {
	return math.FloatEqual(a.X(), b.X()) && math.FloatEqual(b.Y(), b.Y())
}

//Checks equality in 3d
func Equals3D(a, b Coord3D) bool {
	return Equals2D(a, b)  && math.FloatEqual(a.Z(), b.Z())
}
