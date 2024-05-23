package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	a := app.New()
	w := a.NewWindow("Bouncies")
	model := newRandomEnsemble(7)
	widget := newMetaballsWidget(model)
	w.SetContent(widget)
	w.Resize(fyne.NewSize(640, 480))
	w.CenterOnScreen()
	widget.animate()
	w.ShowAndRun()
}
