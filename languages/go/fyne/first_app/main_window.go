package main

import (
	"fyne.io/fyne/v2"
)

// ----------------------------------------------------------------------------
func createMainAppWindow(app fyne.App) fyne.Window {
	newWindow := app.NewWindow("Crazy Dazy Text Editor")

	newWindow.SetMaster()
	newWindow.SetContent(createMainUI())

	newWindow.Resize(fyne.NewSize(640, 480))
	newWindow.CenterOnScreen()

	return newWindow
}
