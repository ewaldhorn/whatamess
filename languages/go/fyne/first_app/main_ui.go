package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// ----------------------------------------------------------------------------
func createMainUI() fyne.CanvasObject {
	entry := widget.NewMultiLineEntry()
	cursorRow := widget.NewLabel("1")
	cursorCol := widget.NewLabel("1")

	toolbar := buildMainToolbar()

	status := container.NewHBox(layout.NewSpacer(), widget.NewLabel("Cursor Row:"), cursorRow, widget.NewLabel("Col:"), cursorCol)

	return container.NewBorder(toolbar, status, nil, nil, container.NewScroll(entry))
}

// ----------------------------------------------------------------------------
func buildMainToolbar() *widget.Toolbar {
	return widget.NewToolbar(
		widget.NewToolbarAction(theme.FolderOpenIcon(), open),
		widget.NewToolbarAction(theme.DocumentSaveIcon(), save),
		widget.NewToolbarSeparator(),
		widget.NewToolbarAction(theme.ContentCutIcon(), cut),
		widget.NewToolbarAction(theme.ContentCopyIcon(), copy),
		widget.NewToolbarAction(theme.ContentPasteIcon(), paste),
	)
}
