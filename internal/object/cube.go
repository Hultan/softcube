package object

import (
	"image/color"
	"math/rand"

	"github.com/hultan/softcube/internal/surface"
	"github.com/hultan/softcube/internal/vector"
)

type Cube struct {
	LUB, RUB vector.Vector3
	LUF, RUF vector.Vector3
	LDB, RDB vector.Vector3
	LDF, RDF vector.Vector3

	ColB, ColF color.Color
	ColU, ColD color.Color
	ColL, ColR color.Color
}

func NewCube(LUB, RUB, LUF, RUF, LDB, RDB, LDF, RDF vector.Vector3) Cube {
	return Cube{
		LUB:  LUB,
		RUB:  RUB,
		LUF:  LUF,
		RUF:  RUF,
		LDB:  LDB,
		RDB:  RDB,
		LDF:  LDF,
		RDF:  RDF,
		ColB: getRandomColor(),
		ColF: getRandomColor(),
		ColU: getRandomColor(),
		ColD: getRandomColor(),
		ColL: getRandomColor(),
		ColR: getRandomColor(),
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

func (c *Cube) GetSurfaces() []surface.Surface3 {
	b := surface.Surface3{
		V1: c.LUB,
		V2: c.RUB,
		V3: c.RDB,
		V4: c.LDB,
		C1: c.ColB,
	}
	f := surface.Surface3{
		V1: c.LUF,
		V2: c.RUF,
		V3: c.RDF,
		V4: c.LDF,
		C1: c.ColF,
	}
	u := surface.Surface3{
		V1: c.LUB,
		V2: c.RUB,
		V3: c.RUF,
		V4: c.LUF,
		C1: c.ColU,
	}
	d := surface.Surface3{
		V1: c.LDB,
		V2: c.RDB,
		V3: c.RDF,
		V4: c.LDF,
		C1: c.ColD,
	}
	l := surface.Surface3{
		V1: c.LUB,
		V2: c.LUF,
		V3: c.LDF,
		V4: c.LDB,
		C1: c.ColL,
	}
	r := surface.Surface3{
		V1: c.RUB,
		V2: c.RUF,
		V3: c.RDF,
		V4: c.RDB,
		C1: c.ColR,
	}
	return []surface.Surface3{b, f, u, d, l, r}
}
