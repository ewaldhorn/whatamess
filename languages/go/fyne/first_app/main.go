package main

import (
	"fyne.io/fyne/v2/app"
)

// ----------------------------------------------------------------------------
func main() {
	mainApp := app.New()
	mainWindow := createMainAppWindow(mainApp)
	mainWindow.ShowAndRun()
}
