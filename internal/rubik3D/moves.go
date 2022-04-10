package rubik3D

import (
	"math"
)

func (c *Cube) X() {
	f := func() {
		c.internalCube = c.internalCube.X()
		c.updateColors()
	}
	a := c.createAnimation(f, math.Pi/2, c.getAllCubits(), axisX)
	c.animatingQueue = append(c.animatingQueue, a)
}

func (c *Cube) Xc() {
	f := func() {
		c.internalCube = c.internalCube.Xc()
		c.updateColors()
	}
	a := c.createAnimation(f, -math.Pi/2, c.getAllCubits(), axisX)
	c.animatingQueue = append(c.animatingQueue, a)
}

func (c *Cube) Y() {
	f := func() {
		c.internalCube = c.internalCube.Y()
		c.updateColors()
	}
	a := c.createAnimation(f, math.Pi/2, c.getAllCubits(), axisY)
	c.animatingQueue = append(c.animatingQueue, a)
}

func (c *Cube) Yc() {
	f := func() {
		c.internalCube = c.internalCube.Yc()
		c.updateColors()
	}
	a := c.createAnimation(f, -math.Pi/2, c.getAllCubits(), axisY)
	c.animatingQueue = append(c.animatingQueue, a)
}

func (c *Cube) Z() {
	f := func() {
		c.internalCube = c.internalCube.Z()
		c.updateColors()
	}
	a := c.createAnimation(f, math.Pi/2, c.getAllCubits(), axisZ)
	c.animatingQueue = append(c.animatingQueue, a)
}

func (c *Cube) Zc() {
	f := func() {
		c.internalCube = c.internalCube.Zc()
		c.updateColors()
	}
	a := c.createAnimation(f, -math.Pi/2, c.getAllCubits(), axisZ)
	c.animatingQueue = append(c.animatingQueue, a)
}

func (c *Cube) U() {
	f := func() {
		c.internalCube = c.internalCube.U()
		c.updateColors()
	}
	a := c.createAnimation(f, math.Pi/2, c.getUMoveCubits(), axisY)
	c.animatingQueue = append(c.animatingQueue, a)
}

func (c *Cube) Uc() {
	f := func() {
		c.internalCube = c.internalCube.Uc()
		c.updateColors()
	}
	a := c.createAnimation(f, -math.Pi/2, c.getUMoveCubits(), axisY)
	c.animatingQueue = append(c.animatingQueue, a)
}

func (c *Cube) D() {
	f := func() {
		c.internalCube = c.internalCube.D()
		c.updateColors()
	}
	a := c.createAnimation(f, -math.Pi/2, c.getDMoveCubits(), axisY)
	c.animatingQueue = append(c.animatingQueue, a)
}

func (c *Cube) Dc() {
	f := func() {
		c.internalCube = c.internalCube.Dc()
		c.updateColors()
	}
	a := c.createAnimation(f, math.Pi/2, c.getDMoveCubits(), axisY)
	c.animatingQueue = append(c.animatingQueue, a)
}

func (c *Cube) L() {
	f := func() {
		c.internalCube = c.internalCube.L()
		c.updateColors()
	}
	a := c.createAnimation(f, math.Pi/2, c.getLMoveCubits(), axisX)
	c.animatingQueue = append(c.animatingQueue, a)
}

func (c *Cube) Lc() {
	f := func() {
		c.internalCube = c.internalCube.Lc()
		c.updateColors()
	}
	a := c.createAnimation(f, -math.Pi/2, c.getLMoveCubits(), axisX)
	c.animatingQueue = append(c.animatingQueue, a)
}

func (c *Cube) R() {
	f := func() {
		c.internalCube = c.internalCube.R()
		c.updateColors()
	}
	a := c.createAnimation(f, -math.Pi/2, c.getRMoveCubits(), axisX)
	c.animatingQueue = append(c.animatingQueue, a)
}

func (c *Cube) Rc() {
	f := func() {
		c.internalCube = c.internalCube.Rc()
		c.updateColors()
	}
	a := c.createAnimation(f, math.Pi/2, c.getRMoveCubits(), axisX)
	c.animatingQueue = append(c.animatingQueue, a)
}

func (c *Cube) F() {
	f := func() {
		c.internalCube = c.internalCube.F()
		c.updateColors()
	}
	a := c.createAnimation(f, math.Pi/2, c.getFMoveCubits(), axisZ)
	c.animatingQueue = append(c.animatingQueue, a)
}

func (c *Cube) Fc() {
	f := func() {
		c.internalCube = c.internalCube.Fc()
		c.updateColors()
	}
	a := c.createAnimation(f, -math.Pi/2, c.getFMoveCubits(), axisZ)
	c.animatingQueue = append(c.animatingQueue, a)
}

func (c *Cube) B() {
	f := func() {
		c.internalCube = c.internalCube.B()
		c.updateColors()
	}
	a := c.createAnimation(f, -math.Pi/2, c.getBMoveCubits(), axisZ)
	c.animatingQueue = append(c.animatingQueue, a)
}

func (c *Cube) Bc() {
	f := func() {
		c.internalCube = c.internalCube.Bc()
		c.updateColors()
	}
	a := c.createAnimation(f, math.Pi/2, c.getBMoveCubits(), axisZ)
	c.animatingQueue = append(c.animatingQueue, a)
}

func (c *Cube) M() {
	f := func() {
		c.internalCube = c.internalCube.M()
		c.updateColors()
	}
	a := c.createAnimation(f, math.Pi/2, c.getMMoveCubits(), axisX)
	c.animatingQueue = append(c.animatingQueue, a)
}

func (c *Cube) Mc() {
	f := func() {
		c.internalCube = c.internalCube.Mc()
		c.updateColors()
	}
	a := c.createAnimation(f, -math.Pi/2, c.getMMoveCubits(), axisX)
	c.animatingQueue = append(c.animatingQueue, a)
}

func (c *Cube) S() {
	f := func() {
		c.internalCube = c.internalCube.S()
		c.updateColors()
	}
	a := c.createAnimation(f, math.Pi/2, c.getSMoveCubits(), axisZ)
	c.animatingQueue = append(c.animatingQueue, a)
}

func (c *Cube) Sc() {
	f := func() {
		c.internalCube = c.internalCube.Sc()
		c.updateColors()
	}
	a := c.createAnimation(f, -math.Pi/2, c.getSMoveCubits(), axisZ)
	c.animatingQueue = append(c.animatingQueue, a)
}

func (c *Cube) E() {
	f := func() {
		c.internalCube = c.internalCube.E()
		c.updateColors()
	}
	a := c.createAnimation(f, -math.Pi/2, c.getEMoveCubits(), axisY)
	c.animatingQueue = append(c.animatingQueue, a)
}

func (c *Cube) Ec() {
	f := func() {
		c.internalCube = c.internalCube.Ec()
		c.updateColors()
	}
	a := c.createAnimation(f, math.Pi/2, c.getEMoveCubits(), axisY)
	c.animatingQueue = append(c.animatingQueue, a)
}

func (c *Cube) Ud() {
	f := func() {
		c.internalCube = c.internalCube.Ud()
		c.updateColors()
	}
	a := c.createAnimation(f, math.Pi/2, c.getUdMoveCubits(), axisY)
	c.animatingQueue = append(c.animatingQueue, a)
}

func (c *Cube) Udc() {
	f := func() {
		c.internalCube = c.internalCube.Udc()
		c.updateColors()
	}
	a := c.createAnimation(f, -math.Pi/2, c.getUdMoveCubits(), axisY)
	c.animatingQueue = append(c.animatingQueue, a)
}

func (c *Cube) Dd() {
	f := func() {
		c.internalCube = c.internalCube.Dd()
		c.updateColors()
	}
	a := c.createAnimation(f, -math.Pi/2, c.getDdMoveCubits(), axisY)
	c.animatingQueue = append(c.animatingQueue, a)
}

func (c *Cube) Ddc() {
	f := func() {
		c.internalCube = c.internalCube.Ddc()
		c.updateColors()
	}
	a := c.createAnimation(f, math.Pi/2, c.getDdMoveCubits(), axisY)
	c.animatingQueue = append(c.animatingQueue, a)
}

func (c *Cube) Fd() {
	f := func() {
		c.internalCube = c.internalCube.Fd()
		c.updateColors()
	}
	a := c.createAnimation(f, math.Pi/2, c.getFdMoveCubits(), axisY)
	c.animatingQueue = append(c.animatingQueue, a)
}

func (c *Cube) Fdc() {
	f := func() {
		c.internalCube = c.internalCube.Fdc()
		c.updateColors()
	}
	a := c.createAnimation(f, -math.Pi/2, c.getFdMoveCubits(), axisY)
	c.animatingQueue = append(c.animatingQueue, a)
}

func (c *Cube) Bd() {
	f := func() {
		c.internalCube = c.internalCube.Bd()
		c.updateColors()
	}
	a := c.createAnimation(f, -math.Pi/2, c.getBdMoveCubits(), axisY)
	c.animatingQueue = append(c.animatingQueue, a)
}

func (c *Cube) Bdc() {
	f := func() {
		c.internalCube = c.internalCube.Bdc()
		c.updateColors()
	}
	a := c.createAnimation(f, math.Pi/2, c.getBdMoveCubits(), axisY)
	c.animatingQueue = append(c.animatingQueue, a)
}

func (c *Cube) Rd() {
	f := func() {
		c.internalCube = c.internalCube.Rd()
		c.updateColors()
	}
	a := c.createAnimation(f, -math.Pi/2, c.getRdMoveCubits(), axisY)
	c.animatingQueue = append(c.animatingQueue, a)
}

func (c *Cube) Rdc() {
	f := func() {
		c.internalCube = c.internalCube.Rdc()
		c.updateColors()
	}
	a := c.createAnimation(f, math.Pi/2, c.getRdMoveCubits(), axisY)
	c.animatingQueue = append(c.animatingQueue, a)
}

func (c *Cube) Ld() {
	f := func() {
		c.internalCube = c.internalCube.Ld()
		c.updateColors()
	}
	a := c.createAnimation(f, math.Pi/2, c.getLdMoveCubits(), axisY)
	c.animatingQueue = append(c.animatingQueue, a)
}

func (c *Cube) Ldc() {
	f := func() {
		c.internalCube = c.internalCube.Ldc()
		c.updateColors()
	}
	a := c.createAnimation(f, -math.Pi/2, c.getLdMoveCubits(), axisY)
	c.animatingQueue = append(c.animatingQueue, a)
}
