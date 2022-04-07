package surface

import (
	"image/color"

	"github.com/hultan/softcube/internal/vector"
)

type Surface3 struct {
	V1, V2, V3, V4 vector.Vector3
	Col            color.Color
}

func (s *Surface3) Z() float64 {
	return (s.V1.Z + s.V2.Z + s.V3.Z + s.V4.Z) / 4
}

func (s *Surface3) Rotate(x, y, z float64) Surface3 {
	return Surface3{
		V1:  s.V1.RotateX(x).RotateY(y).RotateZ(z),
		V2:  s.V2.RotateX(x).RotateY(y).RotateZ(z),
		V3:  s.V3.RotateX(x).RotateY(y).RotateZ(z),
		V4:  s.V4.RotateX(x).RotateY(y).RotateZ(z),
		Col: s.Col,
	}
}

func (s *Surface3) To2DCoords(distance, cubeDistance float64) Surface2 {
	return Surface2{
		V1: to2dCoords(s.V1, distance, cubeDistance),
		V2: to2dCoords(s.V2, distance, cubeDistance),
		V3: to2dCoords(s.V3, distance, cubeDistance),
		V4: to2dCoords(s.V4, distance, cubeDistance),
		C1: s.Col,
	}
}

func to2dCoords(v vector.Vector3, distance, cubeDistance float64) vector.Vector2 {
	return vector.Vector2{
		X: v.X * distance / (v.Z + cubeDistance),
		Y: v.Y * distance / (v.Z + cubeDistance),
	}
}
