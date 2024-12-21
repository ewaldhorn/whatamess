package main

import (
	"io"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
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
func (dazy *DazyApp) fileSave() {
	if dazy.savedURI != nil {
		writer, err := storage.Writer(dazy.savedURI)
		if err != nil {
			dialog.ShowError(err, dazy.mainWindow)
			return
		}

		_, err = writer.Write([]byte(dazy.entry.Text))
		if err != nil {
			dialog.ShowError(err, dazy.mainWindow)
		}
		_ = writer.Close()
	}
}

// ----------------------------------------------------------------------------
func (dazy *DazyApp) fileSaveAs() {
	dialog.ShowFileSave(func(writer fyne.URIWriteCloser, err error) {
		if err != nil {
			dialog.ShowError(err, dazy.mainWindow)
			return
		}

		if writer == nil {
			return
		}

		_, err = writer.Write([]byte(dazy.entry.Text))
		if err != nil {
			dialog.ShowError(err, dazy.mainWindow)
		}
		_ = writer.Close()
	}, dazy.mainWindow)
}
