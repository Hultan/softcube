package cube_ui

import (
	"image/color"

	"github.com/gotk3/gotk3/cairo"
)

func (sc *SoftCube) setColor(ctx *cairo.Context, c color.Color) {
	r, g, b, a := c.RGBA()
	ctx.SetSourceRGBA(col(r), col(g), col(b), col(a))
}

// Convert 0-65535 color to 0-1 color
func col(c uint32) float64 {
	return float64(c) / 65535
}
