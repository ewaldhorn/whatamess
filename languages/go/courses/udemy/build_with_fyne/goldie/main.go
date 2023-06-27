package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"log"
	"net/http"
	"os"
)

const AppName = "za.co.nofuss.goldie"
const AppTitle = "Gold Tracker"

var myApp Config

func main() {
	// create app
	fyneApp := app.NewWithID(AppName)
	myApp.App = fyneApp
	myApp.HTTPClient = &http.Client{}

	// create loggers
	setupLoggers()

	// open connection to database
	// create database repository

	// create and size main window
	myApp.MainWindow = configureMainWindow()

	// now create the window content
	myApp.makeUI()

	// show and run app
	myApp.MainWindow.ShowAndRun()
}

func setupLoggers() {
	myApp.InfoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	myApp.ErrorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
}

func configureMainWindow() fyne.Window {
	mainWindow := myApp.App.NewWindow(AppTitle)

	mainWindow.Resize(fyne.Size{Width: 800.0, Height: 600.0})
	mainWindow.CenterOnScreen()
	mainWindow.SetFixedSize(true)
	mainWindow.SetMaster()

	return mainWindow
}
