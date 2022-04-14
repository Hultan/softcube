package softcube

import (
	"fmt"
	"strconv"

	"github.com/gotk3/gotk3/gtk"

	alg "github.com/hultan/go-rubik/src/rubik-alg"
)

var pllButtons = []string{
	"buttonAaPerm", "buttonAbPerm", "buttonEPerm", "buttonFPerm", "buttonGaPerm",
	"buttonGbPerm", "buttonGcPerm", "buttonGdPerm", "buttonHPerm", "buttonJaPerm",
	"buttonJbPerm", "buttonNaPerm", "buttonNbPerm", "buttonRaPerm", "buttonRbPerm",
	"buttonTPerm", "buttonUaPerm", "buttonUbPerm", "buttonVPerm", "buttonYPerm",
	"buttonZPerm",
}

var ollAlg = []string{
	"F",
}

func getOLLButtonNames() []string {
	var buttons []string
	for i := 0; i < 56; i++ {
		buttons = append(buttons, fmt.Sprintf("buttonOLL%02d", i+1))
	}
	return buttons
}

func performOLL(btn *gtk.Button) {
	label, _ := btn.GetLabel()

	label = label[len(label)-2:]
	n, err := strconv.Atoi(label)
	if err != nil {
		fmt.Println(err)
		return
	}
	if n < 1 || n > len(ollAlg) {
		fmt.Println("invalid oll alg index : ", n)
		return
	}
	cube.ExecuteAlg(ollAlg[n-1])
}

func performPLL(btn *gtk.Button) {
	label, _ := btn.GetLabel()

	switch label {
	case "Aa perm":
		cube.ExecuteAlg(alg.PllPermAa)
	case "Ab perm":
		cube.ExecuteAlg(alg.PllPermAb)
	case "E perm":
		cube.ExecuteAlg(alg.PllPermE)
	case "F perm":
		cube.ExecuteAlg(alg.PllPermF)
	case "Ga perm":
		cube.ExecuteAlg(alg.PllPermGa)
	case "Gb perm":
		cube.ExecuteAlg(alg.PllPermGb)
	case "Gc perm":
		cube.ExecuteAlg(alg.PllPermGc)
	case "Gd perm":
		cube.ExecuteAlg(alg.PllPermGd)
	case "H perm":
		cube.ExecuteAlg(alg.PllPermH)
	case "Ja perm":
		cube.ExecuteAlg(alg.PllPermJa)
	case "Jb perm":
		cube.ExecuteAlg(alg.PllPermJb)
	case "Na perm":
		cube.ExecuteAlg(alg.PllPermNa)
	case "Nb perm":
		cube.ExecuteAlg(alg.PllPermNb)
	case "Ra perm":
		cube.ExecuteAlg(alg.PllPermRa)
	case "Rb perm":
		cube.ExecuteAlg(alg.PllPermRb)
	case "T perm":
		cube.ExecuteAlg(alg.PllPermT)
	case "Ua perm":
		cube.ExecuteAlg(alg.PllPermUa)
	case "Ub perm":
		cube.ExecuteAlg(alg.PllPermUb)
	case "V perm":
		cube.ExecuteAlg(alg.PllPermV)
	case "Y perm":
		cube.ExecuteAlg(alg.PllPermY)
	case "Z perm":
		cube.ExecuteAlg(alg.PllPermZ)
	}
}
