package main

import (
	"os"
	"testing"
)

// Setup for all tests
var app application

// ----------------------------------------------------------------------------
func TestMain(m *testing.M) {
	app.Session = getSession()
	os.Exit(m.Run())
}
