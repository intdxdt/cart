package cart2d

import (
    . "simplex/util/math"
    . "github.com/franela/goblin"
    "testing"
)

const prec = 8

var A2 = &Coord{0.88682, -1.06102}
var B2 = &Coord{3.5, 1}
var C2 = &Coord{-3, 1}
var D2 = &Coord{-1.5, -3}


func TestExtVect(t *testing.T) {
    g := Goblin(t)
    g.Describe("Vector - Extend", func() {
        g.It("should test extending a vector", func() {

            va := A2
            vb := B2
            vc := C2
            vd := D2
            cx, cy := Sub(B2, D2)
            vdb := &Coord{cx, cy}
            cx, cy = Sub(C2, B2)
            vbc := &Coord{cx, cy}

            g.Assert(Round(Direction(va), prec)).Equal(
                Round(Deg2rad(309.889497029295), prec),
            )
            g.Assert(Round(Direction(vb), prec)).Equal(
                Round(Deg2rad(15.945395900922854), prec),
            )
            g.Assert(Round(Direction(vc), prec)).Equal(
                Round(Deg2rad(161.565051177078), prec),
            )
            g.Assert(Round(Direction(vd), prec)).Equal(
                Round(Deg2rad(243.43494882292202), prec),
            )
            g.Assert(Round(Magnitude(vdb), 4)).Equal(
                Round(6.4031242374328485, 4),
            )
            g.Assert(Round(Direction(vdb), prec)).Equal(
                Round(Deg2rad(38.65980825409009), prec),
            )
            deflangle := 157.2855876468
            cx, cy = Extend(vdb, 3.64005494464026, Deg2rad(180 + deflangle), true)
            vo := &Coord{cx, cy}

            g.Assert(Round(vo[0], prec)).Equal(
                Round(-vb[0], prec),
            )
            g.Assert(Round(vo[1], prec)).Equal(
                Round(-vb[1], prec),
            )



            // "vo by extending vdb by angle to origin"
            // "vo by extending vdb by angle to origin"
            deflangle_B := 141.34019174590992

            // extend to c from end
            cx, cy = Extend(vdb, 6.5, Deg2rad(180 + deflangle_B), true)
            vextc := &Coord{cx, cy}
            g.Assert(Round(vbc[0], prec)).Equal(
                Round(vextc[0], prec),
            )
            g.Assert(Round(vbc[1], prec)).Equal(
                Round(vextc[1], prec),
            )

            // "vextc with magnitudie extension from vdb C"
            g.Assert(Round(vextc[0], prec)).Equal(-Magnitude(vextc))
            // "vextc horizontal vector test:  extension from vdb C"
            g.Assert(Round(vextc[1], prec)).Equal(0.)

            vm  := &Coord{5, 0}
            cx, cy = Deflect(vm, 2, Deg2rad(90), true )
            //deflection is the right hand angle
            g.Assert(Round(cx, prec)).Equal(
                Round(0.0, prec),
            )
            g.Assert(Round(cy, prec)).Equal(
                Round(-2, prec),
            )
            cx, cy = Deflect(vm, 2, Deg2rad(90), false )
            g.Assert(Round(cx, prec)).Equal(
                Round(0.0, prec),
            )
            g.Assert(Round(cy, prec)).Equal(
                Round(2, prec),
            )

        })
    })

}