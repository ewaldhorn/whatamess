package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type App struct {
	output *widget.Label
}

var myApp App

func main() {
	application := app.New()

	mainWindow := application.NewWindow(APP_MAIN_WINDOW_TITLE)
	mainWindow.Resize(fyne.NewSize(640.0, 480.0))
	mainWindow.CenterOnScreen()

	output, entry, btn := myApp.makeUI()

	mainWindow.SetContent(container.NewVBox(output, entry, btn))

	mainWindow.ShowAndRun()
}

func (app *App) makeUI() (*widget.Label, *widget.Entry, *widget.Button) {
	output := widget.NewLabel("Hello, world.")
	entry := widget.NewEntry()
	entry.PlaceHolder = "Type Something"
	btn := widget.NewButton("Update", func() { app.output.SetText(entry.Text) })

	app.output = output

	return output, entry, btn
}
