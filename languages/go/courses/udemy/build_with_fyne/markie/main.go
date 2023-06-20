package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
)

var cfg config

func main() {
	myApp := app.New()

	mainWindow := myApp.NewWindow(APP_TITLE)
	mainWindow.Resize(fyne.NewSize(800.0, 600.0))
	mainWindow.CenterOnScreen()

	edit, preview := cfg.createMainUI()

	mainWindow.SetContent(container.NewHSplit(edit, preview))

	mainWindow.ShowAndRun()
}
