package go3d

import (
	"math"
)

type Vector3d struct {
	X, Y, Z float64
}

func (v Vector3d) RotateX(theta float64) Vector3d {
	v2 := Vector3d{
		X: v.X,
		Y: v.Y*math.Cos(theta) - v.Z*math.Sin(theta),
		Z: v.Y*math.Sin(theta) + v.Z*math.Cos(theta),
	}
	return v2
}

func (v Vector3d) RotateY(theta float64) Vector3d {
	v2 := Vector3d{
		X: v.X*math.Cos(theta) + v.Z*math.Sin(theta),
		Y: v.Y,
		Z: -v.X*math.Sin(theta) + v.Z*math.Cos(theta),
	}
	return v2
}

func (v Vector3d) RotateZ(theta float64) Vector3d {
	v2 := Vector3d{
		X: v.X*math.Cos(theta) - v.Y*math.Sin(theta),
		Y: v.X*math.Sin(theta) + v.Y*math.Cos(theta),
		Z: v.Z,
	}
	return v2
}
