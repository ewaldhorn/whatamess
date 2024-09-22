package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// ------------------------------------------------------------- Screen Globals
// Yes yes, it should be in a struct, but this is a single page app...
var inputWidget *widget.Entry

// ----------------------------------------------------------------------------
func createApplicationWindow(app fyne.App, title string) fyne.Window {
	w := app.NewWindow(title)

	w.Resize(fyne.NewSize(640, 480))
	w.CenterOnScreen()

	return w
}

// ----------------------------------------------------------------------------
func createInputSection() fyne.CanvasObject {
	inputWidget = widget.NewEntry()
	inputWidget.PlaceHolder = "Add new todo..."

	return container.NewGridWithColumns(2,
		inputWidget,
		widget.NewButton("Add", func() { fmt.Println("Add was clicked!") }),
	)
}

// ----------------------------------------------------------------------------
func createMainScreen() fyne.CanvasObject {
	return container.NewBorder(
		nil, // TOP of the container

		// this will be a the BOTTOM of the container
		createInputSection(),
		nil, // Right
		nil, // Left

		// the rest will take all the rest of the space
		container.NewCenter(
			widget.NewLabel("Content will end up here"),
		),
	)
}

// ----------------------------------------------------------------------------
func main() {
	a := app.New()
	w := createApplicationWindow(a, "Todo App")

	// t := models.NewTodo("Show this on the window")

	w.SetContent(
		createMainScreen(),
	)

	w.ShowAndRun()
}
