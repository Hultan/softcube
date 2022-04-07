package cube_ui

import (
	"fmt"
	"image/color"
	"strings"

	"github.com/hultan/go-rubik/src/rubik"
	"github.com/hultan/softcube/internal/object"
	"github.com/hultan/softcube/internal/vector"
)

var (
	red    = color.RGBA{R: 128, G: 0, B: 0, A: 255}
	blue   = color.RGBA{R: 0, G: 0, B: 128, A: 255}
	white  = color.RGBA{R: 255, G: 255, B: 255, A: 255}
	orange = color.RGBA{R: 225, G: 100, B: 0, A: 255}
	green  = color.RGBA{R: 24, G: 76, B: 24, A: 255}
	yellow = color.RGBA{R: 200, G: 200, B: 0, A: 255}
)

func createCube(colors rubik.Cube) []object.Cubit {
	var c []object.Cubit
	for z := 0; z < 3; z++ {
		for y := 0; y < 3; y++ {
			for x := 0; x < 3; x++ {
				v1 := createVector(x, y+1, z)
				v2 := createVector(x+1, y+1, z)
				v3 := createVector(x, y+1, z+1)
				v4 := createVector(x+1, y+1, z+1)
				v5 := createVector(x, y, z)
				v6 := createVector(x+1, y, z)
				v7 := createVector(x, y, z+1)
				v8 := createVector(x+1, y, z+1)

				c = append(c, object.NewCubit(v1, v2, v3, v4, v5, v6, v7, v8))
			}
		}
	}

	str := strings.Replace(colors.String(), " ", "", -1)
	fmt.Println(str)

	c[0].U.Col = getColor(str[0])
	c[1].U.Col = getColor(str[1])
	c[2].U.Col = getColor(str[2])
	c[9].U.Col = getColor(str[3])
	c[10].U.Col = getColor(str[4])
	c[11].U.Col = getColor(str[5])
	c[18].U.Col = getColor(str[6])
	c[19].U.Col = getColor(str[7])
	c[20].U.Col = getColor(str[8])

	c[0].B.Col = getColor(str[9])
	c[1].B.Col = getColor(str[10])
	c[2].B.Col = getColor(str[11])
	c[3].B.Col = getColor(str[12])
	c[4].B.Col = getColor(str[13])
	c[5].B.Col = getColor(str[14])
	c[6].B.Col = getColor(str[15])
	c[7].B.Col = getColor(str[16])
	c[8].B.Col = getColor(str[17])

	c[2].R.Col = getColor(str[18])
	c[5].R.Col = getColor(str[19])
	c[8].R.Col = getColor(str[20])
	c[11].R.Col = getColor(str[21])
	c[14].R.Col = getColor(str[22])
	c[17].R.Col = getColor(str[23])
	c[20].R.Col = getColor(str[24])
	c[23].R.Col = getColor(str[25])
	c[26].R.Col = getColor(str[26])

	c[18].F.Col = getColor(str[27])
	c[19].F.Col = getColor(str[28])
	c[20].F.Col = getColor(str[29])
	c[21].F.Col = getColor(str[30])
	c[22].F.Col = getColor(str[31])
	c[23].F.Col = getColor(str[32])
	c[24].F.Col = getColor(str[33])
	c[25].F.Col = getColor(str[34])
	c[26].F.Col = getColor(str[35])

	c[0].L.Col = getColor(str[36])
	c[3].L.Col = getColor(str[37])
	c[6].L.Col = getColor(str[38])
	c[9].L.Col = getColor(str[39])
	c[12].L.Col = getColor(str[40])
	c[15].L.Col = getColor(str[41])
	c[18].L.Col = getColor(str[42])
	c[21].L.Col = getColor(str[43])
	c[24].L.Col = getColor(str[44])

	c[6].D.Col = getColor(str[45])
	c[7].D.Col = getColor(str[46])
	c[8].D.Col = getColor(str[47])
	c[15].D.Col = getColor(str[48])
	c[16].D.Col = getColor(str[49])
	c[17].D.Col = getColor(str[50])
	c[24].D.Col = getColor(str[51])
	c[25].D.Col = getColor(str[52])
	c[26].D.Col = getColor(str[53])

	return c
}

func createVector(x, y, z int) vector.Vector3 {
	return vector.Vector3{
		X: float64(x) - 1.5,
		Y: float64(y) - 1.5,
		Z: float64(z) - 1.5,
	}
}

func getColor(ch byte) color.Color {
	switch strings.ToLower(string(ch))[0] {
	case 'r':
		return red
	case 'b':
		return blue
	case 'w':
		return white
	case 'o':
		return orange
	case 'g':
		return green
	case 'y':
		return yellow
	default:
		panic(fmt.Sprintf("invalid color : %c", ch))
	}
}
