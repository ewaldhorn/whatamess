package main

import (
	"syscall/js"

	"github.com/ewaldhorn/tinycanvas/colour"
	"github.com/ewaldhorn/tinycanvas/tinycanvas"
)

// ----------------------------------------------------------------------------
type Mandelbrotter struct {
	canvas *tinycanvas.TinyCanvas
}

// ----------------------------------------------------------------------------
func NewMandelbrotter() *Mandelbrotter {
	brotter := &Mandelbrotter{}

	brotter.setup()

	return brotter
}

// ----------------------------------------------------------------------------
func (m *Mandelbrotter) setup() {
	m.setupCanvas()
	width, height := m.canvas.GetDimensions()

	for x := range width {
		for y := range height {
			m.canvas.PutColourPixel(x, y, *colour.NewRandomColour())
		}
	}

	m.setupRenderFrameCallback()
}

// ----------------------------------------------------------------------------
func (m *Mandelbrotter) setupCanvas() {
	m.canvas = tinycanvas.NewTinyCanvas(800, 600)
}

// ----------------------------------------------------------------------------
func (m *Mandelbrotter) setupRenderFrameCallback() {
	js.Global().Set("renderFrame", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		m.Refresh()
		return nil
	}))
}

// ----------------------------------------------------------------------------
func (m *Mandelbrotter) Refresh() {
	m.canvas.Render()
}
