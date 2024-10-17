package main

import (
	"os"
	"testing"
	"webapp/pkg/repository/databaserepo"
)

// Setup for all tests
var app application

// ----------------------------------------------------------------------------
func TestMain(m *testing.M) {
	// set up correct default test paths
	pathToTemplates = "./../../templates/"
	pathToPages = pathToTemplates + "pages"

	app.DB = &databaserepo.TestingDBRepo{}

	app.Session = getSession()
	os.Exit(m.Run())
}
