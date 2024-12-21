package main

import (
	"fyne.io/fyne/v2"
)

const (
	APP_TITLE = "Crazy Dazy Text Editor"
)

// ----------------------------------------------------------------------------
func (dazy *DazyApp) createMainAppWindow() fyne.Window {
	newWindow := dazy.app.NewWindow(APP_TITLE)

	newWindow.SetMaster()
	newWindow.SetContent(dazy.createMainUI())

	newWindow.Resize(fyne.NewSize(640, 480))
	newWindow.CenterOnScreen()

	return newWindow
}
