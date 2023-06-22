package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	myApp := app.New()

	mainWindow := myApp.NewWindow("Gold Tracker")

	mainWindow.Resize(fyne.Size{Width: 800.0, Height: 600.0})
	mainWindow.CenterOnScreen()

	mainWindow.ShowAndRun()
}
