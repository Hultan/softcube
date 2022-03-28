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
	sc.setColor(ctx, color.White)
	ctx.Rectangle(0, 0, width, height)
	ctx.Fill()
}

func (sc *SoftCube) drawCube(ctx *cairo.Context, surfaces []surface.Surface) {
	d := 5.0

	var rotated []surface.Surface
	for _, s := range surfaces {
		// Rotate
		r1 := s.V1.RotateX(thetaX).RotateY(thetaY).RotateZ(thetaZ)
		r2 := s.V2.RotateX(thetaX).RotateY(thetaY).RotateZ(thetaZ)
		r3 := s.V3.RotateX(thetaX).RotateY(thetaY).RotateZ(thetaZ)
		r4 := s.V4.RotateX(thetaX).RotateY(thetaY).RotateZ(thetaZ)

		rotated = append(rotated, surface.Surface{
			V1: r1,
			V2: r2,
			V3: r3,
			V4: r4,
			C:  s.C,
		})
	}

	sort.Slice(rotated, func(i, j int) bool {
		return rotated[i].Z() > rotated[j].Z()
	})

	for _, r := range rotated {
		sc.setColor(ctx, r.C)

		// Calculate coords
		xr1 := r.V1.X * d / (r.V1.Z + cubeDistance)
		yr1 := r.V1.Y * d / (r.V1.Z + cubeDistance)
		xr2 := r.V2.X * d / (r.V2.Z + cubeDistance)
		yr2 := r.V2.Y * d / (r.V2.Z + cubeDistance)
		xr3 := r.V3.X * d / (r.V3.Z + cubeDistance)
		yr3 := r.V3.Y * d / (r.V3.Z + cubeDistance)
		xr4 := r.V4.X * d / (r.V4.Z + cubeDistance)
		yr4 := r.V4.Y * d / (r.V4.Z + cubeDistance)

		// Translate to screen coords
		xr1 = xr1*width + width/2
		yr1 = yr1*height + height/2
		xr2 = xr2*width + width/2
		yr2 = yr2*height + height/2
		xr3 = xr3*width + width/2
		yr3 = yr3*height + height/2
		xr4 = xr4*width + width/2
		yr4 = yr4*height + height/2

		// Draw surface
		ctx.SetLineWidth(1)
		ctx.MoveTo(xr1, yr1)
		ctx.LineTo(xr2, yr2)
		ctx.LineTo(xr3, yr3)
		ctx.LineTo(xr4, yr4)
		ctx.LineTo(xr1, yr1)
		ctx.Fill()

		sc.setColor(ctx, color.Black)

		ctx.SetLineWidth(1)
		ctx.MoveTo(xr1, yr1)
		ctx.LineTo(xr2, yr2)
		ctx.LineTo(xr3, yr3)
		ctx.LineTo(xr4, yr4)
		ctx.LineTo(xr1, yr1)
		ctx.Stroke()

		// thetaX += 0.0001
		// thetaY += 0.0001
		// thetaZ += 0.0001
	}
}

func createSurface(a axis, x1, y1, z1, x2, y2, z2 int, col color.Color) surface.Surface {
	switch a {
	case axisX:
		return surface.Surface{
			V1: createVector(x1, 3-y1, z1),
			V2: createVector(x1, 3-y1, z1+1),
			V3: createVector(x1, 3-y1+1, z1+1),
			V4: createVector(x1, 3-y1+1, z1),
			C:  col,
		}
	case axisY:
		return surface.Surface{
			V1: createVector(x1, 3-y1, z1),
			V2: createVector(x1+1, 3-y1, z1),
			V3: createVector(x1+1, 3-y1, z1+1),
			V4: createVector(x1, 3-y1, z1+1),
			C:  col,
		}
	case axisZ:
		return surface.Surface{
			V1: createVector(x1, 3-y1, z1),
			V2: createVector(x1, 3-y1+1, z1),
			V3: createVector(x1+1, 3-y1+1, z1),
			V4: createVector(x1+1, 3-y1, z1),
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
		return color.RGBA{R: 255, A: 255}
	case 'b':
		return color.RGBA{B: 255, A: 255}
	case 'w':
		return color.RGBA{R: 255, G: 255, B: 255, A: 255}
	case 'o':
		return color.RGBA{R: 255, G: 165, B: 0, A: 255}
	case 'g':
		return color.RGBA{G: 255, A: 255}
	case 'y':
		return color.RGBA{R: 255, G: 255, B: 100, A: 255}
	default:
		panic(fmt.Sprintf("invalid color : %c", ch))
	}
}

func createCube(cube rubik.Cube) []surface.Surface {
	var s []surface.Surface
	str := strings.Replace(cube.String(), " ", "", -1)

	// WHITE
	s = append(s, createSurface(axisY, 0, 3, 2, 1, 3, 1, getColor(str[0])))
	s = append(s, createSurface(axisY, 1, 3, 2, 2, 3, 1, getColor(str[1])))
	s = append(s, createSurface(axisY, 2, 3, 2, 3, 3, 1, getColor(str[2])))

	s = append(s, createSurface(axisY, 0, 3, 1, 1, 3, 0, getColor(str[3])))
	s = append(s, createSurface(axisY, 1, 3, 1, 2, 3, 0, getColor(str[4])))
	s = append(s, createSurface(axisY, 2, 3, 1, 3, 3, 0, getColor(str[5])))

	s = append(s, createSurface(axisY, 0, 3, 0, 1, 3, -1, getColor(str[6])))
	s = append(s, createSurface(axisY, 1, 3, 0, 2, 3, -1, getColor(str[7])))
	s = append(s, createSurface(axisY, 2, 3, 0, 3, 3, -1, getColor(str[8])))

	// GREEN
	s = append(s, createSurface(axisZ, 0, 3, 0, 1, 2, 0, getColor(str[9])))
	s = append(s, createSurface(axisZ, 1, 3, 0, 2, 2, 0, getColor(str[10])))
	s = append(s, createSurface(axisZ, 2, 3, 0, 3, 2, 0, getColor(str[11])))

	s = append(s, createSurface(axisZ, 0, 2, 0, 1, 1, 0, getColor(str[12])))
	s = append(s, createSurface(axisZ, 1, 2, 0, 2, 1, 0, getColor(str[13])))
	s = append(s, createSurface(axisZ, 2, 2, 0, 3, 1, 0, getColor(str[14])))

	s = append(s, createSurface(axisZ, 0, 1, 0, 1, 0, 0, getColor(str[15])))
	s = append(s, createSurface(axisZ, 1, 1, 0, 2, 0, 0, getColor(str[16])))
	s = append(s, createSurface(axisZ, 2, 1, 0, 3, 0, 0, getColor(str[17])))

	// RED
	s = append(s, createSurface(axisX, 3, 3, 0, 3, 2, 1, getColor(str[18])))
	s = append(s, createSurface(axisX, 3, 3, 1, 3, 2, 2, getColor(str[19])))
	s = append(s, createSurface(axisX, 3, 3, 2, 3, 2, 3, getColor(str[20])))

	s = append(s, createSurface(axisX, 3, 2, 0, 3, 1, 1, getColor(str[21])))
	s = append(s, createSurface(axisX, 3, 2, 1, 3, 1, 2, getColor(str[22])))
	s = append(s, createSurface(axisX, 3, 2, 2, 3, 1, 3, getColor(str[23])))

	s = append(s, createSurface(axisX, 3, 1, 0, 3, 0, 1, getColor(str[24])))
	s = append(s, createSurface(axisX, 3, 1, 1, 3, 0, 2, getColor(str[25])))
	s = append(s, createSurface(axisX, 3, 1, 2, 3, 0, 3, getColor(str[26])))

	// BLUE
	s = append(s, createSurface(axisZ, 2, 3, 3, 3, 2, 3, getColor(str[27])))
	s = append(s, createSurface(axisZ, 1, 3, 3, 2, 2, 3, getColor(str[28])))
	s = append(s, createSurface(axisZ, 0, 3, 3, 1, 2, 3, getColor(str[29])))

	s = append(s, createSurface(axisZ, 2, 2, 3, 3, 1, 3, getColor(str[30])))
	s = append(s, createSurface(axisZ, 1, 2, 3, 2, 1, 3, getColor(str[31])))
	s = append(s, createSurface(axisZ, 0, 2, 3, 1, 1, 3, getColor(str[32])))

	s = append(s, createSurface(axisZ, 2, 1, 3, 3, 0, 3, getColor(str[33])))
	s = append(s, createSurface(axisZ, 1, 1, 3, 2, 0, 3, getColor(str[34])))
	s = append(s, createSurface(axisZ, 0, 1, 3, 1, 0, 3, getColor(str[35])))

	// ORANGE
	s = append(s, createSurface(axisX, 0, 3, 0, 0, 2, 1, getColor(str[36])))
	s = append(s, createSurface(axisX, 0, 3, 1, 0, 2, 2, getColor(str[37])))
	s = append(s, createSurface(axisX, 0, 3, 2, 0, 2, 3, getColor(str[38])))

	s = append(s, createSurface(axisX, 0, 2, 0, 0, 1, 1, getColor(str[39])))
	s = append(s, createSurface(axisX, 0, 2, 1, 0, 1, 2, getColor(str[40])))
	s = append(s, createSurface(axisX, 0, 2, 2, 0, 1, 3, getColor(str[41])))

	s = append(s, createSurface(axisX, 0, 1, 0, 0, 0, 1, getColor(str[42])))
	s = append(s, createSurface(axisX, 0, 1, 1, 0, 0, 2, getColor(str[43])))
	s = append(s, createSurface(axisX, 0, 1, 2, 0, 0, 3, getColor(str[44])))

	// YELLOW
	s = append(s, createSurface(axisY, 0, 0, 0, 1, 0, 1, getColor(str[45])))
	s = append(s, createSurface(axisY, 1, 0, 0, 2, 0, 1, getColor(str[46])))
	s = append(s, createSurface(axisY, 2, 0, 0, 3, 0, 1, getColor(str[47])))

	s = append(s, createSurface(axisY, 0, 0, 1, 1, 0, 2, getColor(str[48])))
	s = append(s, createSurface(axisY, 1, 0, 1, 2, 0, 2, getColor(str[49])))
	s = append(s, createSurface(axisY, 2, 0, 1, 3, 0, 2, getColor(str[50])))

	s = append(s, createSurface(axisY, 0, 0, 2, 1, 0, 3, getColor(str[51])))
	s = append(s, createSurface(axisY, 1, 0, 2, 2, 0, 3, getColor(str[52])))
	s = append(s, createSurface(axisY, 2, 0, 2, 3, 0, 3, getColor(str[53])))

	return s
}
