package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

// ----------------------------------------------------------------------------
func createMainUI() fyne.CanvasObject {
	return widget.NewLabel("Hello, Fyne!")
}
