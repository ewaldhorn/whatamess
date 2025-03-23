package main

import (
	"syscall/js"

	"github.com/ewaldhorn/dommie/dom"
	"github.com/ewaldhorn/tinycanvas/tinycanvas"
)

// ------------------------------------------ Externally provided by JavaScript

//export bootstrap
func bootstrap()

//export animateCanvasOne
func animateCanvasOne()

// -------------------------------------------------------------------- GLOBALS
var canvasOne *tinycanvas.TinyCanvas

// ----------------------------------------------------------------------------
func main() {
	setCallbacks()
	setupKeyListeners()
	dom.Hide("loading")
	dom.Show("controls")
	dom.Show("information")
	bootstrap()
	canvasOne = tinycanvas.NewTinyCanvas(800, 600)
	initEffects()

	performDemoOnCanvasOne()

	// prevent the app for closing - it stays running for the life of the webpage
	ch := make(chan struct{})
	<-ch
}

// ----------------------------------------------------------------------------
func setCallbacks() {
	setVersionCallback()
	setRefreshCallback()
}

// ----------------------------------------------------------------------------
func setupKeyListeners() {
	js.Global().Set("onkeydown", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		keyPressed := args[0].Get("code").String()

		switch keyPressed {
		case "KeyD":
			effect.toggleDebugging()
		}

		return nil
	}))
}

// ----------------------------------------------------------------------------
func setRefreshCallback() {
	dom.AddEventListener("refreshButton", "click", func() {
		if canvasOne != nil {
			effect.randomise()
		}
	})
}
