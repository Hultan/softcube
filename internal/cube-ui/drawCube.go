package cube_ui

import (
	"fmt"
	"image/color"
	"strings"

	"github.com/hultan/go-rubik/src/rubik"
	"github.com/hultan/softcube/internal/object"
	"github.com/hultan/softcube/internal/surface"
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

func createCube(cube rubik.Cube) []object.Cube {
	// var s []surface.Surface3
	// str := strings.Replace(cube.String(), " ", "", -1)

	var c []object.Cube
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

				c = append(c, object.NewCube(v1, v2, v3, v4, v5, v6, v7, v8))
				fmt.Println(c)
			}
		}
	}
	//
	// // WHITE
	// s = append(s, createSurface(axisY, 0, 3, 2, getColor(str[0])))
	// s = append(s, createSurface(axisY, 1, 3, 2, getColor(str[1])))
	// s = append(s, createSurface(axisY, 2, 3, 2, getColor(str[2])))
	// s = append(s, createSurface(axisY, 0, 3, 1, getColor(str[3])))
	// s = append(s, createSurface(axisY, 1, 3, 1, getColor(str[4])))
	// s = append(s, createSurface(axisY, 2, 3, 1, getColor(str[5])))
	// s = append(s, createSurface(axisY, 0, 3, 0, getColor(str[6])))
	// s = append(s, createSurface(axisY, 1, 3, 0, getColor(str[7])))
	// s = append(s, createSurface(axisY, 2, 3, 0, getColor(str[8])))
	//
	// // GREEN
	// s = append(s, createSurface(axisZ, 0, 3, 0, getColor(str[9])))
	// s = append(s, createSurface(axisZ, 1, 3, 0, getColor(str[10])))
	// s = append(s, createSurface(axisZ, 2, 3, 0, getColor(str[11])))
	// s = append(s, createSurface(axisZ, 0, 2, 0, getColor(str[12])))
	// s = append(s, createSurface(axisZ, 1, 2, 0, getColor(str[13])))
	// s = append(s, createSurface(axisZ, 2, 2, 0, getColor(str[14])))
	// s = append(s, createSurface(axisZ, 0, 1, 0, getColor(str[15])))
	// s = append(s, createSurface(axisZ, 1, 1, 0, getColor(str[16])))
	// s = append(s, createSurface(axisZ, 2, 1, 0, getColor(str[17])))
	//
	// // RED
	// s = append(s, createSurface(axisX, 3, 3, 0, getColor(str[18])))
	// s = append(s, createSurface(axisX, 3, 3, 1, getColor(str[19])))
	// s = append(s, createSurface(axisX, 3, 3, 2, getColor(str[20])))
	// s = append(s, createSurface(axisX, 3, 2, 0, getColor(str[21])))
	// s = append(s, createSurface(axisX, 3, 2, 1, getColor(str[22])))
	// s = append(s, createSurface(axisX, 3, 2, 2, getColor(str[23])))
	// s = append(s, createSurface(axisX, 3, 1, 0, getColor(str[24])))
	// s = append(s, createSurface(axisX, 3, 1, 1, getColor(str[25])))
	// s = append(s, createSurface(axisX, 3, 1, 2, getColor(str[26])))
	//
	// // BLUE
	// s = append(s, createSurface(axisZ, 2, 3, 3, getColor(str[27])))
	// s = append(s, createSurface(axisZ, 1, 3, 3, getColor(str[28])))
	// s = append(s, createSurface(axisZ, 0, 3, 3, getColor(str[29])))
	// s = append(s, createSurface(axisZ, 2, 2, 3, getColor(str[30])))
	// s = append(s, createSurface(axisZ, 1, 2, 3, getColor(str[31])))
	// s = append(s, createSurface(axisZ, 0, 2, 3, getColor(str[32])))
	// s = append(s, createSurface(axisZ, 2, 1, 3, getColor(str[33])))
	// s = append(s, createSurface(axisZ, 1, 1, 3, getColor(str[34])))
	// s = append(s, createSurface(axisZ, 0, 1, 3, getColor(str[35])))
	//
	// // ORANGE
	// s = append(s, createSurface(axisX, 0, 3, 2, getColor(str[36])))
	// s = append(s, createSurface(axisX, 0, 3, 1, getColor(str[37])))
	// s = append(s, createSurface(axisX, 0, 3, 0, getColor(str[38])))
	// s = append(s, createSurface(axisX, 0, 2, 2, getColor(str[39])))
	// s = append(s, createSurface(axisX, 0, 2, 1, getColor(str[40])))
	// s = append(s, createSurface(axisX, 0, 2, 0, getColor(str[41])))
	// s = append(s, createSurface(axisX, 0, 1, 2, getColor(str[42])))
	// s = append(s, createSurface(axisX, 0, 1, 1, getColor(str[43])))
	// s = append(s, createSurface(axisX, 0, 1, 0, getColor(str[44])))
	//
	// // YELLOW
	// s = append(s, createSurface(axisY, 0, 0, 0, getColor(str[45])))
	// s = append(s, createSurface(axisY, 1, 0, 0, getColor(str[46])))
	// s = append(s, createSurface(axisY, 2, 0, 0, getColor(str[47])))
	// s = append(s, createSurface(axisY, 0, 0, 1, getColor(str[48])))
	// s = append(s, createSurface(axisY, 1, 0, 1, getColor(str[49])))
	// s = append(s, createSurface(axisY, 2, 0, 1, getColor(str[50])))
	// s = append(s, createSurface(axisY, 0, 0, 2, getColor(str[51])))
	// s = append(s, createSurface(axisY, 1, 0, 2, getColor(str[52])))
	// s = append(s, createSurface(axisY, 2, 0, 2, getColor(str[53])))

	return c
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

func createSurface(a axis, x int, y int, z int, col1 color.Color) surface.Surface3 {
	// Invert the Y-axis
	y = 3 - y

	switch a {
	case axisX:
		return surface.Surface3{
			V1: createVector(x, y, z),
			V2: createVector(x, y, z+1),
			V3: createVector(x, y+1, z+1),
			V4: createVector(x, y+1, z),
			C1: col1,
		}
	case axisY:
		return surface.Surface3{
			V1: createVector(x, y, z),
			V2: createVector(x+1, y, z),
			V3: createVector(x+1, y, z+1),
			V4: createVector(x, y, z+1),
			C1: col1,
		}
	case axisZ:
		return surface.Surface3{
			V1: createVector(x, y, z),
			V2: createVector(x, y+1, z),
			V3: createVector(x+1, y+1, z),
			V4: createVector(x+1, y, z),
			C1: col1,
		}
	default:
		panic(fmt.Sprintf("invalid axis : %d", a))
	}
}

func createVector(x, y, z int) vector.Vector3 {
	return vector.Vector3{
		X: float64(x) - 1.5,
		Y: float64(y) - 1.5,
		Z: float64(z) - 1.5,
	}
}
