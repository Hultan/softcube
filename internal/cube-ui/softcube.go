package cube_ui

import (
	"fmt"
	"time"

	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"

	"github.com/hultan/go-rubik/src/rubik"
	rubik_alg "github.com/hultan/go-rubik/src/rubik-alg"
	"github.com/hultan/softteam/framework"
)

type SoftCube struct {
	builder     *framework.GtkBuilder
	window      *gtk.ApplicationWindow
	drawingArea *gtk.DrawingArea

	tickerQuit chan struct{}
	ticker     *time.Ticker
}

var permButtons = []string{
	"buttonAaPerm", "buttonAbPerm", "buttonEPerm", "buttonFPerm", "buttonGaPerm",
	"buttonGbPerm", "buttonGcPerm", "buttonGdPerm", "buttonHPerm", "buttonJaPerm",
	"buttonJbPerm", "buttonNaPerm", "buttonNbPerm", "buttonRaPerm", "buttonRbPerm",
	"buttonTPerm", "buttonUaPerm", "buttonUbPerm", "buttonVPerm", "buttonYPerm",
	"buttonZPerm",
}

func NewCube(b *framework.GtkBuilder, w *gtk.ApplicationWindow, da *gtk.DrawingArea) *SoftCube {
	t := &SoftCube{builder: b, window: w, drawingArea: da}
	t.window.Connect("key-press-event", t.onKeyPressed)

	for _, permButton := range permButtons {
		button := t.builder.GetObject(permButton).(*gtk.Button)
		button.Connect("clicked", performPermutation)
	}

	return t
}

func performPermutation(btn *gtk.Button) {
	label, _ := btn.GetLabel()

	switch label {
	case "Aa perm":
		cube = rubik_alg.ExecuteAlg(cube, rubik_alg.PllPermAa)
	case "Ab perm":
		cube = rubik_alg.ExecuteAlg(cube, rubik_alg.PllPermAb)
	case "E perm":
		cube = rubik_alg.ExecuteAlg(cube, rubik_alg.PllPermE)
	case "F perm":
		cube = rubik_alg.ExecuteAlg(cube, rubik_alg.PllPermF)
	case "Ga perm":
		cube = rubik_alg.ExecuteAlg(cube, rubik_alg.PllPermGa)
	case "Gb perm":
		cube = rubik_alg.ExecuteAlg(cube, rubik_alg.PllPermGb)
	case "Gc perm":
		cube = rubik_alg.ExecuteAlg(cube, rubik_alg.PllPermGc)
	case "Gd perm":
		cube = rubik_alg.ExecuteAlg(cube, rubik_alg.PllPermGd)
	case "H perm":
		cube = rubik_alg.ExecuteAlg(cube, rubik_alg.PllPermH)
	case "Ja perm":
		cube = rubik_alg.ExecuteAlg(cube, rubik_alg.PllPermJa)
	case "Jb perm":
		cube = rubik_alg.ExecuteAlg(cube, rubik_alg.PllPermJb)
	case "Na perm":
		cube = rubik_alg.ExecuteAlg(cube, rubik_alg.PllPermNa)
	case "Nb perm":
		cube = rubik_alg.ExecuteAlg(cube, rubik_alg.PllPermNb)
	case "Ra perm":
		cube = rubik_alg.ExecuteAlg(cube, rubik_alg.PllPermRa)
	case "Rb perm":
		cube = rubik_alg.ExecuteAlg(cube, rubik_alg.PllPermRb)
	case "T perm":
		cube = rubik_alg.ExecuteAlg(cube, rubik_alg.PllPermT)
	case "Ua perm":
		cube = rubik_alg.ExecuteAlg(cube, rubik_alg.PllPermUa)
	case "Ub perm":
		cube = rubik_alg.ExecuteAlg(cube, rubik_alg.PllPermUb)
	case "V perm":
		cube = rubik_alg.ExecuteAlg(cube, rubik_alg.PllPermV)
	case "Y perm":
		cube = rubik_alg.ExecuteAlg(cube, rubik_alg.PllPermY)
	case "Z perm":
		cube = rubik_alg.ExecuteAlg(cube, rubik_alg.PllPermZ)
	}
}

func (sc *SoftCube) StartCube() {
	sc.drawingArea.Connect("draw", sc.onDraw)

	sc.ticker = time.NewTicker(50 * time.Millisecond)
	sc.tickerQuit = make(chan struct{})

	cube = rubik.NewSolvedCube()
	// alg := rubikAlg.PllPermGb
	// alg = cube_alg.ReverseAlg(alg)
	// cube = rubik.NewCube("oywbwbroy yboggrbrb bygwrywwr rrygbobbg gwggowoyw rgooyowry")
	// alg := cube_alg.OLL_9
	// cube = rubikAlg.ExecuteAlg(cube, alg)

	resetRotation()

	go sc.mainLoop()
}

func (sc *SoftCube) mainLoop() {
	for {
		select {
		case <-sc.ticker.C:
			sc.drawingArea.QueueDraw()
		case <-sc.tickerQuit:
			sc.ticker.Stop()
			return
		}
	}
}

// onKeyPressed : The onKeyPressed signal handler
func (sc *SoftCube) onKeyPressed(_ *gtk.ApplicationWindow, e *gdk.Event) {
	key := gdk.EventKeyNewFromEvent(e)

	fmt.Println(key.KeyVal())

	switch key.KeyVal() {
	// Misc controls
	case 113: // Button "Q" => Quit game
		sc.window.Close() // Close window
	case 65307: // Button "Esc" => Reset to solved cube
		cube = rubik.NewSolvedCube()

	// Cube rotations
	case 120: // Button "x" => Reset rotation
		cube = cube.X()
	case 121: // Button "y" => Move camera left
		cube = cube.Y()
	case 122: // Button "z" => Move camera left
		cube = cube.Z()
	case 88: // Button "X" => Reset rotation
		cube = cube.Xc()
	case 89: // Button "Y" => Move camera left
		cube = cube.Yc()
	case 90: // Button "Z" => Move camera left
		cube = cube.Zc()

	// Turns
	case 117: // Button "u" => Move camera back
		cube = cube.U()
	case 100: // Button "d" => Move camera forward
		cube = cube.D()
	case 85: // Button "U" => Move camera left
		cube = cube.Uc()
	case 68: // Button "D" => Move camera right
		cube = cube.Dc()

	case 114: // Button "r" => Move camera back
		cube = cube.R()
	case 108: // Button "l" => Move camera forward
		cube = cube.L()
	case 82: // Button "R" => Move camera left
		cube = cube.Rc()
	case 76: // Button "L" => Move camera right
		cube = cube.Lc()

	case 102: // Button "f" => Move camera back
		cube = cube.F()
	case 98: // Button "b" => Move camera forward
		cube = cube.B()
	case 70: // Button "F" => Move camera left
		cube = cube.Fc()
	case 66: // Button "B" => Move camera right
		cube = cube.Bc()

	// TODO : Double turns CTRL + u,d,r,l,f,b

	// Slice moves
	case 109: // Button "m" => Reset rotation
		cube = cube.M()
	case 101: // Button "e" => Move camera left
		cube = cube.E()
	case 115: // Button "s" => Move camera left
		cube = cube.S()
	case 77: // Button "M" => Reset rotation
		cube = cube.Mc()
	case 69: // Button "E" => Move camera left
		cube = cube.Ec()
	case 83: // Button "S" => Move camera left
		cube = cube.Sc()
	}

	sc.drawingArea.QueueDraw()
}

func resetRotation() {
	thetaX = 0.2
	thetaY = -0.2
	thetaZ = 0
}
