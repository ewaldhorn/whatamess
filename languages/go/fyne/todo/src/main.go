package main

import (
	"fmt"
	"todoapp/src/models"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// ----------------------------------------------------------------------------
func createApplicationWindow(app fyne.App, title string) fyne.Window {
	w := app.NewWindow(title)

	w.Resize(fyne.NewSize(640, 480))
	w.CenterOnScreen()

	return w
}

// ----------------------------------------------------------------------------
func main() {
	a := app.New()
	w := createApplicationWindow(a, "Todo App")

	t := models.NewTodo("Show this on the window")

	w.SetContent(
		container.NewBorder(
			nil, // TOP of the container

			// this will be a the BOTTOM of the container
			container.NewGridWithColumns(2,
				widget.NewEntry(),
				widget.NewButton("Add", func() { fmt.Println("Add was clicked!") }),
			),

			nil, // Right
			nil, // Left

			// the rest will take all the rest of the space
			container.NewCenter(
				widget.NewLabel(t.String()),
			),
		),
	)
	w.ShowAndRun()
}
