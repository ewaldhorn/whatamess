package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

type DazyApp struct {
	app        fyne.App
	mainWindow fyne.Window

	entry    *widget.Entry
	savedURI fyne.URI

	cursorRow, cursorCol *widget.Label

	unsavedChanges binding.Bool
}

// ----------------------------------------------------------------------------
func CreateNewApp() *DazyApp {
	dazy := &DazyApp{}
	dazy.unsavedChanges = binding.NewBool()

	dazy.initFyneApp()
	dazy.initMainWindow()

	// visual indicator of unsaved changes
	dazy.unsavedChanges.AddListener(binding.NewDataListener(func() {
		edited, _ := dazy.unsavedChanges.Get()
		if edited {
			dazy.mainWindow.SetTitle(APP_TITLE + " *")
		} else {
			dazy.mainWindow.SetTitle(APP_TITLE)
		}
	}))

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
