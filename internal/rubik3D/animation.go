package rubik3D

type axis int

const (
	axisX axis = iota
	axisY
	axisZ
)

const animationSteps = 10

type animation struct {
	isAnimation    bool
	afterAnimation func()
	step           int
	angle          float64
	endAngle       float64
	cubits         []int
	axis           axis
}

func (c *Cube) createAnimation(after func(), endAngle float64, cubits []int, a axis) *animation {
	return &animation{
		afterAnimation: after,
		isAnimation:    true,
		endAngle:       endAngle,
		cubits:         cubits,
		axis:           a,
	}
}

func (c *Cube) createNonAnimation(after func()) *animation {
	return &animation{
		afterAnimation: after,
		isAnimation:    false,
	}
}
