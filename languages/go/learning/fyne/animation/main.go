package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"someanimation/collection"
	"someanimation/ui"
)

func main() {
	a := app.New()
	w := a.NewWindow("Metaballs")
	model := collection.NewRandomCollection(8)
	widget := ui.NewAmoebaWidget(model)
	w.SetContent(widget)
	w.Resize(fyne.NewSize(512, 512))
	widget.Animate()
	w.ShowAndRun()
}
