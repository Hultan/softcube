package rubik3D

import (
	"fmt"
	"image/color"
	"math"
	"strings"

	"github.com/gotk3/gotk3/cairo"
	"github.com/gotk3/gotk3/gtk"

	"github.com/hultan/go-rubik/src/rubik"
	"github.com/hultan/softcube/internal/vector"
)

type Cube struct {
	BackgroundColor        color.Color
	AngleX, AngleY, AngleZ float64

	cubits       []Cubit
	internalCube rubik.Cube

	currentAnimation *animation
	animatingQueue   []*animation
}

var (
	red    = color.RGBA{R: 128, G: 0, B: 0, A: 255}
	blue   = color.RGBA{R: 0, G: 0, B: 128, A: 255}
	white  = color.RGBA{R: 255, G: 255, B: 255, A: 255}
	orange = color.RGBA{R: 225, G: 100, B: 0, A: 255}
	green  = color.RGBA{R: 24, G: 76, B: 24, A: 255}
	yellow = color.RGBA{R: 200, G: 200, B: 0, A: 255}
)

func NewCube() *Cube {
	c := &Cube{}

	for z := 0; z < 3; z++ {
		for y := 0; y < 3; y++ {
			for x := 0; x < 3; x++ {
				v1 := createVector(x, y+1, z)
				v2 := createVector(x+1, y+1, z)
				v3 := createVector(x, y+1, z+1)
				v4 := createVector(x+1, y+1, z+1)
				v5 := createVector(x, y, z)
				v6 := createVector(x+1, y, z)
				v7 := createVector(x, y, z+1)
				v8 := createVector(x+1, y, z+1)

				c.cubits = append(c.cubits, NewCubit(v1, v2, v3, v4, v5, v6, v7, v8))
			}
		}
	}

	c.BackgroundColor = color.White
	c.internalCube = rubik.NewSolvedCube()
	c.updateColors()

	return c
}

// ExecuteAlg executes the provided algorithm on the cube
func (c *Cube) ExecuteAlg(alg string) {
	// Remove ():s and []:s
	alg = c.cleanAlg(alg)

	// Split up alg, and execute moves
	moves := strings.Split(alg, " ")
	for _, move := range moves {
		move = strings.Trim(move, " ")
		c.executeMove(move)
	}
}

func (c *Cube) IsAnimating() bool {
	return c.currentAnimation != nil
}

func (c *Cube) Reset() {
	f := func() {
		c.internalCube = rubik.NewSolvedCube()
		c.updateColors()
	}
	a := c.createNonAnimation(f)
	c.animatingQueue = append(c.animatingQueue, a)
}

func (c *Cube) Draw(da *gtk.DrawingArea, ctx *cairo.Context) {
	width = float64(da.GetAllocatedWidth())
	height = float64(da.GetAllocatedHeight())

	c.drawBackground(ctx)
	c.drawCube(ctx)
}

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

//
// Private functions
//

func createVector(x, y, z int) vector.Vector3 {
	return vector.Vector3{
		X: float64(x) - 1.5,
		Y: float64(y) - 1.5,
		Z: float64(z) - 1.5,
	}
}

func getColor(ch byte) color.Color {
	switch strings.ToLower(string(ch))[0] {
	case 'r':
		return red
	case 'b':
		return blue
	case 'w':
		return white
	case 'o':
		return orange
	case 'g':
		return green
	case 'y':
		return yellow
	default:
		panic(fmt.Sprintf("invalid color : %c", ch))
	}
}

func (c *Cube) updateColors() {
	str := strings.Replace(c.internalCube.String(), " ", "", -1)
	fmt.Println(str)

	// Indexing of cube colors : https://github.com/yackx/go-rubik
	//
	// Indexing of cubits
	// 18  19  20
	// 21  22  23
	// 24  25  26
	//   \        \
	// 		9   10  11
	// 		12  13  14	// Cubit 13 is the center cubit, never visible
	// 		15  16  17
	//        \        \
	// 			0   1   2
	// 			3   4   5
	// 			6   7   8

	// White (up) side (on a solved cube)
	c.cubits[18].U.Col = getColor(str[0])
	c.cubits[19].U.Col = getColor(str[1])
	c.cubits[20].U.Col = getColor(str[2])
	c.cubits[9].U.Col = getColor(str[3])
	c.cubits[10].U.Col = getColor(str[4])
	c.cubits[11].U.Col = getColor(str[5])
	c.cubits[0].U.Col = getColor(str[6])
	c.cubits[1].U.Col = getColor(str[7])
	c.cubits[2].U.Col = getColor(str[8])

	// Green (front) side (on a solved cube) (since Y axis is flipped we use B instead of F here)
	c.cubits[0].B.Col = getColor(str[9])
	c.cubits[1].B.Col = getColor(str[10])
	c.cubits[2].B.Col = getColor(str[11])
	c.cubits[3].B.Col = getColor(str[12])
	c.cubits[4].B.Col = getColor(str[13])
	c.cubits[5].B.Col = getColor(str[14])
	c.cubits[6].B.Col = getColor(str[15])
	c.cubits[7].B.Col = getColor(str[16])
	c.cubits[8].B.Col = getColor(str[17])

	// Red (right) side (on a solved cube)
	c.cubits[2].R.Col = getColor(str[18])
	c.cubits[11].R.Col = getColor(str[19])
	c.cubits[20].R.Col = getColor(str[20])
	c.cubits[5].R.Col = getColor(str[21])
	c.cubits[14].R.Col = getColor(str[22])
	c.cubits[23].R.Col = getColor(str[23])
	c.cubits[8].R.Col = getColor(str[24])
	c.cubits[17].R.Col = getColor(str[25])
	c.cubits[26].R.Col = getColor(str[26])

	// Blue (back) side (on a solved cube) (since Y axis is flipped we use F instead of B here)
	c.cubits[20].F.Col = getColor(str[27])
	c.cubits[19].F.Col = getColor(str[28])
	c.cubits[18].F.Col = getColor(str[29])
	c.cubits[23].F.Col = getColor(str[30])
	c.cubits[22].F.Col = getColor(str[31])
	c.cubits[21].F.Col = getColor(str[32])
	c.cubits[26].F.Col = getColor(str[33])
	c.cubits[25].F.Col = getColor(str[34])
	c.cubits[24].F.Col = getColor(str[35])

	// Orange (left) side (on a solved cube)
	c.cubits[18].L.Col = getColor(str[36])
	c.cubits[9].L.Col = getColor(str[37])
	c.cubits[0].L.Col = getColor(str[38])
	c.cubits[21].L.Col = getColor(str[39])
	c.cubits[12].L.Col = getColor(str[40])
	c.cubits[3].L.Col = getColor(str[41])
	c.cubits[24].L.Col = getColor(str[42])
	c.cubits[15].L.Col = getColor(str[43])
	c.cubits[6].L.Col = getColor(str[44])

	// Yellow (down) side (on a solved cube)
	c.cubits[6].D.Col = getColor(str[45])
	c.cubits[7].D.Col = getColor(str[46])
	c.cubits[8].D.Col = getColor(str[47])
	c.cubits[15].D.Col = getColor(str[48])
	c.cubits[16].D.Col = getColor(str[49])
	c.cubits[17].D.Col = getColor(str[50])
	c.cubits[24].D.Col = getColor(str[51])
	c.cubits[25].D.Col = getColor(str[52])
	c.cubits[26].D.Col = getColor(str[53])
}

// getCubits returns a slice of copies of the cubits
func (c *Cube) getCubits() []Cubit {
	var cubits []Cubit

	for i := 0; i < 27; i++ {
		cubits = append(cubits, c.cubits[i])
	}

	return cubits
}

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

// cleanAlg removes spaces, enter, tabs, () and []
func (c *Cube) cleanAlg(alg string) string {
	alg = strings.Trim(alg, " \t\n\r")
	alg = strings.Replace(alg, "(", "", -1)
	alg = strings.Replace(alg, ")", "", -1)
	alg = strings.Replace(alg, "[", "", -1)
	alg = strings.Replace(alg, "]", "", -1)
	return alg
}
