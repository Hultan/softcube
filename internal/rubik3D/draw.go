package rubik3D

import (
	"sort"

	"github.com/gotk3/gotk3/cairo"
)

var width, height float64

const cubeDistance = 30.0
const distance = 5
const animationSteps = 20

type animation struct {
	afterAnimation func()
	step           int
	angle          float64
	startAngle     float64
	endAngle       float64
	cubits         []int
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
			cubits[i] = cubits[i].rotate(c.currentAnimation.angle, 0, 0)
		}
	}

	// Rotate cubits
	for i := 0; i < 27; i++ {
		cubits[i] = cubits[i].rotate(c.AngleX, c.AngleY, c.AngleZ)
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
