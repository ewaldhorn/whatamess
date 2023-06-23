package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	myApp := app.New()

	myWindow := myApp.NewWindow("Simple App")
	myWindow.Resize(fyne.NewSize(800.0, 600.0))
	myWindow.CenterOnScreen()

	myWindow.ShowAndRun()
}
