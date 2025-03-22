package main

import (
	"syscall/js"

	"github.com/ewaldhorn/tinycanvas/colour"
	"github.com/ewaldhorn/tinycanvas/tinycanvas"
)

var x, y = 0, 0
var width, height int

var white = *colour.NewColourWhite()

// ----------------------------------------------------------------------------
func initEffects() {
	effect := NewEffect(width, height)
	effect.init()
}

// ----------------------------------------------------------------------------
func performDemoOnCanvasOne() {
	width, height = canvasOne.GetDimensions()
	canvasOne.ClearScreen(*colour.NewColour(80, 80, 180, 255))
	setRefreshCanvasOneCallback()
	animateCanvasOne()
}

// ----------------------------------------------------------------------------
func renderTriangle() {
	canvasOne.SetColour(white)

	for i := 0; i < 40; i += 2 {
		canvasOne.Triangle(
			tinycanvas.Point{X: (width / 2), Y: (height / 3) + i},
			tinycanvas.Point{X: (width - (width / 3)) + i, Y: (height - (height / 3)) - i},
			tinycanvas.Point{X: (width / 3) - i, Y: (height - (height / 3)) - i},
		)
	}
}

// ----------------------------------------------------------------------------
func updateCanvasOne() {
	canvasOne.SetColour(*colour.NewRandomColour())

	renderTriangle()
	canvasOne.Render()
}

// ----------------------------------------------------------------------------
func setRefreshCanvasOneCallback() {
	js.Global().Set("refreshCanvasOne", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		updateCanvasOne()
		return nil
	}))
}
