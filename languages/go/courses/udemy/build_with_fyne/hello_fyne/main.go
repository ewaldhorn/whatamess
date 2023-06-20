package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	application := app.New()

	mainWindow := application.NewWindow(APP_MAIN_WINDOW_TITLE)
	mainWindow.Resize(fyne.NewSize(640.0, 480.0))
	mainWindow.CenterOnScreen()
	mainWindow.SetContent(widget.NewLabel("Hello, world."))

	mainWindow.ShowAndRun()
}
