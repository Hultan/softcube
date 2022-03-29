package surface

import (
	"image/color"

	"github.com/hultan/softcube/internal/vector"
)

type Surface2 struct {
	V1, V2, V3, V4 vector.Vector2
	C              color.Color
}

func (s Surface2) ToScreenCoords(width, height float64) Surface2 {
	return Surface2{
		V1: s.V1.MulScalars(width, height).AddScalars(width/2, height/2),
		V2: s.V2.MulScalars(width, height).AddScalars(width/2, height/2),
		V3: s.V3.MulScalars(width, height).AddScalars(width/2, height/2),
		V4: s.V4.MulScalars(width, height).AddScalars(width/2, height/2),
		C:  s.C,
	}
}
