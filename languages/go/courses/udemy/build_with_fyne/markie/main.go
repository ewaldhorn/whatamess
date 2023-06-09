package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
)

var cfg config

func main() {
	myApp := app.New()

	// use a custom theme
	myApp.Settings().SetTheme(&myTheme{})

	mainWindow := myApp.NewWindow(APP_TITLE)
	mainWindow.Resize(fyne.NewSize(800.0, 600.0))
	mainWindow.CenterOnScreen()

	edit, preview := cfg.createMainUI()
	cfg.createMainMenu(mainWindow)

	mainWindow.SetContent(container.NewHSplit(edit, preview))

	mainWindow.ShowAndRun()
}
