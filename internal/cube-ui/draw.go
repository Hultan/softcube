package cube_ui

import (
	"image/color"
	"sort"

	"github.com/gotk3/gotk3/cairo"
	"github.com/gotk3/gotk3/gtk"

	"github.com/hultan/softcube/internal/rubik3D"
	"github.com/hultan/softcube/internal/surface"
)

var width, height float64
var thetaX, thetaY, thetaZ = 0.0, 0.0, 0.0
var cube *rubik3D.Cube

const cubeDistance = 30.0
const distance = 5

// onDraw : The onDraw signal handler
func (sc *SoftCube) onDraw(da *gtk.DrawingArea, ctx *cairo.Context) {
	width = float64(da.GetAllocatedWidth())
	height = float64(da.GetAllocatedHeight())

	sc.drawBackground(ctx)
	sc.drawCube(ctx)
}

// drawBackground : Draws the background
func (sc *SoftCube) drawBackground(ctx *cairo.Context) {
	setColor(ctx, color.White)
	ctx.Rectangle(0, 0, width, height)
	ctx.Fill()
}

func (sc *SoftCube) drawCube(ctx *cairo.Context) {
	var sf []*surface.Surface3

	// Collect all the surfaces
	for _, o := range cube.GetCubits() {
		sf = append(sf, o.GetSurfaces()...)
	}

	// TODO : Rotate layer here?

	// Rotate the cube
	var rotated []surface.Surface3
	for _, s := range sf {
		rotated = append(rotated, s.Rotate(thetaX, thetaY, thetaZ))
	}

	// Sort by Z-coord
	// We want draw surfaces in the back first
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
