package vector

import (
	"fmt"
	"math"
)

type Vector3 struct {
	X, Y, Z float64
}

func (v Vector3) AddVector(v2 Vector3) Vector3 {
	return Vector3{
		X: v.X + v2.X,
		Y: v.Y + v2.Y,
		Z: v.Z + v2.Z,
	}
}

func (v Vector3) AddScalar(f float64) Vector3 {
	return Vector3{
		X: v.X + f,
		Y: v.Y + f,
		Z: v.Z + f,
	}
}

func (v Vector3) AddScalars(x, y, z float64) Vector3 {
	return Vector3{
		X: v.X + x,
		Y: v.Y + y,
		Z: v.Z + z,
	}
}

func (v Vector3) SubVector(v2 Vector3) Vector3 {
	return Vector3{
		X: v.X - v2.X,
		Y: v.Y - v2.Y,
		Z: v.Z - v2.Z,
	}
}

func (v Vector3) SubScalar(f float64) Vector3 {
	return Vector3{
		X: v.X - f,
		Y: v.Y - f,
		Z: v.Z - f,
	}
}

func (v Vector3) SubScalars(x, y, z float64) Vector3 {
	return Vector3{
		X: v.X - x,
		Y: v.Y - y,
		Z: v.Z - z,
	}
}

func (v Vector3) MulVector(v2 Vector3) Vector3 {
	return Vector3{
		X: v.X * v2.X,
		Y: v.Y * v2.Y,
		Z: v.Z * v2.Z,
	}
}

func (v Vector3) MulScalar(f float64) Vector3 {
	return Vector3{
		X: v.X * f,
		Y: v.Y * f,
		Z: v.Z * f,
	}
}

func (v Vector3) MulScalars(x, y, z float64) Vector3 {
	return Vector3{
		X: v.X * x,
		Y: v.Y * y,
		Z: v.Z * z,
	}
}

func (v Vector3) DivVector(v2 Vector3) Vector3 {
	return Vector3{
		X: v.X / v2.X,
		Y: v.Y / v2.Y,
		Z: v.Z / v2.Z,
	}
}

func (v Vector3) DivScalar(f float64) Vector3 {
	return Vector3{
		X: v.X / f,
		Y: v.Y / f,
		Z: v.Z / f,
	}
}

func (v Vector3) DivScalars(x, y, z float64) Vector3 {
	return Vector3{
		X: v.X / x,
		Y: v.Y / y,
		Z: v.Z / z,
	}
}

func (v Vector3) RotateX(theta float64) Vector3 {
	return Vector3{
		X: v.X,
		Y: v.Y*math.Cos(theta) - v.Z*math.Sin(theta),
		Z: v.Y*math.Sin(theta) + v.Z*math.Cos(theta),
	}
}

func (v Vector3) RotateY(theta float64) Vector3 {
	return Vector3{
		X: v.X*math.Cos(theta) + v.Z*math.Sin(theta),
		Y: v.Y,
		Z: -v.X*math.Sin(theta) + v.Z*math.Cos(theta),
	}
}

func (v Vector3) RotateZ(theta float64) Vector3 {
	return Vector3{
		X: v.X*math.Cos(theta) - v.Y*math.Sin(theta),
		Y: v.X*math.Sin(theta) + v.Y*math.Cos(theta),
		Z: v.Z,
	}
}

func (v Vector3) Dist() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
}

func (v Vector3) DistTo(v2 Vector3) float64 {
	dx := v.X - v2.X
	dy := v.Y - v2.Y
	dz := v.Z - v2.Z
	return math.Sqrt(dx*dx + dy*dy*dz*dz)
}

func (v Vector3) Dot(v2 Vector3) float64 {
	return v.X*v2.X + v.Y*v2.Y + v.Z*v2.Z
}

func (v Vector3) Cross(v2 Vector3) Vector3 {
	return Vector3{
		v.Y*v2.Z - v.Z*v2.Y,
		v.Z*v2.X - v.X*v2.Z,
		v.X*v2.Y - v.Y*v2.X,
	}
}

func (v Vector3) Normalize() Vector3 {
	m := v.Dist()

	if m > 0 {
		return v.DivScalar(m)
	} else {
		return v
	}
}

func (v Vector3) String() string {
	return fmt.Sprintf("(%f, %f, %f)", v.X, v.Y, v.Z)
}

func (v Vector3) Equals(v2 Vector3) bool {
	return v.X == v2.X && v.Y == v2.Y && v.Z == v2.Z
}
