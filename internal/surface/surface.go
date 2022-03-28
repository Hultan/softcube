package surface

import (
	"image/color"

	"github.com/hultan/softcube/internal/vector"
)

type Surface struct {
	V1, V2, V3, V4 vector.Vector3
	C              color.Color
}

func (s *Surface) Z() float64 {
	return (s.V1.Z + s.V2.Z + s.V3.Z + s.V4.Z) / 4
}
