package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

const AppName = "za.co.nofuss.goldie"

var myApp Config

func main() {
	// create app
	fyneApp := app.NewWithID(AppName)
	myApp.App = fyneApp

	// create loggers

	// open connection to database
	// create database repository
	// create and size main window
	mainWindow := myApp.App.NewWindow("Gold Tracker")
	mainWindow.Resize(fyne.Size{Width: 800.0, Height: 600.0})
	mainWindow.CenterOnScreen()

	// show and run app
	mainWindow.ShowAndRun()
}
