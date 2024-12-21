package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

type DazyApp struct {
	app        fyne.App
	mainWindow fyne.Window

	entry    *widget.Entry
	savedURI fyne.URI

	cursorRow, cursorCol *widget.Label
}

// ----------------------------------------------------------------------------
func CreateNewApp() *DazyApp {
	dazy := &DazyApp{}

	dazy.initFyneApp()
	dazy.initMainWindow()

	return dazy
}

// ----------------------------------------------------------------------------
func (dazy *DazyApp) initFyneApp() {
	dazy.app = app.NewWithID("nofuss_crazy_dazy_text_editor")
}

// ----------------------------------------------------------------------------
func (dazy *DazyApp) initMainWindow() {
	dazy.mainWindow = dazy.createMainAppWindow()
}

// ----------------------------------------------------------------------------
func (dazy *DazyApp) Run() {
	dazy.mainWindow.ShowAndRun()
}
