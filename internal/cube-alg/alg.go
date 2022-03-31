package cube_alg

import (
	"fmt"
	"strings"

	rubik "github.com/hultan/softcube/internal/cube"
)

// Aa : x R' U R' D2 R U' R' D2 R2
// Ab : x R2 D2 R U R' D2 R U' R x'
// E  : z U2 R2 F R U R' U' R U R' U' R U R' U' F' R2 U2 z'
// F  : y R' U' F' R U R' U' R' F R2 U' R' U' R U R' U R
// Ga : R2 u R' U R' U' R u' R2 y' R' U R
// Gb : L' U' L y L2 U L' U L U' L u' L2
// Gc : L2 u' L U' L U L' u L2 y L U' L'
// Gd : R U R' y' R2 u' R U' R' U R' u R2
// H  : M2 U M2 U2 M2 U M2
// Ja : (R' U L' U2) (R U' R' U2 R) L [U']
// Jb : (R U R' F') (R U R' U') R' F R2 U' R' [U']
// Na : L U' R U2 L' U R' L U' R U2 L' U R' U
// Nb : R' U L' U2 R U' L R' U L' U2 R U' L U'
// Ra : (R U R' F') (R U2' R' U2') (R' F R U) (R U2' R') [U']
// Rb : (R' U2 R U2') R' F (R U R' U') R' F' R2 [U']
// T  : (R U R' U') (R' F R2 U') R' U' (R U R' F')
// Ua : y2 R U' R U R U R U' R' U' R2
// Ub : y2 R2 U R U R' U' R' U' R' U R'
// V  : R' U R' U' y R' F' R2 U' R' U R' F R F
// Y  : R2 U' R2 U' R2 U R' F' R U R2 U' R' F R
// Z  : M2 U M2 U M' U2 M2 U2 M'

func ReverseAlg(alg string) string {
	alg = cleanAlg(alg)
	moves := strings.Split(alg, " ")
	var reverse []string

	for _, move := range moves {
		move = strings.Trim(move, " ")
		if move == "" {
			continue
		}
		reverse = append(reverse, reverseMove(move))
	}

	for i, j := 0, len(reverse)-1; i < j; i, j = i+1, j-1 {
		reverse[i], reverse[j] = reverse[j], reverse[i]
	}

	return strings.Join(reverse, " ")
}

func PerformAlg(cube rubik.Cube, alg string) rubik.Cube {
	alg = cleanAlg(alg)
	moves := strings.Split(alg, " ")

	for _, move := range moves {
		move = strings.Trim(move, " ")
		if move == "" {
			continue
		}
		cube = performMove(cube, move)
	}

	return cube
}

func reverseMove(move string) string {
	switch move {
	case "R":
		return "R'"
	case "R2":
		return "R2"
	case "R'":
		return "R"
	case "L":
		return "L1"
	case "L2":
		return "L2"
	case "L'":
		return "L"

	case "U":
		return "U'"
	case "U2":
		return "U2"
	case "U'":
		return "U"
	case "D":
		return "D'"
	case "D2":
		return "D2"
	case "D'":
		return "D"

	case "F":
		return "F'"
	case "F2":
		return "F2"
	case "F'":
		return "F"
	case "B":
		return "B'"
	case "B2":
		return "B2"
	case "B'":
		return "B"

	case "M":
		return "M'"
	case "M2":
		return "M2"
	case "M'":
		return "M"
	case "E":
		return "E'"
	case "E2":
		return "E2"
	case "E'":
		return "E"
	case "S":
		return "S'"
	case "S2":
		return "S2"
	case "S'":
		return "S"

	case "u":
		return "u'"
	case "u'":
		return "u"
	case "u2":
		return "u2"
	case "d":
		return "d'"
	case "d'":
		return "d"
	case "d2":
		return "d2"

	case "f":
		return "f'"
	case "f'":
		return "f"
	case "f2":
		return "f2"
	case "b":
		return "b'"
	case "b'":
		return "b"
	case "b2":
		return "b2"

	case "l":
		return "l'"
	case "l'":
		return "l"
	case "l2":
		return "l2"
	case "r":
		return "r'"
	case "r'":
		return "r"
	case "r2":
		return "r2"

	default:
		panic(fmt.Sprintf("illegal move : %v", move))
	}
}

func performMove(cube rubik.Cube, move string) rubik.Cube {
	switch move {
	case "R":
		return cube.R()
	case "R2":
		return cube.R().R()
	case "R'":
		return cube.Rc()
	case "L":
		return cube.L()
	case "L2":
		return cube.L().L()
	case "L'":
		return cube.Lc()

	case "U":
		return cube.U()
	case "U2":
		return cube.U().U()
	case "U'":
		return cube.Uc()
	case "D":
		return cube.D()
	case "D2":
		return cube.D().D()
	case "D'":
		return cube.Dc()

	case "F":
		return cube.F()
	case "F2":
		return cube.F().F()
	case "F'":
		return cube.Fc()
	case "B":
		return cube.B()
	case "B2":
		return cube.B().B()
	case "B'":
		return cube.Bc()

	case "M":
		return cube.M()
	case "M2":
		return cube.M().M()
	case "M'":
		return cube.Mc()
	case "E":
		return cube.E()
	case "E2":
		return cube.E().E()
	case "E'":
		return cube.Ec()
	case "S":
		return cube.S()
	case "S2":
		return cube.S().S()
	case "S'":
		return cube.Sc()

	case "u":
		return cube.Ec().U()
	case "u'":
		return cube.E().Uc()
	case "u2":
		return cube.U().U().Ec().Ec()
	case "d":
		return cube.E().D()
	case "d'":
		return cube.Ec().Dc()
	case "d2":
		return cube.D().D().E().E()

	case "f":
		return cube.S().F()
	case "f'":
		return cube.Sc().Fc()
	case "f2":
		return cube.S().S().F().F()
	case "b":
		return cube.Sc().B()
	case "b'":
		return cube.S().Bc()
	case "b2":
		return cube.Sc().Sc().B().B()

	case "l":
		return cube.L().M()
	case "l'":
		return cube.Lc().Mc()
	case "l2":
		return cube.L().L().M().M()
	case "r":
		return cube.R().Mc()
	case "r'":
		return cube.Rc().M()
	case "r2":
		return cube.R().R().Mc().Mc()

	case "x":
		return cube.L().M().Rc()
	case "x'":
		return cube.Lc().Mc().R()
	case "y":
		return cube.U().Ec().Dc()
	case "y'":
		return cube.Uc().E().D()
	case "z":
		return cube.F().S().Bc()
	case "z'":
		return cube.Fc().Sc().B()

	default:
		panic(fmt.Sprintf("illegal move : %v", move))
	}
}

func cleanAlg(alg string) string {
	alg = strings.Replace(alg, "(", "", -1)
	alg = strings.Replace(alg, ")", "", -1)
	alg = strings.Replace(alg, "[", "", -1)
	alg = strings.Replace(alg, "]", "", -1)
	return alg
}
