package main

import (
	"database/sql"
	"goldie/repository"
	"log"
	"net/http"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"

	_ "github.com/glebarez/go-sqlite"
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
	sqlDB, err := myApp.connectSQL()
	if err != nil {
		log.Panic(err)
	}

	// create database repository
	myApp.setupDB(sqlDB)

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

func (myApp *Config) connectSQL() (*sql.DB, error) {
	path := ""
	if os.Getenv("DB_PATH") != "" {
		path = os.Getenv("DB_PATH")
	} else {
		path = myApp.App.Storage().RootURI().Path() + "/sql.db"
		myApp.InfoLog.Println("db in:", path)
	}

	db, err := sql.Open("sqlite", path)
	if err != nil {
		myApp.ErrorLog.Println("Error opening database:", err)
		return nil, err
	}

	return db, nil
}

func (myApp *Config) setupDB(sqlDB *sql.DB) {
	myApp.DB = repository.NewSQLiteRepository(sqlDB)

	err := myApp.DB.Migrate()
	if err != nil {
		myApp.ErrorLog.Println(err)
		log.Panic()
	}
}
