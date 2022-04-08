package cube_ui

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"

	"github.com/hultan/go-rubik/src/rubik"
	alg "github.com/hultan/go-rubik/src/rubik-alg"
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

func NewCube(b *framework.GtkBuilder, w *gtk.ApplicationWindow, da *gtk.DrawingArea) *SoftCube {
	t := &SoftCube{builder: b, window: w, drawingArea: da}
	t.window.Connect("key-press-event", t.onKeyPressed)

	for _, permButton := range permButtons {
		button := t.builder.GetObject(permButton).(*gtk.Button)
		button.Connect("clicked", performPermutation)
	}

	btns := getOLLButtons()
	for _, btn := range btns {
		button := t.builder.GetObject(btn).(*gtk.Button)
		button.Connect("clicked", performOLL)
	}

	return t
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
	colors = alg.ExecuteAlg(colors, ollAlg[n-1])
}

func performPermutation(btn *gtk.Button) {
	label, _ := btn.GetLabel()

	switch label {
	case "Aa perm":
		colors = alg.ExecuteAlg(colors, alg.PllPermAa)
	case "Ab perm":
		colors = alg.ExecuteAlg(colors, alg.PllPermAb)
	case "E perm":
		colors = alg.ExecuteAlg(colors, alg.PllPermE)
	case "F perm":
		colors = alg.ExecuteAlg(colors, alg.PllPermF)
	case "Ga perm":
		colors = alg.ExecuteAlg(colors, alg.PllPermGa)
	case "Gb perm":
		colors = alg.ExecuteAlg(colors, alg.PllPermGb)
	case "Gc perm":
		colors = alg.ExecuteAlg(colors, alg.PllPermGc)
	case "Gd perm":
		colors = alg.ExecuteAlg(colors, alg.PllPermGd)
	case "H perm":
		colors = alg.ExecuteAlg(colors, alg.PllPermH)
	case "Ja perm":
		colors = alg.ExecuteAlg(colors, alg.PllPermJa)
	case "Jb perm":
		colors = alg.ExecuteAlg(colors, alg.PllPermJb)
	case "Na perm":
		colors = alg.ExecuteAlg(colors, alg.PllPermNa)
	case "Nb perm":
		colors = alg.ExecuteAlg(colors, alg.PllPermNb)
	case "Ra perm":
		colors = alg.ExecuteAlg(colors, alg.PllPermRa)
	case "Rb perm":
		colors = alg.ExecuteAlg(colors, alg.PllPermRb)
	case "T perm":
		colors = alg.ExecuteAlg(colors, alg.PllPermT)
	case "Ua perm":
		colors = alg.ExecuteAlg(colors, alg.PllPermUa)
	case "Ub perm":
		colors = alg.ExecuteAlg(colors, alg.PllPermUb)
	case "V perm":
		colors = alg.ExecuteAlg(colors, alg.PllPermV)
	case "Y perm":
		colors = alg.ExecuteAlg(colors, alg.PllPermY)
	case "Z perm":
		colors = alg.ExecuteAlg(colors, alg.PllPermZ)
	}
}

func (sc *SoftCube) StartCube() {
	sc.drawingArea.Connect("draw", sc.onDraw)

	sc.ticker = time.NewTicker(50 * time.Millisecond)
	sc.tickerQuit = make(chan struct{})

	colors = rubik.NewSolvedCube()
	cube = createCube(colors)

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
		colors = rubik.NewSolvedCube()

	// Cube rotations
	case 120: // Button "x" => Rotate around X
		// colors = colors.X()
		thetaX += 0.1
	case 121: // Button "y" => Rotate around Y
		// colors = colors.Y()
		thetaY += 0.1
	case 122: // Button "z" => Rotate around Z
		// colors = colors.Z()
		thetaZ += 0.1
	case 88: // Button "X" => Rotate around X counter-clockwise
		// colors = colors.Xc()
		thetaX -= 0.1
	case 89: // Button "Y" => Rotate around Y counter-clockwise
		// colors = colors.Yc()
		thetaY -= 0.1
	case 90: // Button "Z" => Rotate around Z counter-clockwise
		// colors = colors.Zc()
		thetaZ -= 0.1

	// Turns
	case 117: // Button "u" => Move camera back
		colors = colors.U()
	case 100: // Button "d" => Move camera forward
		colors = colors.D()
	case 85: // Button "U" => Move camera left
		colors = colors.Uc()
	case 68: // Button "D" => Move camera right
		colors = colors.Dc()

	case 114: // Button "r" => Move camera back
		colors = colors.R()
	case 108: // Button "l" => Move camera forward
		colors = colors.L()
	case 82: // Button "R" => Move camera left
		colors = colors.Rc()
	case 76: // Button "L" => Move camera right
		colors = colors.Lc()

	case 102: // Button "f" => Move camera back
		colors = colors.F()
	case 98: // Button "b" => Move camera forward
		colors = colors.B()
	case 70: // Button "F" => Move camera left
		colors = colors.Fc()
	case 66: // Button "B" => Move camera right
		colors = colors.Bc()

	// TODO : Double turns CTRL + u,d,r,l,f,b

	// Slice moves
	case 109: // Button "m" => Reset rotation
		colors = colors.M()
	case 101: // Button "e" => Move camera left
		colors = colors.E()
	case 115: // Button "s" => Move camera left
		colors = colors.S()
	case 77: // Button "M" => Reset rotation
		colors = colors.Mc()
	case 69: // Button "E" => Move camera left
		colors = colors.Ec()
	case 83: // Button "S" => Move camera left
		colors = colors.Sc()
	}

	sc.drawingArea.QueueDraw()
}

func resetRotation() {
	thetaX = 0.2
	thetaY = -0.2
	thetaZ = 0
}

func getOLLButtons() []string {
	var btns []string

	for i := 0; i < 56; i++ {
		btns = append(btns, fmt.Sprintf("buttonOLL%02d", i+1))
	}
	return btns
}
