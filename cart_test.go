package cart

import (
	"testing"
	"github.com/intdxdt/math"
	"github.com/franela/goblin"
)

func TestCart(t *testing.T) {
	g := goblin.Goblin(t)
	p0 := NewCoord(0.0, 0.0)
	pn := NewCoord(0.0, math.NaN())
	p1 := NewCoord(4, 5)
	p2 := NewCoord(4.0, 5.0)
	p3 := &Coord{4.0, math.NaN()}

	g.Describe("geom.point", func() {
		g.It("x, y access & null", func() {
			g.Assert(IsZero(p0)).IsTrue()
			g.Assert(IsNull(pn)).IsTrue()

			g.Assert(IsZero(p1)).IsFalse()
			g.Assert(Equals(p1, p2)).IsTrue()

			g.Assert(p1.X()).Equal(4.0)
			g.Assert(p1.Y()).Equal(5.0)
			g.Assert(p0.IsNull()).IsFalse()
			g.Assert(IsNull(p3)).IsTrue()
		})

	})

	g.Describe("Point distance and to polygon ", func() {
		g.It("sqrt(3**2,4**2) ", func() {
			pt := &Coord{3., 0.}
			g.Assert(Magnitude(pt, &Coord{0., 4.})).Equal(5.0)
			g.Assert(MagnitudeSquare(pt, &Coord{0., 4.})).Equal(25.0)
		})
		g.It("sqrt(2)", func() {
			pt := &Coord{3, 4}
			g.Assert(Magnitude(pt, &Coord{4, 5})).Equal(math.Sqrt2)
			g.Assert(MagnitudeSquare(pt, &Coord{4, 5})).Equal(2.0)
		})
	})

	g.Describe("operators", func() {
		g.It("component ", func() {
			cx, cy := Component(5, math.Deg2rad(53.13010235415598))
			g.Assert(math.FloatEqual(cx, 3.0)).IsTrue()
			g.Assert(math.FloatEqual(cy, 4.0)).IsTrue()
		})
		g.It("add ", func() {
			a, b := &Coord{3., 0.}, &Coord{0., 4.}
			cx, cy := Add(a, b)
			g.Assert(&Coord{cx, cy}).Equal(&Coord{3., 4.})
		})

		g.It("sub & neg ", func() {
			a, b := &Coord{3., 4.}, &Coord{4, 5}
			cx, cy := Neg(b)
			g.Assert(&Coord{cx, cy}).Equal(&Coord{-4, -5})
			cx, cy = Sub(a, b)
			g.Assert(&Coord{cx, cy}).Equal(&Coord{-1.0, -1.0})
		})
	})
}

//Test Neg
func TestNeg(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("Negate Vector", func() {
		g.It("should test vector negation", func() {
			a := []float64{10, 150, 6.5}
			e := []float64{280, 280, 12.8}
			A := NewCoord(a[x], a[y])
			B := NewCoord(e[x], e[y])

			vx, vy := Sub(B, A)
			pv := NewCoord(vx, vy)
			vx, vy = Neg(pv)
			nv := NewCoord(vx, vy)
			negA := NewCoord(0, 0)
			for i, v := range A {
				negA[i] = -v
			}
			vx, vy = KProduct(pv, -1)
			g.Assert(nv).Eql(NewCoord(vx, vy))
			vx, vy = Neg(A)
			g.Assert(NewCoord(vx, vy)).Eql(negA)

		})
	})

}

func TestMagDist(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("Point - Vector Magnitude", func() {
		g.It("should test vector magnitude and distance", func() {
			a := &Coord{0, 0 }
			b := &Coord{3, 4 }
			z := NewCoord(0, 0)
			g.Assert(Magnitude(NewCoord(1, 1), z)).Equal(math.Sqrt2)
			g.Assert(math.Round(Magnitude(NewCoord(-3, 2), z), 8)).Equal(
				math.Round(3.605551275463989, 8),
			)
			g.Assert(Magnitude(NewCoord(3, 4))).Equal(5.0)
			g.Assert(Magnitude(NewCoord(3, 4), z)).Equal(5.0)
			g.Assert(Magnitude(a, b)).Equal(5.0)
			g.Assert(MagnitudeSquare(NewCoord(3, 4), z)).Equal(25.0)
			g.Assert(MagnitudeSquare(NewCoord(3, 4))).Equal(25.0)
			g.Assert(MagnitudeSquare(a, b)).Equal(25.0)
			g.Assert(Magnitude(NewCoord(4.587, 0.), z)).Equal(4.587)
		})
	})

}

func TestDotProduct(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("Point - Vector Dot Product", func() {
		g.It("should test dot product", func() {
			dotProd := DotProduct(NewCoord(1.2, -4.2), NewCoord(1.2, -4.2))
			g.Assert(19.08).Equal(math.Round(dotProd, 8))
		})
	})

}

func TestSideOf(t *testing.T) {
	g := goblin.Goblin(t)

	a := NewCoord(237, 289)
	b := NewCoord(404.25, 357.25)
	c := NewCoord(460, 380)

	d := NewCoord(297.13043478260863, 339.30434782608694)
	e := NewCoord(445.8260869565217, 350.17391304347825)

	cx, cy := Sub(b, a)
	ab := NewCoord(cx, cy)

	cx, cy = Sub(c, a)
	ac := NewCoord(cx, cy)

	cx, cy = Sub(d, a)
	ad := NewCoord(cx, cy)

	cx, cy = Sub(e, a)
	ae := NewCoord(cx, cy)

	g.Describe("Orientation and cross product", func() {
		g.It("orientation", func() {
			g.Assert(Orientation2D(a, b, c) == 0).IsTrue()
			g.Assert(Orientation2D(a, c, d) < 0).IsTrue()
			g.Assert(Orientation2D(a, c, e) > 0).IsTrue()

		})
		g.It("cross product", func() {
			g.Assert(CrossProduct(ab, ac) == 0).IsTrue()
			g.Assert(CrossProduct(ac, ad) > 0).IsTrue()
			g.Assert(CrossProduct(ac, ae) < 0).IsTrue()
		})
	})

}

func TestCCW(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("Vector Sidedness", func() {
		g.It("should test side of point to vector", func() {
			k := &Coord{-0.887, -1.6128}
			u := &Coord{4.55309, 1.42996}

			testpoints := []*Coord{
				{2, 2}, {0, 2}, {0, -2}, {2, -2}, {0, 0}, {2, 0}, u, k,
			}

			left, right, on := func(x float64) bool {
				return x < 0
			}, func(x float64) bool {
				return x > 0
			}, func(x float64) bool {
				return math.FloatEqual(x, 0)
			}

			sides := make([]float64, len(testpoints))
			for i, pnt := range testpoints {
				sides[i] = Orientation2D(k, u, pnt)
			}
			g.Assert(Orientation2D(k, u, &Coord{2, 2}) < 0).IsTrue()

			sideOut := []func(x float64) bool{
				left, left, right, right, left,
				right, on, on,
			}

			for i := range sideOut {
				g.Assert(sideOut[i](sides[i])).IsTrue()
			}
		})
	})

}

func TestProj(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("Vector - unit & Project", func() {
		var A = &Coord{0.88682, -1.06102}
		var B = &Coord{3.5, 1.0}
		g.It("should test projection", func() {
			g.Assert(math.Round(Project(A, B), 5)).Equal(0.56121)
		})
		g.It("should test Unit", func() {
			Z := &Coord{0., 0.}
			cx, cy := Unit(Z)
			g.Assert(math.FloatEqual(cx, 0)).IsTrue()
			g.Assert(math.FloatEqual(cy, 0)).IsTrue()
		})
	})
}

func TestDirection(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("Vector Direction", func() {
		g.It("should test vector direction", func() {
			A := &Coord{0, 0}
			B := &Coord{-1, 0}
			cx, cy := Sub(B, A)
			v := NewCoord(cx, cy)
			g.Assert(Direction(1, 1)).Equal(0.7853981633974483)
			g.Assert(Direction(-1, 0)).Equal(math.Pi)
			g.Assert(Direction(v[x], v[y])).Equal(math.Pi)
			g.Assert(Direction(1, math.Sqrt(3))).Equal(math.Deg2rad(60))
			g.Assert(Direction(0, -1)).Equal(math.Deg2rad(270))
		})
	})

}

func TestReverseDirection(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("Vector RevDirection", func() {
		g.It("should test reverse vector direction", func() {
			A := &Coord{0, 0}
			B := &Coord{-1, 0}
			cx, cy := Sub(B, A)
			v := NewCoord(cx, cy)
			g.Assert(ReverseDirection(Direction(v[x], v[y]))).Equal(0.0)
			g.Assert(ReverseDirection(0.7853981633974483)).Equal(0.7853981633974483 + math.Pi)
			g.Assert(ReverseDirection(0.7853981633974483 + math.Pi)).Equal(0.7853981633974483)
		})
	})

}

func TestDeflection(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("Vector Deflection", func() {
		g.It("should test reverse vector direction", func() {
			ln0 := []*Coord{{0, 0}, {20, 30}}
			ln1 := []*Coord{{20, 30}, {40, 15}}
			cx, cy := Sub(ln0[1], ln0[0])
			v0 := &Coord{cx, cy }
			cx, cy = Sub(ln1[1], ln1[0])
			v1 := &Coord{cx, cy }

			g.Assert(math.Round(DeflectionAngle(
				Direction(v0[x],v0[y]),
				Direction(v1[x],v1[y]),
			), 10)).Equal(math.Round(math.Deg2rad(93.17983011986422), 10))
			g.Assert(math.Round(DeflectionAngle(
				Direction(v0[x],v0[y]),
				Direction(v0[x],v0[y]),
			), 10)).Equal(math.Deg2rad(0.0))

			ln1 = []*Coord{{20, 30}, {20, 60}}
			cx, cy = Sub(ln1[1], ln1[0])
			v1 = &Coord{cx, cy }
			g.Assert(math.Round(DeflectionAngle(
				Direction(v0[x], v0[y]),
				Direction(v1[x], v1[y]),
			), 10)).Equal(
				math.Round(math.Deg2rad(-33.690067525979806), 10),
			)
		})
	})

}

func TestDistanceToPoint(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("Vector - Dist2Vect", func() {
		g.It("should test distance vector", func() {
			a := &Coord{16.82295, 10.44635}
			b := &Coord{28.99656, 15.76452}
			onAb := &Coord{25.32, 14.16}

			tpoints := []*Coord{
				{30., 0.},
				{15.78786, 25.26468},
				{-2.61504, -3.09018},
				{28.85125, 27.81773},
				a,
				b,
				onAb,
			}

			tDists := []float64{14.85, 13.99, 23.69, 12.05, 0.00, 0.00, 0.00}
			dists := make([]float64, len(tpoints))

			for i, tp := range tpoints {
				dists[i] = DistanceToPoint(a, b, tp)
			}

			for i := range tpoints {
				g.Assert(math.Round(dists[i], 2)).Equal(math.Round(tDists[i], 2))
			}
		})
	})

}
