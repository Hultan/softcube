package object

import (
	"github.com/hultan/softcube/internal/surface"
)

type Object struct {
	Surfaces []surface.Surface3
}

func (o *Object) Rotate(x, y, z float64) []surface.Surface3 {
	var rotated []surface.Surface3
	for _, s := range o.Surfaces {
		// Rotate
		r1 := s.V1.RotateX(x).RotateY(y).RotateZ(z)
		r2 := s.V2.RotateX(x).RotateY(y).RotateZ(z)
		r3 := s.V3.RotateX(x).RotateY(y).RotateZ(z)
		r4 := s.V4.RotateX(x).RotateY(y).RotateZ(z)

		rotated = append(rotated, surface.Surface3{
			V1: r1,
			V2: r2,
			V3: r3,
			V4: r4,
			C1: s.C1,
		})
	}
	return rotated
}
