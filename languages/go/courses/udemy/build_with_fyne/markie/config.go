package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

type config struct {
	EditWidget    *widget.Entry
	PreviewWidget *widget.RichText
	CurrentFile   fyne.URI
	SaveMenuItem  *fyne.MenuItem
}

func (c *config) createMainUI() (*widget.Entry, *widget.RichText) {
	edit := widget.NewMultiLineEntry()
	preview := widget.NewRichTextFromMarkdown("")

	c.EditWidget = edit
	c.PreviewWidget = preview

	edit.OnChanged = preview.ParseMarkdown

	return edit, preview
}

func (c *config) createMainMenu(window fyne.Window) {
	newMenuItem := fyne.NewMenuItem("New", func() {})
	openMenuItem := fyne.NewMenuItem("Open...", func() {})
	saveMenuItem := fyne.NewMenuItem("Save", func() {})
	saveAsMenuItem := fyne.NewMenuItem("Save As...", c.saveAs(window))

	fileMenu := fyne.NewMenu("File", newMenuItem, fyne.NewMenuItemSeparator(), openMenuItem,
		fyne.NewMenuItemSeparator(), saveMenuItem, saveAsMenuItem)
	menu := fyne.NewMainMenu(fileMenu)

	// disable save menu item and set Save As menu item pointer
	saveMenuItem.Disabled = true
	c.SaveMenuItem = saveAsMenuItem

	window.SetMainMenu(menu)
}

func (c *config) saveAs(window fyne.Window) func() {
	return func() {
		saveDialog := dialog.NewFileSave(func(writer fyne.URIWriteCloser, err error) {
			if err != nil {
				// oops
				dialog.ShowError(err, window)
				return
			}

			if writer == nil {
				// user cancelled
				return
			}

			// all good, try to save the file
			writer.Write([]byte(c.EditWidget.Text))
			c.CurrentFile = writer.URI()

			defer writer.Close()

			window.SetTitle(APP_TITLE + " - " + writer.URI().Name())
			c.SaveMenuItem.Disabled = false
		}, window)

		saveDialog.Show()
	}
}
