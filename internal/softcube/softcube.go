package softcube

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gotk3/gotk3/cairo"
	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"

	alg "github.com/hultan/go-rubik/src/rubik-alg"
	"github.com/hultan/softcube/internal/rubik3D"
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

var ollAlg = []string{
	"F",
}

var cube *rubik3D.Cube

func NewCube(b *framework.GtkBuilder, w *gtk.ApplicationWindow, da *gtk.DrawingArea) *SoftCube {
	t := &SoftCube{builder: b, window: w, drawingArea: da}
	t.window.Connect("key-press-event", t.onKeyPressed)

	for _, permButton := range permButtons {
		button := t.builder.GetObject(permButton).(*gtk.Button)
		button.Connect("clicked", performPLL)
	}

	buttons := getOLLButtonNames()
	for _, button := range buttons {
		btn := t.builder.GetObject(button).(*gtk.Button)
		btn.Connect("clicked", performOLL)
	}

	cube = rubik3D.NewCube()
	cube.R()
	cube.R()
	cube.R()
	cube.R()
	cube.Rc()
	cube.Rc()
	cube.Rc()
	cube.Rc()
	resetRotation()

	return t
}

func (sc *SoftCube) StartCube() {
	sc.drawingArea.Connect("draw", sc.onDraw)

	sc.ticker = time.NewTicker(50 * time.Millisecond)
	sc.tickerQuit = make(chan struct{})

	go sc.mainLoop()
}

//
// Private functions
//

// onDraw : The onDraw signal handler
func (sc *SoftCube) onDraw(da *gtk.DrawingArea, ctx *cairo.Context) {
	cube.Draw(da, ctx)
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

	rotate := true

	fmt.Println(key.KeyVal())

	switch key.KeyVal() {
	// Misc controls
	case 113: // Button "Q" => Quit game
		sc.window.Close() // Close window
	case 65307: // Button "Esc" => Reset to solved cube
		cube.Reset()

	// Cube rotations
	case 120: // Button "x" => Rotate around X
		if rotate {
			cube.X()
		} else {
			cube.AngleX += 0.1
		}
	case 121: // Button "y" => Rotate around Y
		if rotate {
			cube.Y()
		} else {
			cube.AngleY += 0.1
		}
	case 122: // Button "z" => Rotate around Z
		if rotate {
			cube.Z()
		} else {
			cube.AngleZ += 0.1
		}
	case 88: // Button "X" => Rotate around X counter-clockwise
		if rotate {
			cube.Xc()
		} else {
			cube.AngleX -= 0.1
		}
	case 89: // Button "Y" => Rotate around Y counter-clockwise
		if rotate {
			cube.Yc()
		} else {
			cube.AngleY -= 0.1
		}
	case 90: // Button "Z" => Rotate around Z counter-clockwise
		if rotate {
			cube.Zc()
		} else {
			cube.AngleZ -= 0.1
		}

	// Turns
	case 117: // Button "u" => Move camera back
		cube.U()
	case 100: // Button "d" => Move camera forward
		cube.D()
	case 85: // Button "U" => Move camera left
		cube.Uc()
	case 68: // Button "D" => Move camera right
		cube.Dc()

	case 114: // Button "r" => Move camera back
		cube.R()
	case 108: // Button "l" => Move camera forward
		cube.L()
	case 82: // Button "R" => Move camera left
		cube.Rc()
	case 76: // Button "L" => Move camera right
		cube.Lc()

	case 102: // Button "f" => Move camera back
		cube.F()
	case 98: // Button "b" => Move camera forward
		cube.B()
	case 70: // Button "F" => Move camera left
		cube.Fc()
	case 66: // Button "B" => Move camera right
		cube.Bc()

	// TODO : Double turns CTRL + u,d,r,l,f,b

	// Slice moves
	case 109: // Button "m" => Reset rotation
		cube.M()
	case 101: // Button "e" => Move camera left
		cube.E()
	case 115: // Button "s" => Move camera left
		cube.S()
	case 77: // Button "M" => Reset rotation
		cube.Mc()
	case 69: // Button "E" => Move camera left
		cube.Ec()
	case 83: // Button "S" => Move camera left
		cube.Sc()
	}

	sc.drawingArea.QueueDraw()
}

func resetRotation() {
	cube.AngleX = 0.2
	cube.AngleY = -0.2
	cube.AngleZ = 0
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
