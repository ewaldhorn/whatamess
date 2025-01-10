package main

import "github.com/ewaldhorn/tinycanvas/dom"

// ----------------------------------------------------------------------------
// The following functions are externally provided by the WASM runtime environment.

//export setVersion
func setVersion()

// ----------------------------------------------------------------------------
// Standard entry point for Go programs
func main() {
	// general bootstrapping
	setupCallbacks()
	setVersion()
	dom.Hide("loading")

	// now we should be ready to rock and roll
	mandelbrotter := NewMandelbrotter()
	mandelbrotter.Refresh()

	// prevent the app for closing - it stays running for the life of the page
	// if we don't do this, it exits and all future calls into it will fail
	ch := make(chan struct{})
	<-ch
}

// ----------------------------------------------------------------------------
// Set up the various callbacks used by the JavaScript side of things.
func setupCallbacks() {
	setVersionCallback()
}
