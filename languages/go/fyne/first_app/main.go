package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Crazy Dazy Text Editor")

	w.SetContent(widget.NewLabel("Hello, Fyne!"))

	w.Resize(fyne.NewSize(640, 480))
	w.CenterOnScreen()
	w.ShowAndRun()
}
