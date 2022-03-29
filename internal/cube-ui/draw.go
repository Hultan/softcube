package cube_ui

import (
	"fmt"
	"image/color"
	"sort"
	"strings"

	"github.com/gotk3/gotk3/cairo"
	"github.com/gotk3/gotk3/gtk"

	rubik "github.com/hultan/softcube/internal/cube"
	"github.com/hultan/softcube/internal/surface"
	"github.com/hultan/softcube/internal/vector"
)

type axis int

const (
	axisX axis = iota
	axisY
	axisZ
)

var width, height float64
var thetaX, thetaY, thetaZ = 0.0, 0.0, 0.0
var cube rubik.Cube

const cubeSize = 1
const cubePosition = -1.5
const cubeDistance = 30.0
const distance = 5

var (
	red    = color.RGBA{R: 128, G: 0, B: 0, A: 255}
	blue   = color.RGBA{R: 0, G: 0, B: 128, A: 255}
	white  = color.RGBA{R: 255, G: 255, B: 255, A: 255}
	orange = color.RGBA{R: 225, G: 100, B: 0, A: 255}
	green  = color.RGBA{R: 4, G: 56, B: 4, A: 255}
	yellow = color.RGBA{R: 200, G: 200, B: 0, A: 255}
)

// onDraw : The onDraw signal handler
func (sc *SoftCube) onDraw(da *gtk.DrawingArea, ctx *cairo.Context) {
	width = float64(da.GetAllocatedWidth())
	height = float64(da.GetAllocatedHeight())

	cube = rubik.NewSolvedCube()
	//  R U R' U' R' F R2 U' R' U' R U R' F'
	cube = cube.R().U().Rc().Uc().Rc().F().R().R().Uc().Rc().Uc().R().U().Rc().Fc()

	sc.drawBackground(ctx)
	sc.drawCube(ctx, createCube(cube))
}

// drawBackground : Draws the background
func (sc *SoftCube) drawBackground(ctx *cairo.Context) {
	setColor(ctx, color.White)
	ctx.Rectangle(0, 0, width, height)
	ctx.Fill()
}

func (sc *SoftCube) drawCube(ctx *cairo.Context, surfaces []surface.Surface3) {
	// Rotate the cube
	var rotated []surface.Surface3
	for _, s := range surfaces {
		rotated = append(rotated, s.Rotate(thetaX, thetaY, thetaZ))
	}

	// Sort by Z-coord
	sort.Slice(rotated, func(i, j int) bool {
		return rotated[i].Z() > rotated[j].Z()
	})

	// Draw the cube
	for _, r := range rotated {
		// Calculate 2d coords
		s := r.To2DCoords(distance, cubeDistance)

		// Translate to screen coords
		s = s.ToScreenCoords(width, height)

		// Draw surface
		drawRectangle(ctx, true, 1, s, s.C)
		drawRectangle(ctx, false, 2, s, color.Black)
	}
}

func createSurface(a axis, x int, y int, z int, col color.Color) surface.Surface3 {
	switch a {
	case axisX:
		return surface.Surface3{
			V1: createVector(x, 3-y, z),
			V2: createVector(x, 3-y, z+1),
			V3: createVector(x, 3-y+1, z+1),
			V4: createVector(x, 3-y+1, z),
			C:  col,
		}
	case axisY:
		return surface.Surface3{
			V1: createVector(x, 3-y, z),
			V2: createVector(x+1, 3-y, z),
			V3: createVector(x+1, 3-y, z+1),
			V4: createVector(x, 3-y, z+1),
			C:  col,
		}
	case axisZ:
		return surface.Surface3{
			V1: createVector(x, 3-y, z),
			V2: createVector(x, 3-y+1, z),
			V3: createVector(x+1, 3-y+1, z),
			V4: createVector(x+1, 3-y, z),
			C:  col,
		}
	default:
		panic(fmt.Sprintf("invalid axis : %d", a))
	}
}

func createVector(x, y, z int) vector.Vector3 {
	return vector.Vector3{
		X: (float64(x) + cubePosition) * cubeSize,
		Y: (float64(y) + cubePosition) * cubeSize,
		Z: (float64(z) + cubePosition) * cubeSize,
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

func createCube(cube rubik.Cube) []surface.Surface3 {
	var s []surface.Surface3
	str := strings.Replace(cube.String(), " ", "", -1)

	// WHITE
	s = append(s, createSurface(axisY, 0, 3, 2, getColor(str[0])))
	s = append(s, createSurface(axisY, 1, 3, 2, getColor(str[1])))
	s = append(s, createSurface(axisY, 2, 3, 2, getColor(str[2])))

	s = append(s, createSurface(axisY, 0, 3, 1, getColor(str[3])))
	s = append(s, createSurface(axisY, 1, 3, 1, getColor(str[4])))
	s = append(s, createSurface(axisY, 2, 3, 1, getColor(str[5])))

	s = append(s, createSurface(axisY, 0, 3, 0, getColor(str[6])))
	s = append(s, createSurface(axisY, 1, 3, 0, getColor(str[7])))
	s = append(s, createSurface(axisY, 2, 3, 0, getColor(str[8])))

	// GREEN
	s = append(s, createSurface(axisZ, 0, 3, 0, getColor(str[9])))
	s = append(s, createSurface(axisZ, 1, 3, 0, getColor(str[10])))
	s = append(s, createSurface(axisZ, 2, 3, 0, getColor(str[11])))

	s = append(s, createSurface(axisZ, 0, 2, 0, getColor(str[12])))
	s = append(s, createSurface(axisZ, 1, 2, 0, getColor(str[13])))
	s = append(s, createSurface(axisZ, 2, 2, 0, getColor(str[14])))

	s = append(s, createSurface(axisZ, 0, 1, 0, getColor(str[15])))
	s = append(s, createSurface(axisZ, 1, 1, 0, getColor(str[16])))
	s = append(s, createSurface(axisZ, 2, 1, 0, getColor(str[17])))

	// RED
	s = append(s, createSurface(axisX, 3, 3, 0, getColor(str[18])))
	s = append(s, createSurface(axisX, 3, 3, 1, getColor(str[19])))
	s = append(s, createSurface(axisX, 3, 3, 2, getColor(str[20])))

	s = append(s, createSurface(axisX, 3, 2, 0, getColor(str[21])))
	s = append(s, createSurface(axisX, 3, 2, 1, getColor(str[22])))
	s = append(s, createSurface(axisX, 3, 2, 2, getColor(str[23])))

	s = append(s, createSurface(axisX, 3, 1, 0, getColor(str[24])))
	s = append(s, createSurface(axisX, 3, 1, 1, getColor(str[25])))
	s = append(s, createSurface(axisX, 3, 1, 2, getColor(str[26])))

	// BLUE
	s = append(s, createSurface(axisZ, 2, 3, 3, getColor(str[27])))
	s = append(s, createSurface(axisZ, 1, 3, 3, getColor(str[28])))
	s = append(s, createSurface(axisZ, 0, 3, 3, getColor(str[29])))

	s = append(s, createSurface(axisZ, 2, 2, 3, getColor(str[30])))
	s = append(s, createSurface(axisZ, 1, 2, 3, getColor(str[31])))
	s = append(s, createSurface(axisZ, 0, 2, 3, getColor(str[32])))

	s = append(s, createSurface(axisZ, 2, 1, 3, getColor(str[33])))
	s = append(s, createSurface(axisZ, 1, 1, 3, getColor(str[34])))
	s = append(s, createSurface(axisZ, 0, 1, 3, getColor(str[35])))

	// ORANGE
	s = append(s, createSurface(axisX, 0, 3, 0, getColor(str[36])))
	s = append(s, createSurface(axisX, 0, 3, 1, getColor(str[37])))
	s = append(s, createSurface(axisX, 0, 3, 2, getColor(str[38])))

	s = append(s, createSurface(axisX, 0, 2, 0, getColor(str[39])))
	s = append(s, createSurface(axisX, 0, 2, 1, getColor(str[40])))
	s = append(s, createSurface(axisX, 0, 2, 2, getColor(str[41])))

	s = append(s, createSurface(axisX, 0, 1, 0, getColor(str[42])))
	s = append(s, createSurface(axisX, 0, 1, 1, getColor(str[43])))
	s = append(s, createSurface(axisX, 0, 1, 2, getColor(str[44])))

	// YELLOW
	s = append(s, createSurface(axisY, 0, 0, 0, getColor(str[45])))
	s = append(s, createSurface(axisY, 1, 0, 0, getColor(str[46])))
	s = append(s, createSurface(axisY, 2, 0, 0, getColor(str[47])))

	s = append(s, createSurface(axisY, 0, 0, 1, getColor(str[48])))
	s = append(s, createSurface(axisY, 1, 0, 1, getColor(str[49])))
	s = append(s, createSurface(axisY, 2, 0, 1, getColor(str[50])))

	s = append(s, createSurface(axisY, 0, 0, 2, getColor(str[51])))
	s = append(s, createSurface(axisY, 1, 0, 2, getColor(str[52])))
	s = append(s, createSurface(axisY, 2, 0, 2, getColor(str[53])))

	return s
}

func drawRectangle(ctx *cairo.Context, fill bool, width float64, s surface.Surface2, col color.Color) {
	setColor(ctx, col)

	ctx.SetLineWidth(width)
	ctx.MoveTo(s.V1.X, s.V1.Y)
	ctx.LineTo(s.V2.X, s.V2.Y)
	ctx.LineTo(s.V3.X, s.V3.Y)
	ctx.LineTo(s.V4.X, s.V4.Y)
	ctx.LineTo(s.V1.X, s.V1.Y)

	if fill {
		ctx.Fill()
	} else {
		ctx.Stroke()
	}
}

func to2dCoords(v vector.Vector3) vector.Vector2 {
	return vector.Vector2{
		X: v.X * distance / (v.Z + cubeDistance),
		Y: v.Y * distance / (v.Z + cubeDistance),
	}
}
