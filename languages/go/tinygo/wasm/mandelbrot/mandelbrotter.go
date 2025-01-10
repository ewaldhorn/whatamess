package main

import (
	"syscall/js"

	"github.com/ewaldhorn/tinycanvas/colour"
	"github.com/ewaldhorn/tinycanvas/tinycanvas"
)

// ----------------------------------------------------------------------------
type Mandelbrotter struct {
	canvas        *tinycanvas.TinyCanvas
	width, height int
}

// ----------------------------------------------------------------------------
// Mandelbrotter creates a new mandelbrot generator using the supplied size as width and height.
// It will set up a canvas of "size" pixels by "size" pixels ready for rendering.
// - size int : the required size of the drawing. Drawings are square.
func NewMandelbrotter(size int) *Mandelbrotter {
	brotter := &Mandelbrotter{width: size, height: size}

	brotter.setup()

	return brotter
}

// ----------------------------------------------------------------------------
func (m *Mandelbrotter) setup() {
	m.setupCanvas()

	for x := range m.width {
		for y := range m.height {
			m.canvas.PutColourPixel(x, y, *colour.NewRandomColour())
		}
	}

	m.setupRenderFrameCallback()
}

// ----------------------------------------------------------------------------
func (m *Mandelbrotter) setupCanvas() {
	m.canvas = tinycanvas.NewTinyCanvas(m.width, m.height)
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
