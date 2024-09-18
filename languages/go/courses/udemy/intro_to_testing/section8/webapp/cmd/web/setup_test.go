package main

import (
	"log"
	"os"
	"testing"
	"webapp/pkg/db"
)

// Setup for all tests
var app application

// ----------------------------------------------------------------------------
func TestMain(m *testing.M) {
	// set up correct default test paths
	pathToTemplates = "./../../templates/"
	pathToPages = pathToTemplates + "pages"

	conStr := "host=localhost port=5432 user=postgres password=postgres dbname=users sslmode=disable timezone=UCT+2 connect_timeout=5"
	app.DSN = conStr
	dbConnection, err := app.connectToDatabase()
	if err != nil {
		log.Fatal("TEST: Database connection failed:", err)
	}
	defer dbConnection.Close()

	app.DB = db.PostgresConn{DB: dbConnection}

	app.Session = getSession()
	os.Exit(m.Run())
}
