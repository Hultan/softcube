package rubik3D

import (
	"sort"

	"github.com/gotk3/gotk3/cairo"
)

var width, height float64

const cubeDistance = 30.0
const distance = 5

// drawBackground : Draws the background
func (c *Cube) drawBackground(ctx *cairo.Context) {
	setColor(ctx, c.BackgroundColor)
	ctx.Rectangle(0, 0, width, height)
	ctx.Fill()
}

// drawCube : Draws the actual cube
func (c *Cube) drawCube(ctx *cairo.Context) {
	var cubits []Cubit

	// Rotate cubits
	for i := 0; i < 27; i++ {
		cubits = append(cubits, c.cubits[i].rotate(c.AngleX, c.AngleY, c.AngleZ))
	}

	// Sort by Z-coord
	// We want draw surfaces in the back first
	sort.Slice(cubits, func(i, j int) bool {
		return cubits[i].Z() > cubits[j].Z()
	})

	// Draw cubits
	for i := 0; i < 27; i++ {
		cubits[i].draw(ctx)
	}
}
