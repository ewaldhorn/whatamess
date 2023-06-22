package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
)

func main() {
	application := app.New()
	mainWindow := application.NewWindow("Transparency")

	mainWindow.Resize(fyne.Size{Width: 640.0, Height: 480.0})
	mainWindow.CenterOnScreen()

	img := canvas.NewImageFromFile("./some_fool_at_devconf.jpg")
	img.Translucency = 0.90

	mainWindow.SetContent(img)

	mainWindow.ShowAndRun()
}
