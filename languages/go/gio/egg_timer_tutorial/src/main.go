package main

import (
	"os"

	"gioui.org/app"
	"gioui.org/io/system"
	"gioui.org/op"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
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

		// ops are the operations from the UI
		var ops op.Ops

		// startButton is a clickable widget
		var startButton widget.Clickable

		// th defines the material design style
		th := material.NewTheme()

		// listen for events in the window
		for {
			evt := mainWindow.Event()

			// then detect the type
			switch typ := evt.(type) {
			// this is sent when the application should re-render.
			case app.FrameEvent:
				windowContext := app.NewContext(&ops, typ)
				btn := material.Button(th, &startButton, "Start")
				btn.Layout(windowContext)
				typ.Frame(windowContext.Ops)
			// and this is sent when the application should exits
			case app.DestroyEvent:
				os.Exit(0)
			}
		}
	}()
	app.Main()
}
