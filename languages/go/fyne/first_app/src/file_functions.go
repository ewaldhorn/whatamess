package main

import (
	"fmt"
	"io"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
)

// ----------------------------------------------------------------------------
func (dazy *DazyApp) fileOpen() {
	dialog.ShowFileOpen(func(reader fyne.URIReadCloser, err error) {
		if err != nil {
			dialog.ShowError(err, dazy.mainWindow)
			return
		}
		if reader == nil {
			return
		}

		data, err := io.ReadAll(reader)
		if err == nil {
			_ = reader.Close()
		} else {
			dialog.ShowError(err, dazy.mainWindow)
			return
		}
		_ = data
		fmt.Println("File:", reader.URI().String())
	}, dazy.mainWindow)
}

// ----------------------------------------------------------------------------
func fileSave() {}

// ----------------------------------------------------------------------------
func fileSaveAs() {}
