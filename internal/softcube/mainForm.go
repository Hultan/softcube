package softcube

import (
	"os"

	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"

	cube_ui "github.com/hultan/softcube/internal/cube-ui"
	"github.com/hultan/softteam/framework"
)

const applicationTitle = "softcube"
const applicationVersion = "v 0.01"
const applicationCopyRight = "©SoftTeam AB, 2020"

type MainForm struct {
	window      *gtk.ApplicationWindow
	builder     *framework.GtkBuilder
	aboutDialog *gtk.AboutDialog
	da          *gtk.DrawingArea
}

// NewMainForm : Creates a new MainForm object
func NewMainForm() *MainForm {
	mainForm := new(MainForm)
	return mainForm
}

// OpenMainForm : Opens the MainForm window
func (m *MainForm) OpenMainForm(app *gtk.Application) {
	// Initialize gtk
	gtk.Init(&os.Args)

	// Create a new softBuilder
	fw := framework.NewFramework()
	builder, err := fw.Gtk.CreateBuilder("main.ui")
	if err != nil {
		panic(err)
	}
	m.builder = builder

	// Get the main window from the glade file
	m.window = m.builder.GetObject("main_window").(*gtk.ApplicationWindow)

	// Set up main window
	m.window.Maximize()
	m.window.SetApplication(app)
	m.window.SetTitle("softcube main window")

	// Hook up the destroy event
	m.window.Connect("destroy", m.window.Close)

	// Quit button
	button := m.builder.GetObject("main_window_quit_button").(*gtk.ToolButton)
	button.Connect("clicked", m.window.Close)

	// Status bar
	statusBar := m.builder.GetObject("main_window_status_bar").(*gtk.Statusbar)
	statusBar.Push(statusBar.GetContextId("softcube"), "softcube : version 0.1.0")

	// Menu
	m.setupMenu(fw)

	// Drawing area
	m.da = m.builder.GetObject("drawingArea").(*gtk.DrawingArea)

	m.setupCSS()

	// Show the main window
	m.window.ShowAll()

	// Create new game object
	t := cube_ui.NewCube(m.window, m.da)
	t.StartCube()
}

func (m *MainForm) setupMenu(fw *framework.Framework) {
	menuQuit := m.builder.GetObject("menu_file_quit").(*gtk.MenuItem)
	menuQuit.Connect("activate", m.window.Close)
}

func (m *MainForm) setupCSS() {
	// Create CSS provider
	provider, _ := gtk.CssProviderNew()

	err := provider.LoadFromPath("assets/main.css")
	if err != nil {
		panic(err)
	}

	screen, err := gdk.ScreenGetDefault()
	gtk.AddProviderForScreen(screen, provider, gtk.STYLE_PROVIDER_PRIORITY_APPLICATION)
	// // Set CSS provider
	// box := m.builder.GetObject("mainBox").(*gtk.Box)
	// context, _ := box.GetStyleContext()
	// context.AddProvider(provider, gtk.STYLE_PROVIDER_PRIORITY_APPLICATION)
	// image := m.builder.GetObject("image").(*gtk.Image)
	// context, _ = image.GetStyleContext()
	// context.AddProvider(provider, gtk.STYLE_PROVIDER_PRIORITY_APPLICATION)
}
