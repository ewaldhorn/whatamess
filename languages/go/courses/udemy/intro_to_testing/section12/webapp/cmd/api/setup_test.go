package main

import (
	"os"
	"testing"
	"webapp/pkg/repository/databaserepo"
)

// ----------------------------------------------------------------------------
var app application

// ----------------------------------------------------------------------------
func TestMain(m *testing.M) {
	app.DB = &databaserepo.TestingDBRepo{}
	app.Domain = "example.com"
	app.JWTSecret = "ee334ab190919293949a9b9cc1c2c3c4c5c6c7c8c910"

	os.Exit(m.Run())
}
