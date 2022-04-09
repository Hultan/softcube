package rubik3D

import (
	"image/color"
	"sort"

	"github.com/gotk3/gotk3/cairo"

	"github.com/hultan/softcube/internal/surface"
)

var width, height float64

const cubeDistance = 30.0
const distance = 5

// drawBackground : Draws the background
func (c *Cube) drawBackground(ctx *cairo.Context) {
	setColor(ctx, color.White)
	ctx.Rectangle(0, 0, width, height)
	ctx.Fill()
}

// drawCube : Draws the actual cube
func (c *Cube) drawCube(ctx *cairo.Context) {
	var surfaces []*surface.Surface3

	// Collect all the surfaces
	for _, cubit := range c.cubits {
		surfaces = append(surfaces, cubit.getSurfaces()...)
	}

	// TODO : Rotate layer here?

	// Rotate the cube
	var rotatedSurfaces []surface.Surface3
	for _, s := range surfaces {
		rotatedSurfaces = append(rotatedSurfaces, s.Rotate(c.ThetaX, c.ThetaY, c.ThetaZ))
	}

	// Sort by Z-coord
	// We want draw surfaces in the back first
	sort.Slice(rotatedSurfaces, func(i, j int) bool {
		return rotatedSurfaces[i].Z() > rotatedSurfaces[j].Z()
	})

	// Draw the cube
	for _, r := range rotatedSurfaces {
		// Calculate 2d coords
		surface2D := r.To2DCoords(distance, cubeDistance)

		// Translate to screen coords
		surface2D = surface2D.ToScreenCoords(width, height)

		// Draw surface
		drawQuadrilateral(ctx, true, 1, surface2D, surface2D.C1)
		drawQuadrilateral(ctx, false, 2, surface2D, color.Black)
	}
}
