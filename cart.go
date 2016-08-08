package cart2d

import (
    "math"
    . "simplex/util/math"
)

const (
    x = iota
    y
)

const ε = 1e-12

type Cart2D interface {
    X() float64
    Y() float64
}

//Component vector
func Component(m, d float64) (float64, float64) {
	return m * math.Cos(d), m * math.Sin(d)
}


//Equals evaluates whether two points are the same
func Equals(v, o Cart2D) bool {
    return FloatEqual(v.X(), o.X()) && FloatEqual(v.Y(), o.Y())
}


//Computes the addition of x and y components
func Add(v Cart2D, o Cart2D) (float64, float64) {
    return v.X() + o.X(), v.Y() + o.Y()
}
//Computes the difference between x , y components
func Sub(v Cart2D, o Cart2D) (float64, float64) {
    return v.X() - o.X(), v.Y() - o.Y()
}

//KProduct scales x and y components by constant  k
func KProduct(v Cart2D, k float64) (float64, float64) {
    return k * v.X(), k * v.Y()
}

//Negates components x and y
func Neg(v Cart2D) (float64, float64) {
    return KProduct(v, -1.0)
}

//Dot Product of two points as vectors
func DotProduct(v, o Cart2D) float64 {
    return (v.X() * o.X()) + (v.Y() * o.Y())
}


//Unit vector of point
func Unit(v Cart2D) (float64, float64) {
    m := Magnitude(v)
    if FloatEqual(m, 0.0) {
        m = ε
    }
    return v.X() / m, v.Y() / m
}


//Projects  u on to v
func Project(u, onv Cart2D) float64 {
    cx, cy := Unit(onv)
    return DotProduct(u, NewCoord(cx, cy))
}

//2D cross product of OA and OB vectors,
//i.e. z-component of their 3D cross product.
//Returns a positive value, if OAB makes a counter-clockwise turn,
//negative for clockwise turn, and zero if the points are collinear.
func CCW(o, a, b Cart2D) float64 {
    return (b.X() - a.X()) * (o.Y() - a.Y()) - (b.Y() - a.Y()) * (o.X() - a.X())
}

//Computes the square vector magnitude of pt as vector: x , y as components
//This has a potential overflow problem based on coordinates of pt x^2 + y^2
func SquareMagnitude(v Cart2D, other ...Cart2D) float64 {
    var dx, dy float64
    if len(other) == 0 {
        dx, dy = v.X(), v.Y()
    } else {
        o := other[0]
        dx, dy = o.X() - v.X(), o.Y() - v.Y()
    }
    return sqr(dx) + sqr(dy)
}


//Computes vector magnitude of pt as vector: x , y as components
func Magnitude(v Cart2D, other ...Cart2D) float64 {
    var m float64
    if len(other) == 0 {
        m = math.Hypot(v.X(), v.Y())
    } else {
        o := other[0]
        m = math.Hypot(o.X() - v.X(), o.Y() - v.Y())
    }
    return m
}


//Checks if catesian coordinate is null ( has NaN )
func IsNull(v Cart2D) bool {
    return math.IsNaN(v.X()) || math.IsNaN(v.Y())
}

//Checks if x and y components are zero
func IsZero(v Cart2D) bool {
    return FloatEqual(v.X(), 0.0) && FloatEqual(v.Y(), 0.0)
}



//computes square of a number
func sqr(x float64) float64 {
    return x * x
}




