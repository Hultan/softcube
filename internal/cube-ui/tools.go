package cube_ui

import (
	"image/color"

	"github.com/gotk3/gotk3/cairo"

	"github.com/hultan/softcube/internal/surface"
)

func setColor(ctx *cairo.Context, c color.Color) {
	r, g, b, a := c.RGBA()
	ctx.SetSourceRGBA(col(r), col(g), col(b), col(a))
}

// Convert 0-65535 color to 0-1 color
func col(c uint32) float64 {
	return float64(c) / 65535
}

func drawQuadrilateral(ctx *cairo.Context, fill bool, width float64, s surface.Surface2, col color.Color) {
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
