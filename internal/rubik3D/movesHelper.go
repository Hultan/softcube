package rubik3D

func (c *Cube) getAllCubits() []int {
	var cubits []int

	for i := 0; i < 27; i++ {
		cubits = append(cubits, i)
	}
	return cubits
}

func (c *Cube) getUMoveCubits() []int {
	return []int{18, 19, 20, 9, 10, 11, 0, 1, 2}
}

func (c *Cube) getDMoveCubits() []int {
	return []int{6, 7, 8, 15, 16, 17, 24, 25, 26}
}

func (c *Cube) getRMoveCubits() []int {
	return []int{2, 11, 20, 5, 14, 23, 8, 17, 26}
}

func (c *Cube) getLMoveCubits() []int {
	return []int{18, 9, 0, 21, 12, 3, 24, 15, 6}
}

func (c *Cube) getFMoveCubits() []int {
	return []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
}

func (c *Cube) getBMoveCubits() []int {
	return []int{18, 19, 20, 21, 22, 23, 24, 25, 26}
}

func (c *Cube) getMMoveCubits() []int {
	return []int{1, 4, 7, 10, 13, 16, 19, 22, 25}
}

func (c *Cube) getEMoveCubits() []int {
	return []int{3, 4, 5, 12, 13, 14, 21, 22, 23}
}

func (c *Cube) getSMoveCubits() []int {
	return []int{9, 10, 11, 12, 13, 14, 15, 16, 17}
}

func (c *Cube) getUdMoveCubits() []int {
	return []int{18, 19, 20, 9, 10, 11, 0, 1, 2, 3, 4, 5, 12, 13, 14, 21, 22, 23}
}

func (c *Cube) getDdMoveCubits() []int {
	return []int{3, 4, 5, 12, 13, 14, 21, 22, 23, 6, 7, 8, 15, 16, 17, 24, 25, 26}
}

func (c *Cube) getFdMoveCubits() []int {
	return []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17}
}

func (c *Cube) getBdMoveCubits() []int {
	return []int{9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26}
}

func (c *Cube) getRdMoveCubits() []int {
	return []int{2, 11, 20, 5, 14, 23, 8, 17, 26, 1, 4, 7, 10, 13, 16, 19, 22, 25}
}

func (c *Cube) getLdMoveCubits() []int {
	return []int{1, 4, 7, 10, 13, 16, 19, 22, 25, 18, 9, 0, 21, 12, 3, 24, 15, 6}
}

// Performs the provided move on the cube.
// Illegal moves are ignored.
func (c *Cube) executeMove(move string) {
	switch move {
	case "R":
		c.R()
	case "R2":
		c.R()
		c.R()
	case "R2'":
		c.Rc()
		c.Rc()
	case "R'":
		c.Rc()
	case "L":
		c.L()
	case "L2":
		c.L()
		c.L()
	case "L2'":
		c.Lc()
		c.Lc()
	case "L'":
		c.Lc()

	case "U":
		c.U()
	case "U2":
		c.U()
		c.U()
	case "U2'":
		c.Uc()
		c.Uc()
	case "U'":
		c.Uc()
	case "D":
		c.D()
	case "D2":
		c.D()
		c.D()
	case "D2'":
		c.Dc()
		c.Dc()
	case "D'":
		c.Dc()

	case "F":
		c.F()
	case "F2":
		c.F()
		c.F()
	case "F2'":
		c.Fc()
		c.Fc()
	case "F'":
		c.Fc()
	case "B":
		c.B()
	case "B2":
		c.B()
		c.B()
	case "B2'":
		c.Bc()
		c.Bc()
	case "B'":
		c.Bc()

	case "M":
		c.M()
	case "M2":
		c.M()
		c.M()
	case "M2'":
		c.Mc()
		c.Mc()
	case "M'":
		c.Mc()
	case "E":
		c.E()
	case "E2":
		c.E()
		c.E()
	case "E2'":
		c.Ec()
		c.Ec()
	case "E'":
		c.Ec()
	case "S":
		c.S()
	case "S2":
		c.S()
		c.S()
	case "S2'":
		c.Sc()
		c.Sc()
	case "S'":
		c.Sc()

	case "u":
		c.Ud()
	case "u'":
		c.Udc()
	case "u2":
		c.Ud()
		c.Ud()
	case "u2'":
		c.Udc()
		c.Udc()
	case "d":
		c.Dd()
	case "d'":
		c.Ddc()
	case "d2":
		c.Dd()
		c.Dd()
	case "d2'":
		c.Ddc()
		c.Ddc()

	case "f":
		c.Fd()
	case "f'":
		c.Fdc()
	case "f2":
		c.Fd()
		c.Fd()
	case "f2'":
		c.Fdc()
		c.Fdc()
	case "b":
		c.Bd()
	case "b'":
		c.Bdc()
	case "b2":
		c.Bd()
		c.Bd()
	case "b2'":
		c.Bdc()
		c.Bdc()

	case "l":
		c.Ld()
	case "l'":
		c.Ldc()
	case "l2":
		c.Ld()
		c.Ld()
	case "l2'":
		c.Ldc()
		c.Ldc()
	case "r":
		c.Rd()
	case "r'":
		c.Rdc()
	case "r2":
		c.Rd()
		c.Rd()
	case "r2'":
		c.Rdc()
		c.Rdc()

	case "x":
		c.X()
	case "x2":
		c.X()
		c.X()
	case "x2'":
		c.Xc()
		c.Xc()
	case "x'":
		c.Xc()
	case "y":
		c.Y()
	case "y2":
		c.Y()
		c.Y()
	case "y2'":
		c.Yc()
		c.Yc()
	case "y'":
		c.Yc()
	case "z":
		c.Z()
	case "z2":
		c.Z()
		c.Z()
	case "z2'":
		c.Zc()
		c.Zc()
	case "z'":
		c.Zc()

	default:
		// Invalid moves are ignored
	}
}
