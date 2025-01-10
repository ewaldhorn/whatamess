package main

import "github.com/ewaldhorn/tinycanvas/dom"

// ----------------------------------------------------------------------------
// The following functions are externally provided by the WASM runtime environment.

//export setVersion
func setVersion()

//export startAnimation
func startAnimation()

// ----------------------------------------------------------------------------
// Standard entry point for Go programs
func main() {
	// general bootstrapping
	setupCallbacks()
	setVersion()
	dom.Hide("loading")
	dom.Show("controls")

	// now we should be ready to rock and roll
	sorter := NewSorter(800, 200)
	sorter.Refresh()

	startAnimation()
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
