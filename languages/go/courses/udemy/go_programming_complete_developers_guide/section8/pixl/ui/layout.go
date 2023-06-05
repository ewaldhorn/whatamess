package ui

import (
	"fyne.io/fyne/v2/container"
)

func Setup(app *AppInit) {
	SetupMenus(app)
	swatchesContainer := BuildSwatches(app)
	colourPicker := SetupColourPicker(app)

	appLayout := container.NewBorder(nil, swatchesContainer, nil, colourPicker, app.PixlCanvas)

	app.PixlWindow.SetContent(appLayout)
}
