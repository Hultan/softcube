package cube_ui

import (
	"math"
)

func (v vector3d) rotateX(theta float64) vector3d {
	v2 := vector3d{
		x: v.x,
		y: v.y*math.Cos(theta) - v.z*math.Sin(theta),
		z: v.y*math.Sin(theta) + v.z*math.Cos(theta),
	}
	return v2
}

func (v vector3d) rotateY(theta float64) vector3d {
	v2 := vector3d{
		x: v.x*math.Cos(theta) + v.z*math.Sin(theta),
		y: v.y,
		z: -v.x*math.Sin(theta) + v.z*math.Cos(theta),
	}
	return v2
}

func (v vector3d) rotateZ(theta float64) vector3d {
	v2 := vector3d{
		x: v.x*math.Cos(theta) - v.y*math.Sin(theta),
		y: v.x*math.Sin(theta) + v.y*math.Cos(theta),
		z: v.z,
	}
	return v2
}
