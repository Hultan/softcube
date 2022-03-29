package vector

import (
	"fmt"
	"math"
)

type Vector2 struct {
	X, Y float64
}

func (v Vector2) AddVector(v2 Vector2) Vector2 {
	return Vector2{
		X: v.X + v2.X,
		Y: v.Y + v2.Y,
	}
}

func (v Vector2) AddScalar(f float64) Vector2 {
	return Vector2{
		X: v.X + f,
		Y: v.Y + f,
	}
}

func (v Vector2) AddScalars(x, y float64) Vector2 {
	return Vector2{
		X: v.X + x,
		Y: v.Y + y,
	}
}

func (v Vector2) SubVector(v2 Vector2) Vector2 {
	return Vector2{
		X: v.X - v2.X,
		Y: v.Y - v2.Y,
	}
}

func (v Vector2) SubScalar(f float64) Vector2 {
	return Vector2{
		X: v.X - f,
		Y: v.Y - f,
	}
}

func (v Vector2) SubScalars(x, y float64) Vector2 {
	return Vector2{
		X: v.X - x,
		Y: v.Y - y,
	}
}

func (v Vector2) MulVector(v2 Vector2) Vector2 {
	return Vector2{
		X: v.X * v2.X,
		Y: v.Y * v2.Y,
	}
}

func (v Vector2) MulScalar(f float64) Vector2 {
	return Vector2{
		X: v.X * f,
		Y: v.Y * f,
	}
}

func (v Vector2) MulScalars(x, y float64) Vector2 {
	return Vector2{
		X: v.X * x,
		Y: v.Y * y,
	}
}

func (v Vector2) DivVector(v2 Vector2) Vector2 {
	return Vector2{
		X: v.X / v2.X,
		Y: v.Y / v2.Y,
	}
}

func (v Vector2) DivScalar(f float64) Vector2 {
	return Vector2{
		X: v.X / f,
		Y: v.Y / f,
	}
}

func (v Vector2) DivScalars(x, y float64) Vector2 {
	return Vector2{
		X: v.X / x,
		Y: v.Y / y,
	}
}

func (v Vector2) Dist() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v Vector2) DistTo(v2 Vector2) float64 {
	dx := v.X - v2.X
	dy := v.Y - v2.Y
	return math.Sqrt(dx*dx + dy*dy)
}

func (v Vector2) Dot(v2 Vector2) float64 {
	return v.X*v2.X + v.Y*v2.Y
}

func (v Vector2) Normalize() Vector2 {
	m := v.Dist()

	if m > 0 {
		return v.DivScalar(m)
	} else {
		return v
	}
}

func (v Vector2) String() string {
	return fmt.Sprintf("(%f, %f)", v.X, v.Y)
}

func (v Vector2) Equals(v2 Vector2) bool {
	return v.X == v2.X && v.Y == v2.Y
}
