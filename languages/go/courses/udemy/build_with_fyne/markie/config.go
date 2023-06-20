package main

import (
	"fyne.io/fyne/v2"
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
	saveAsMenuItem := fyne.NewMenuItem("Save As...", func() {})

	fileMenu := fyne.NewMenu("File", newMenuItem, fyne.NewMenuItemSeparator(), openMenuItem,
		fyne.NewMenuItemSeparator(), saveMenuItem, saveAsMenuItem)
	menu := fyne.NewMainMenu(fileMenu)

	window.SetMainMenu(menu)
}
