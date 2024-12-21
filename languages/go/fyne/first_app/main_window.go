package main

import (
	"fyne.io/fyne/v2"
)

// ----------------------------------------------------------------------------
func (dazy *DazyApp) createMainAppWindow() fyne.Window {
	newWindow := dazy.app.NewWindow("Crazy Dazy Text Editor")

	newWindow.SetMaster()
	newWindow.SetContent(dazy.createMainUI())

	newWindow.Resize(fyne.NewSize(640, 480))
	newWindow.CenterOnScreen()

	return newWindow
}
