package cube_ui

import (
	"fmt"
	"time"

	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"
)

type SoftCube struct {
	window      *gtk.ApplicationWindow
	drawingArea *gtk.DrawingArea

	tickerQuit chan struct{}
	ticker     *time.Ticker
}

func NewCube(w *gtk.ApplicationWindow, da *gtk.DrawingArea) *SoftCube {
	t := &SoftCube{window: w, drawingArea: da}
	t.window.Connect("key-press-event", t.onKeyPressed)

	return t
}

func (sc *SoftCube) StartCube() {
	sc.drawingArea.Connect("draw", sc.onDraw)

	sc.ticker = time.NewTicker(50 * time.Millisecond)
	sc.tickerQuit = make(chan struct{})

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
	case 113: // Button "Q" => Quit game
		sc.window.Close() // Close window
	case 115: // Button "S" => Move camera back
		thetaX -= 0.05
	case 119: // Button "W" => Move camera forward
		thetaX += 0.05
	case 97: // Button "A" => Move camera left
		thetaY -= 0.05
	case 100: // Button "D" => Move camera right
		thetaY += 0.05
	case 122: // Button "Z" => Move camera left
		thetaZ -= 0.05
	case 99: // Button "C" => Move camera right
		thetaZ += 0.05
	}

	sc.drawingArea.QueueDraw()
}
