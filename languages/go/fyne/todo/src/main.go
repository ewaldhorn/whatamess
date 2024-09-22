package main

import (
	"fmt"
	"strings"
	"todoapp/src/models"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// ------------------------------------------------------------- Screen Globals
// Yes yes, it should be in a struct, but this is a single page app...
var inputWidget *widget.Entry
var addButton *widget.Button

// ----------------------------------------------------------------------------
func createApplicationWindow(app fyne.App, title string) fyne.Window {
	newWindow := app.NewWindow(title)

	newWindow.Resize(fyne.NewSize(640, 480))
	newWindow.CenterOnScreen()

	return newWindow
}

// ----------------------------------------------------------------------------
func createInputSection() fyne.CanvasObject {
	addButton = widget.NewButton("Add", func() { fmt.Println("Addind") })
	addButton.Disable()

	inputWidget = widget.NewEntry()
	inputWidget.PlaceHolder = "Add new todo..."
	inputWidget.OnChanged = func(txt string) {
		addButton.Disable()
		if len(strings.TrimSpace(txt)) > 2 {
			addButton.Enable()
		}
	}

	return container.NewGridWithColumns(2,
		inputWidget,
		addButton,
	)
}

// ----------------------------------------------------------------------------
func createTodoListSection() fyne.CanvasObject {
	data := []models.Todo{
		models.NewTodo("Some stuff"),
		models.NewTodo("Some more stuff"),
		models.NewTodo("Some other things"),
	}

	return widget.NewList(
		// func that returns the number of items in the list
		func() int {
			return len(data)
		},
		// func that returns the component structure of the List Item
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		// func that is called for each item in the list and allows
		// you to show the content on the previously defined ui structure
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(data[i].Description)
		})
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
		createTodoListSection(),
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
