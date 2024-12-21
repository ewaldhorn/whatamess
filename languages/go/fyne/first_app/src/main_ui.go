package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// ----------------------------------------------------------------------------
func (dazy *DazyApp) updateCursorStatus() {
	dazy.cursorRow.SetText(fmt.Sprintf("%d", dazy.entry.CursorRow+1))
	dazy.cursorCol.SetText(fmt.Sprintf("%d", dazy.entry.CursorColumn+1))
}

// ----------------------------------------------------------------------------
func (dazy *DazyApp) createMainUI() fyne.CanvasObject {
	dazy.entry = widget.NewMultiLineEntry()
	dazy.entry.OnCursorChanged = dazy.updateCursorStatus
	dazy.entry.OnChanged = func(s string) {
		dazy.unsavedChanges.Set(true)
	}

	dazy.cursorRow = widget.NewLabel("1")
	dazy.cursorCol = widget.NewLabel("1")

	toolbar := dazy.buildMainToolbar()

	status := container.NewHBox(layout.NewSpacer(), widget.NewLabel("Cursor Row:"), dazy.cursorRow, widget.NewLabel("Col:"), dazy.cursorCol)

	return container.NewBorder(toolbar, status, nil, nil, dazy.entry)
}

// ----------------------------------------------------------------------------
func (dazy *DazyApp) buildMainToolbar() *widget.Toolbar {
	return widget.NewToolbar(
		widget.NewToolbarAction(theme.FolderOpenIcon(), dazy.fileOpen),
		widget.NewToolbarAction(theme.DocumentSaveIcon(), dazy.fileSave),
		widget.NewToolbarAction(theme.DocumentCreateIcon(), dazy.fileSaveAs),
		widget.NewToolbarSeparator(),
		widget.NewToolbarAction(theme.ContentCutIcon(), dazy.textCut),
		widget.NewToolbarAction(theme.ContentCopyIcon(), dazy.textCopy),
		widget.NewToolbarAction(theme.ContentPasteIcon(), dazy.textPaste),
	)
}
