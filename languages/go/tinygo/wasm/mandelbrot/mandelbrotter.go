package main

import (
	"syscall/js"

	"github.com/ewaldhorn/tinycanvas/tinycanvas"
)

// ----------------------------------------------------------------------------
type Mandelbrotter struct {
	canvas                       *tinycanvas.TinyCanvas
	width, height, maxIterations int
	x, y, currentIter            int
	cx, cy                       float64
	done                         bool
}

// ----------------------------------------------------------------------------
// Mandelbrotter creates a new mandelbrot generator using the supplied size as width and height.
// It will set up a canvas of "size" pixels by "size" pixels ready for rendering.
// - size int : the required size of the drawing. Drawings are square.
func NewMandelbrotter(size int) *Mandelbrotter {
	brotter := &Mandelbrotter{width: size, height: size, maxIterations: 255}

	brotter.setup()

	return brotter
}

// ----------------------------------------------------------------------------
func (m *Mandelbrotter) setup() {
	m.setupCanvas()
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
func (m *Mandelbrotter) step() {
	if m.y >= m.height {
		// we are done
		m.done = true
		return
	}

	if m.x >= m.width {
		m.x = 0
		m.y += 1
	}

	m.cx = float64(m.x-m.width/2) * 4.0 / float64(m.width)
	m.cy = float64(m.y-m.height/2) * 4.0 / float64(m.height)
	m.currentIter = mandelbrot(m.cx, m.cy, m.maxIterations)
	colour := getColour(m.currentIter, m.maxIterations)
	m.canvas.ColourPutPixel(m.x, m.y, *colour)

	m.x++
}

// ----------------------------------------------------------------------------
func (m *Mandelbrotter) Refresh() {
	if !m.done {
		for range m.width {
			m.step()
		}

		m.canvas.Render()
	}
}
