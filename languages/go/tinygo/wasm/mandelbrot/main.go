package main

import "github.com/ewaldhorn/tinycanvas/dom"

//export setVersion
func setVersion()

// ----------------------------------------------------------------------------
func main() {
	setupCallbacks()
	setVersion()
	dom.Hide("loading")
}

// ----------------------------------------------------------------------------
func setupCallbacks() {
	setVersionCallback()
}
