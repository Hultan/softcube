package softcube

type keyHandler struct {
	handlers         map[uint]func()
	ctrlHandlers     map[uint]func()
	isCtrlKeyPressed bool
	rotate           bool
}

func newKeyHandler() *keyHandler {
	k := &keyHandler{isCtrlKeyPressed: false, rotate: true}
	k.setupKeyHandlers()
	return k
}

func (h *keyHandler) setupKeyHandlers() {
	h.handlers = map[uint]func(){
		// Misc controls
		keyESC: func() { cube.Reset() }, // ESC

		// Cube rotations

		// Button "x" => Rotate around X
		keyx: func() {
			if h.rotate {
				cube.X()
			} else {
				cube.AngleX += 0.1
			}
		},
		// Button "y" => Rotate around y
		keyy: func() {
			if h.rotate {
				cube.Y()
			} else {
				cube.AngleY += 0.1
			}
		},
		// Button "z" => Rotate around z
		keyz: func() {
			if h.rotate {
				cube.Z()
			} else {
				cube.AngleZ += 0.1
			}
		},
		// Button "X" => Rotate around X counterclockwise
		keyX: func() {
			if h.rotate {
				cube.Xc()
			} else {
				cube.AngleX -= 0.1
			}
		},
		// Button "Y" => Rotate around y counterclockwise
		keyY: func() {
			if h.rotate {
				cube.Yc()
			} else {
				cube.AngleY -= 0.1
			}
		},
		// Button "Z" => Rotate around z counterclockwise
		keyZ: func() {
			if h.rotate {
				cube.Zc()
			} else {
				cube.AngleZ -= 0.1
			}
		},

		// Turns
		keyu: func() { cube.U() },  // Button "u" => U turn
		keyd: func() { cube.D() },  // Button "d" => D turn
		keyU: func() { cube.Uc() }, // Button "U" => U' turn
		keyD: func() { cube.Dc() }, // Button "D" => D' turn

		keyr: func() { cube.R() },  // Button "r" => R turn
		keyl: func() { cube.L() },  // Button "l" => L turn
		keyR: func() { cube.Rc() }, // Button "R" => R' turn
		keyL: func() { cube.Lc() }, // Button "L" => L' turn

		keyf: func() { cube.F() },  // Button "f" => F turn
		keyb: func() { cube.B() },  // Button "b" => B turn
		keyF: func() { cube.Fc() }, // Button "F" => F' turn
		keyB: func() { cube.Bc() }, // Button "B" => B' turn

		// Slice moves
		keym: func() { cube.M() },  // Button "m" => M turn
		keye: func() { cube.E() },  // Button "e" => E turn
		keys: func() { cube.S() },  // Button "s" => S turn
		keyM: func() { cube.Mc() }, // Button "M" => M' turn
		keyE: func() { cube.Ec() }, // Button "E" => E' turn
		keyS: func() { cube.Sc() }, // Button "S" => S' turn
	}

	h.ctrlHandlers = map[uint]func(){
		// Double turns
		keyu: func() { cube.Ud() },  // Button CTRL + "u" => u turn
		keyd: func() { cube.Dd() },  // Button CTRL + "d" => d turn
		keyU: func() { cube.Udc() }, // Button CTRL + "U" => u' turn
		keyD: func() { cube.Ddc() }, // Button CTRL + "D" => d' turn

		keyr: func() { cube.Rd() },  // Button CTRL + "r" => r turn
		keyl: func() { cube.Ld() },  // Button CTRL + "l" => l turn
		keyR: func() { cube.Rdc() }, // Button CTRL + "R" => r' turn
		keyL: func() { cube.Ldc() }, // Button CTRL + "L" => l' turn

		keyf: func() { cube.Fd() },  // Button CTRL + "f" => f turn
		keyb: func() { cube.Bd() },  // Button CTRL + "b" => b turn
		keyF: func() { cube.Fdc() }, // Button CTRL + "F" => f' turn
		keyB: func() { cube.Bdc() }, // Button CTRL + "B" => b' turn
	}
}

func (h *keyHandler) onKeyPressed(key uint) {
	if key == keyCTRL {
		h.isCtrlKeyPressed = true
	}
}

func (h *keyHandler) onKeyReleased(key uint) bool {
	// Handle CTRL key release
	switch key {
	case keyCTRL: // Button "CTRL"
		h.isCtrlKeyPressed = false
		return true
	}

	// Pick the correct handlers map (is CTRL pressed)
	var handlers map[uint]func()
	if h.isCtrlKeyPressed {
		handlers = h.ctrlHandlers
	} else {
		handlers = h.handlers
	}

	// Execute handler for the pressed key
	f, ok := handlers[key]
	if ok {
		f()
	}
	return ok
}
