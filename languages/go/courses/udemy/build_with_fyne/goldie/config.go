package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"log"
	"net/http"
)

type Config struct {
	App            fyne.App
	InfoLog        *log.Logger
	ErrorLog       *log.Logger
	MainWindow     fyne.Window
	PriceContainer *fyne.Container
	Toolbar        *widget.Toolbar
	HTTPClient     *http.Client
}
