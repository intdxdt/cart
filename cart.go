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
    IsNull() bool
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
    return DotProductXY(v.X(), v.Y(), o.X(), o.Y())
}

//Dot Product of two points as vectors
func DotProductXY(vx, vy, ox, oy float64) float64 {
    return (vx * ox) + (vy * oy)
}


//Unit vector of point
func Unit(v Cart2D) (float64, float64) {
    return UnitXY(v.X(), v.Y())
}

//Unit vector of point
func UnitXY(x, y float64) (float64, float64) {
    m := MagnitudeXY(x, y)
    if FloatEqual(m, 0.0) {
        m = ε
    }
    return x / m, y / m
}


//Projects  u on to v
func Project(u, onv Cart2D) float64 {
    return ProjectXY(u.X(), u.Y(), onv.X(), onv.Y())
}

//Projects  u on to v, using x and y compoents of u and v
func ProjectXY(ux, uy , onv_x, onv_y float64) float64 {
    cx, cy := UnitXY(onv_x, onv_y)
    return DotProductXY(ux, uy, cx, cy)
}

//2D cross product of AB and AC vectors given A, B, and C as points,
//i.e. z-component of their 3D cross product.
//Returns a positive value, if ABC makes a counter-clockwise turn,
//negative for clockwise turn, and zero if the points are collinear.
func CCW(a, b, c Cart2D) float64 {
    return (b.X() - a.X()) * (c.Y() - a.Y()) - (b.Y() - a.Y()) * (c.X() - a.X())
}
//2D cross product of AB and AC vectors,
//i.e. z-component of their 3D cross product.
//Returns a positive value, if AB-->BC makes a counter-clockwise turn,
//negative for clockwise turn, and zero if the points are collinear.
func CCWVector(ab, ac Cart2D) float64 {
    return (ab.X() * ac.Y()) - (ab.Y() * ac.X())
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
    var dx, dy float64
    if len(other) == 0 {
        dx, dy = v.X(), v.Y()
    } else {
        o := other[0]
        dx, dy = o.X() - v.X(), o.Y() - v.Y()
    }
    return MagnitudeXY(dx, dy)
}

//Computes vector magnitude given x an dy component
func MagnitudeXY(dx, dy float64) float64 {
    return math.Hypot(dx, dy)
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




