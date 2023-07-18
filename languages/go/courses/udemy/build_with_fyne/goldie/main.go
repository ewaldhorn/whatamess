package main

import (
	"log"
	"net/http"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

const AppName = "za.co.nofuss.goldie"
const AppTitle = "Gold Tracker"

func main() {
	// create app
	var myApp Config
	myApp.App = app.NewWithID(AppName)
	myApp.HTTPClient = &http.Client{}

	// create loggers
	setupLoggers(&myApp)

	// open connection to database
	// create database repository

	// create and size main window
	myApp.MainWindow = configureMainWindow(&myApp)

	// now create the window content
	myApp.makeUI(myApp)

	// show and run app
	myApp.MainWindow.ShowAndRun()
}

func setupLoggers(myApp *Config) {
	myApp.InfoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	myApp.ErrorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
}

func configureMainWindow(myApp *Config) fyne.Window {
	mainWindow := myApp.App.NewWindow(AppTitle)

	mainWindow.Resize(fyne.Size{Width: 800.0, Height: 600.0})
	mainWindow.CenterOnScreen()
	mainWindow.SetFixedSize(true)
	mainWindow.SetMaster()

	return mainWindow
}
