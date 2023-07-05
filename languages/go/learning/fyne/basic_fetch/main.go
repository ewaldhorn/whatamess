package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	application := app.New()

	window := application.NewWindow("Fyne Fetch")
	window.Resize(fyne.NewSize(800.0, 600.0))
	window.CenterOnScreen()

	window.ShowAndRun()
}
