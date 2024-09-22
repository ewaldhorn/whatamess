package main

import (
	"todoapp/src/models"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("TODO App")
	w.Resize(fyne.NewSize(640, 480))
	w.CenterOnScreen()

	t := models.NewTodo("Show this on the window")

    w.SetContent(widget.NewLabel(t.String()))

	w.ShowAndRun()
}
