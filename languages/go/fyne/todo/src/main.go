package main

import (
	"fmt"
	"strings"
	"todoapp/src/models"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
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

	todos := binding.NewUntypedList()
	for _, t := range data {
		todos.Append(t)
	}

	return widget.NewListWithData(
		todos,
		// func that returns the component structure of the List Item
		func() fyne.CanvasObject {
			return container.NewBorder(nil, nil, nil, widget.NewCheck("", func(b bool) {}), widget.NewLabel(""))
		},
		// func that is called for each item in the list and allows
		// you to show the content on the previously defined ui structure
		// func that is called for each item in the list and allows
		// but this time we get the actual DataItem we need to cast
		func(di binding.DataItem, object fyne.CanvasObject) {
			container, _ := object.(*fyne.Container)
			// ideally we should check `ok` for each one of those casting
			// but we know that they are those types for sure
			label := container.Objects[0].(*widget.Label)
			checkbox := container.Objects[1].(*widget.Check)

			diu, _ := di.(binding.Untyped).Get()
			todo := diu.(models.Todo)

			label.SetText(todo.Description)
			checkbox.SetChecked(todo.Done)
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
