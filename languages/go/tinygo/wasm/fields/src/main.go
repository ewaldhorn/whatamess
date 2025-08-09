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
const (
	CANVAS_WIDTH      = 800
	CANVAS_HEIGHT     = 600
	CELL_SIZE         = 10
	ROWS              = CANVAS_HEIGHT / CELL_SIZE
	COLS              = CANVAS_WIDTH / CELL_SIZE
	RANDOM_COLOUR_IDX = 0
)

var canvasOne *tinycanvas.TinyCanvas

// ----------------------------------------------------------------------------
func main() {
	initColours()
	setCallbacks()
	setupKeyListeners()
	initUX()
	bootstrap()
	canvasOne = tinycanvas.NewTinyCanvas(CANVAS_WIDTH, CANVAS_HEIGHT)
	initEffects()

	performDemoOnCanvasOne()

	// prevent the app for closing - it stays running for the life of the webpage
	ch := make(chan struct{})
	<-ch
}

// ----------------------------------------------------------------------------
func initUX() {
	dom.Hide("loading")
	dom.Show("controls")
	dom.Show("information")
}

// ----------------------------------------------------------------------------
func setCallbacks() {
	setVersionCallback()
	setRefreshCallback()
}

// ----------------------------------------------------------------------------
func setupKeyListeners() {
	js.Global().Set("onkeydown", js.FuncOf(func(this js.Value, args []js.Value) any {
		keyPressed := args[0].Get("code").String()

		isCtrlPressed := args[0].Get("ctrlKey").Bool()
		isMetaPressed := args[0].Get("metaKey").Bool()

		switch keyPressed {
		case "KeyD":
			effect.toggleDebugging()
		case "KeyR":
			if !isCtrlPressed && !isMetaPressed {
				effect.randomise(true)
			}
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
		case "KeyZ":
			effect.switchToRandomColour()
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
