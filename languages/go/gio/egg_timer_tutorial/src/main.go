package main

import (
	"gioui.org/app"
	"gioui.org/io/system"
	"gioui.org/unit"
)

// ----------------------------------------------------------------------------
const (
	MAIN_WIDTH  = 640
	MAIN_HEIGHT = 480
)

// ----------------------------------------------------------------------------
func main() {
	go func() {
		// create new window
		mainWindow := new(app.Window)

		// set window properties
		mainWindow.Option(app.Title("Egg Timer in Gio"))
		mainWindow.Option(app.Size(unit.Dp(MAIN_WIDTH), unit.Dp(MAIN_HEIGHT)))

		// center window on the desktop
		mainWindow.Perform(system.ActionCenter) // doesn't work?

		// listen for events in the window
		for {
			mainWindow.Event()
		}
	}()
	app.Main()
}
