package main

import (
	"os"
	"testing"
)

// Setup for all tests
var app application

// ----------------------------------------------------------------------------
func TestMain(m *testing.M) {
	os.Exit(m.Run())
}
