package main

import (
	"syscall/js"

	"github.com/ewaldhorn/tinycanvas/colour"
)

var x, y = 0, 0
var width, height int
var effect *Effect
var white = *colour.NewColourWhite()
var black = *colour.NewColourBlack()

// ----------------------------------------------------------------------------
func initEffects() {
	effect = NewEffect(canvasOne.Width(), canvasOne.Height())
}

// ----------------------------------------------------------------------------
func performDemoOnCanvasOne() {
	width, height = canvasOne.GetDimensions()
	canvasOne.ClearScreen(*colour.NewColour(80, 80, 180, 255))
	setRefreshCanvasOneCallback()
	animateCanvasOne()
}

// ----------------------------------------------------------------------------
func updateCanvasOne() {
	canvasOne.ClearScreen(black)

	effect.render()

	canvasOne.Render()
}

// ----------------------------------------------------------------------------
func setRefreshCanvasOneCallback() {
	js.Global().Set("refreshCanvasOne", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		updateCanvasOne()
		return nil
	}))
}
