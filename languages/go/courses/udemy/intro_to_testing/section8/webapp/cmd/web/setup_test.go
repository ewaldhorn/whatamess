package main

import (
	"os"
	"testing"
)

// Setup for all tests
var app application

// ----------------------------------------------------------------------------
func TestMain(m *testing.M) {
	// set up correct default test paths
	pathToTemplates = "./../../templates/"
	pathToPages = pathToTemplates + "pages"

	app.Session = getSession()
	os.Exit(m.Run())
}
