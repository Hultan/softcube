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
const animationSteps = 20

type axis int

const (
	axisX axis = iota
	axisY
	axisZ
)

type animation struct {
	afterAnimation func()
	step           int
	angle          float64
	startAngle     float64
	endAngle       float64
	cubits         []int
	axis           axis
}

// drawBackground : Draws the background
func (c *Cube) drawBackground(ctx *cairo.Context) {
	setColor(ctx, c.BackgroundColor)
	ctx.Rectangle(0, 0, width, height)
	ctx.Fill()
}

// drawCube : Draws the actual cube
func (c *Cube) drawCube(ctx *cairo.Context) {
	cubits := c.getCubits()

	if !c.IsAnimating() {
		if c.currentAnimation == nil && len(c.animatingQueue) > 0 {
			c.currentAnimation = c.animatingQueue[0]
			c.animatingQueue = c.animatingQueue[1:]
		}
	} else {
		c.currentAnimation.step += 1
		c.currentAnimation.angle += (c.currentAnimation.endAngle - c.currentAnimation.startAngle) / animationSteps
		if c.currentAnimation.step == animationSteps {
			c.currentAnimation.angle = c.currentAnimation.endAngle
			defer func() {
				c.currentAnimation.afterAnimation()
				c.currentAnimation = nil
			}()
		}

		// Animate
		for _, i := range c.currentAnimation.cubits {
			switch c.currentAnimation.axis {
			case axisX:
				cubits[i] = cubits[i].rotate(c.currentAnimation.angle, 0, 0)
			case axisY:
				cubits[i] = cubits[i].rotate(0, c.currentAnimation.angle, 0)
			case axisZ:
				cubits[i] = cubits[i].rotate(0, 0, c.currentAnimation.angle)
			}

		}
	}

	// Rotate cubits
	for i := 0; i < 27; i++ {
		cubits[i] = cubits[i].rotate(c.AngleX, c.AngleY, c.AngleZ)
	}

	// Collect surfaces
	var s []surface.Surface3
	for _, cubit := range cubits {
		s = append(s, cubit.getSurfaces()...)
	}

	// Sort by Z-coord
	// We want draw surfaces in the back first
	sort.Slice(s, func(i, j int) bool {
		return s[i].Z() > s[j].Z()
	})

	// Draw surfaces
	for i := 0; i < len(s); i++ {
		surface2D := s[i].To2DCoords(distance, cubeDistance)

		// Translate to screen coords
		surface2D = surface2D.ToScreenCoords(width, height)

		// Draw surface
		drawQuadrilateral(ctx, true, 1, surface2D, surface2D.C1)
		drawQuadrilateral(ctx, false, 2, surface2D, color.Black)
	}
}
