package main

import (
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
			dazy.entry.SetText(string(data))
			dazy.savedURI = reader.URI()
		} else {
			dialog.ShowError(err, dazy.mainWindow)
			return
		}
	}, dazy.mainWindow)
}

// ----------------------------------------------------------------------------
func fileSave() {}

// ----------------------------------------------------------------------------
func fileSaveAs() {}
