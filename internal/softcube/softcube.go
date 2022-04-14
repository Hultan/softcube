package softcube

import (
	"time"

	"github.com/gotk3/gotk3/cairo"
	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"

	"github.com/hultan/softcube/internal/rubik3D"
	"github.com/hultan/softteam/framework"
)

type SoftCube struct {
	builder     *framework.GtkBuilder
	window      *gtk.ApplicationWindow
	drawingArea *gtk.DrawingArea
	keyHandler  *keyHandler

	tickerQuit chan struct{}
	ticker     *time.Ticker
}

var cube *rubik3D.Cube

func NewCube(b *framework.GtkBuilder, w *gtk.ApplicationWindow, da *gtk.DrawingArea) *SoftCube {
	sc := &SoftCube{builder: b, window: w, drawingArea: da}
	sc.keyHandler = newKeyHandler()

	sc.window.Connect("key-press-event", sc.onKeyPressed)
	sc.window.Connect("key-release-event", sc.onKeyReleased)

	// PLL
	for _, pllButton := range pllButtons {
		button := sc.builder.GetObject(pllButton).(*gtk.Button)
		button.Connect("clicked", performPLL)
	}

	// OLL
	ollButtons := getOLLButtonNames()
	for _, ollButton := range ollButtons {
		btn := sc.builder.GetObject(ollButton).(*gtk.Button)
		btn.Connect("clicked", performOLL)
	}

	cube = rubik3D.NewCube()
	// cube.ExecuteAlg(alg.PllPermT)
	resetRotation()

	return sc
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
	sc.keyHandler.onKeyPressed(key.KeyVal())
}

// onKeyReleased : The onKeyPressed signal handler
func (sc *SoftCube) onKeyReleased(_ *gtk.ApplicationWindow, e *gdk.Event) {
	key := gdk.EventKeyNewFromEvent(e)

	if sc.keyHandler.onKeyReleased(key.KeyVal()) {
		sc.drawingArea.QueueDraw()
		return
	}

	switch key.KeyVal() {
	// Misc controls
	case keyq: // Button "Q" => Quit game
		sc.window.Close() // Close window
	}
}

func resetRotation() {
	cube.AngleX = 0.2
	cube.AngleY = -0.2
	cube.AngleZ = 0
}
