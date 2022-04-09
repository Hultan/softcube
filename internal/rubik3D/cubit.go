package rubik3D

import (
	"image/color"

	"github.com/hultan/softcube/internal/surface"
	"github.com/hultan/softcube/internal/vector"
)

type Cubit struct {
	F, B *surface.Surface3
	U, D *surface.Surface3
	L, R *surface.Surface3
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
		B: &b,
		F: &f,
		U: &u,
		D: &d,
		L: &l,
		R: &r,
	}
}

func (c *Cubit) GetSurfaces() []*surface.Surface3 {
	return []*surface.Surface3{c.B, c.F, c.U, c.D, c.L, c.R}
}
