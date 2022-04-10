package rubik3D

import (
	"image/color"
	"math"
	"sort"

	"github.com/gotk3/gotk3/cairo"

	"github.com/hultan/softcube/internal/surface"
	"github.com/hultan/softcube/internal/vector"
)

type Cubit struct {
	F, B surface.Surface3
	U, D surface.Surface3
	L, R surface.Surface3
}

func NewCubit(LUB, RUB, LUF, RUF, LDB, RDB, LDF, RDF vector.Vector3) Cubit {

	// Corners of the cube:
	//
	// LUB --- RUB
	//  |  \	|  \
	// 	|	LUF-|-- RUF
	//  |	 |	|	 |
	// LDB --| RDB	 |
	//	  \	 |	  \  |
	// 		LDF --- RDF

	b := surface.Surface3{
		V1:  LUB,
		V2:  RUB,
		V3:  RDB,
		V4:  LDB,
		Col: color.Black,
	}
	f := surface.Surface3{
		V1:  LUF,
		V2:  RUF,
		V3:  RDF,
		V4:  LDF,
		Col: color.Black,
	}
	// NOTE : Switched U and D surface, since Y axis is reverted
	u := surface.Surface3{
		V1:  LDB,
		V2:  RDB,
		V3:  RDF,
		V4:  LDF,
		Col: color.Black,
	}
	// NOTE : Switched U and D surface, since Y axis is reverted
	d := surface.Surface3{
		V1:  LUB,
		V2:  RUB,
		V3:  RUF,
		V4:  LUF,
		Col: color.Black,
	}
	l := surface.Surface3{
		V1:  LUB,
		V2:  LUF,
		V3:  LDF,
		V4:  LDB,
		Col: color.Black,
	}
	r := surface.Surface3{
		V1:  RUB,
		V2:  RUF,
		V3:  RDF,
		V4:  RDB,
		Col: color.Black,
	}
	return Cubit{
		B: b,
		F: f,
		U: u,
		D: d,
		L: l,
		R: r,
	}
}

// rotate : Returns a new rotated cubit
func (c *Cubit) rotate(x, y, z float64) Cubit {
	return Cubit{
		B: c.B.Rotate(x, y, z),
		F: c.F.Rotate(x, y, z),
		U: c.U.Rotate(x, y, z),
		D: c.D.Rotate(x, y, z),
		L: c.L.Rotate(x, y, z),
		R: c.R.Rotate(x, y, z),
	}
}

// getSurfaces : Returns a slice of surfaces
func (c *Cubit) getSurfaces() []surface.Surface3 {
	return []surface.Surface3{c.B, c.F, c.U, c.D, c.L, c.R}
}

func (c *Cubit) Z() float64 {
	max := math.Max(c.B.Z(), c.F.Z())
	max = math.Max(max, c.U.Z())
	max = math.Max(max, c.D.Z())
	max = math.Max(max, c.R.Z())
	max = math.Max(max, c.L.Z())

	return max
}

func (c *Cubit) draw(ctx *cairo.Context) {
	s := c.getSurfaces()

	// Sort by Z-coord
	// We want draw surfaces in the back first
	sort.Slice(s, func(i, j int) bool {
		return s[i].Z() > s[j].Z()
	})

	// Draw the cube
	for _, r := range s {
		// Calculate 2d coords
		surface2D := r.To2DCoords(distance, cubeDistance)

		// Translate to screen coords
		surface2D = surface2D.ToScreenCoords(width, height)

		// Draw surface
		drawQuadrilateral(ctx, true, 1, surface2D, surface2D.C1)
		drawQuadrilateral(ctx, false, 2, surface2D, color.Black)
	}
}
