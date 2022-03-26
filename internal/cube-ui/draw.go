package cube_ui

import (
	"image/color"
	"sort"

	"github.com/gotk3/gotk3/cairo"
	"github.com/gotk3/gotk3/gtk"
)

type vector3d struct {
	x, y, z float64
}

var camera = vector3d{0, 0, -10}
var screen = vector3d{-0.5, -0.5, -5}
var width, height float64
var thetaX, thetaY, thetaZ = 0.0, 0.0, 0.0

const cubeSize = 1
const cubePosition = -1.5
const cubeDistance = 30.0

// onDraw : The onDraw signal handler
func (sc *SoftCube) onDraw(da *gtk.DrawingArea, ctx *cairo.Context) {
	width = float64(da.GetAllocatedWidth())
	height = float64(da.GetAllocatedHeight())

	var cube []vector3d

	for x := 0; x <= 3; x++ {
		for y := 0; y <= 3; y++ {
			for z := 0; z <= 3; z++ {
				if x == 0 || x == 3 || y == 0 || y == 3 || z == 0 || z == 3 {
					v := vector3d{
						x: (float64(x) + cubePosition) * cubeSize,
						y: (float64(y) + cubePosition) * cubeSize,
						z: (float64(z) + cubePosition) * cubeSize,
					}

					cube = append(cube, v)
				}
			}
		}
	}
	
	sc.drawBackground(da, ctx)
	sc.drawCube(da, ctx, cube, color.RGBA{R: 255, A: 255})
}

// drawBackground : Draws the background
func (sc *SoftCube) drawBackground(da *gtk.DrawingArea, ctx *cairo.Context) {
	sc.setColor(ctx, color.White)
	ctx.Rectangle(0, 0, width, height)
	ctx.Fill()
}

func (sc *SoftCube) drawCube(da *gtk.DrawingArea, ctx *cairo.Context, cube []vector3d, rgba color.RGBA) {
	sc.setColor(ctx, rgba)
	d := 5.0

	sort.Slice(cube, func(i, j int) bool {
		return cube[i].z < cube[j].z
	})

	for _, s := range cube {
		// Rotate
		r := s.rotateX(thetaX).rotateY(thetaY).rotateZ(thetaZ)

		// Calculate coords
		x := r.x * d / (r.z + cubeDistance)
		y := r.y * d / (r.z + cubeDistance)

		// Translate to screen coords
		x = x*width + width/2
		y = y*height + height/2

		// Draw point
		ctx.Rectangle(x-3, y-3, 6, 6)
		ctx.Fill()

		thetaX += 0.0001
		thetaY += 0.0001
		thetaZ += 0.0001
	}
}
