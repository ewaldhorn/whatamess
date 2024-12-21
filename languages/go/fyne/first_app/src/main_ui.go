package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// ----------------------------------------------------------------------------
func (dazy *DazyApp) createMainUI() fyne.CanvasObject {
	entry := widget.NewMultiLineEntry()

	cursorRow := widget.NewLabel("1")
	cursorCol := widget.NewLabel("1")

	toolbar := dazy.buildMainToolbar()

	status := container.NewHBox(layout.NewSpacer(), widget.NewLabel("Cursor Row:"), cursorRow, widget.NewLabel("Col:"), cursorCol)

	return container.NewBorder(toolbar, status, nil, nil, entry)
}

// ----------------------------------------------------------------------------
func (dazy *DazyApp) buildMainToolbar() *widget.Toolbar {
	return widget.NewToolbar(
		widget.NewToolbarAction(theme.FolderOpenIcon(), dazy.fileOpen),
		widget.NewToolbarAction(theme.DocumentSaveIcon(), fileSave),
		widget.NewToolbarSeparator(),
		widget.NewToolbarAction(theme.ContentCutIcon(), textCut),
		widget.NewToolbarAction(theme.ContentCopyIcon(), textCopy),
		widget.NewToolbarAction(theme.ContentPasteIcon(), textPaste),
	)
}
