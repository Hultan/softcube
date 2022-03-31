package cube_ui

import (
	"image/color"
	"sort"

	"github.com/gotk3/gotk3/cairo"
	"github.com/gotk3/gotk3/gtk"

	rubik "github.com/hultan/softcube/internal/cube"
	cube_alg "github.com/hultan/softcube/internal/cube-alg"
	"github.com/hultan/softcube/internal/surface"
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

// onDraw : The onDraw signal handler
func (sc *SoftCube) onDraw(da *gtk.DrawingArea, ctx *cairo.Context) {
	width = float64(da.GetAllocatedWidth())
	height = float64(da.GetAllocatedHeight())

	// Y Perm (altered)
	// alg := "R U' R' U' R U R' F' R U R' U' R' F R"
	// T Perm
	// alg := "R U R' U' R' F R2 U' R' U' R U R' F'"
	// R Perm
	// alg := "R U' R' U' R U R D R' U' R D' R' U2 R' U'"
	alg := "z'"
	// alg := "l"
	// alg = cube_alg.ReverseAlg(alg)
	cube = cube_alg.PerformAlg(rubik.NewSolvedCube(), alg)

	// //  R U R' U' R' F R2 U' R' U' R U R' F'
	// cube = cube.R().U().Rc().Uc().Rc().F().R().R().Uc().Rc().Uc().R().U().Rc().Fc()

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
		drawQuadrilateral(ctx, true, 1, s, s.C1)
		drawQuadrilateral(ctx, false, 2, s, color.Black)
	}
}
