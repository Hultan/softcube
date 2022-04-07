package object

import (
	"image/color"
	"math/rand"

	"github.com/hultan/softcube/internal/surface"
	"github.com/hultan/softcube/internal/vector"
)

type Cubit struct {
	F, B *surface.Surface3
	U, D *surface.Surface3
	L, R *surface.Surface3
}

func NewCubit(LUB, RUB, LUF, RUF, LDB, RDB, LDF, RDF vector.Vector3) Cubit {
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
	u := surface.Surface3{
		V1:  LDB,
		V2:  RDB,
		V3:  RDF,
		V4:  LDF,
		Col: color.Black,
	}
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

func getRandomColor() color.Color {
	return color.RGBA{
		uint8(rand.Intn(255)),
		uint8(rand.Intn(255)),
		uint8(rand.Intn(255)),
		255,
	}
}

//
// func (c *Cubit) GetSurfaces() []surface.Surface3 {
// 	b := surface.Surface3{
// 		V1: c.LUB,
// 		V2: c.RUB,
// 		V3: c.RDB,
// 		V4: c.LDB,
// 		Col: c.ColB,
// 	}
// 	f := surface.Surface3{
// 		V1: c.LUF,
// 		V2: c.RUF,
// 		V3: c.RDF,
// 		V4: c.LDF,
// 		Col: c.ColF,
// 	}
// 	u := surface.Surface3{
// 		V1: c.LUB,
// 		V2: c.RUB,
// 		V3: c.RUF,
// 		V4: c.LUF,
// 		Col: c.ColU,
// 	}
// 	d := surface.Surface3{
// 		V1: c.LDB,
// 		V2: c.RDB,
// 		V3: c.RDF,
// 		V4: c.LDF,
// 		Col: c.ColD,
// 	}
// 	l := surface.Surface3{
// 		V1: c.LUB,
// 		V2: c.LUF,
// 		V3: c.LDF,
// 		V4: c.LDB,
// 		Col: c.ColL,
// 	}
// 	r := surface.Surface3{
// 		V1: c.RUB,
// 		V2: c.RUF,
// 		V3: c.RDF,
// 		V4: c.RDB,
// 		Col: c.ColR,
// 	}
// 	return []surface.Surface3{b, f, u, d, l, r}
// }
