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
const CANVAS_WIDTH = 800
const CANVAS_HEIGHT = 600
const CELL_SIZE = 10
const ROWS = CANVAS_HEIGHT / CELL_SIZE
const COLS = CANVAS_WIDTH / CELL_SIZE
const RANDOM_COLOUR_IDX = 0

var canvasOne *tinycanvas.TinyCanvas

// ----------------------------------------------------------------------------
func main() {
	InitColours()
	setCallbacks()
	setupKeyListeners()
	dom.Hide("loading")
	dom.Show("controls")
	dom.Show("information")
	bootstrap()
	canvasOne = tinycanvas.NewTinyCanvas(CANVAS_WIDTH, CANVAS_HEIGHT)
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
		case "KeyR":
			effect.randomise(true)
		case "KeyC":
			effect.colourChange()
		case "KeyP":
			effect.patternChange()
		case "KeyM":
			if effect.colourRange == RANDOM_COLOUR_IDX {
				RandomiseRandomColour()
			}
		case "KeyS":
			if effect.colourRange == RANDOM_COLOUR_IDX {
				RandomColourShade()
			}
		}

		return nil
	}))
}

// ----------------------------------------------------------------------------
func setRefreshCallback() {
	dom.AddEventListener("refreshButton", "click", func() {
		if canvasOne != nil {
			effect.randomise(true)
		}
	})
}
