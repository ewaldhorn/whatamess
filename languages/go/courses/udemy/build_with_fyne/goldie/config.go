package main

import (
	"goldie/repository"
	"log"
	"net/http"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type Config struct {
	App                 fyne.App
	InfoLog             *log.Logger
	ErrorLog            *log.Logger
	DB                  repository.Repository
	MainWindow          fyne.Window
	PriceContainer      *fyne.Container
	PriceChartContainer *fyne.Container
	Toolbar             *widget.Toolbar
	HTTPClient          *http.Client
}
