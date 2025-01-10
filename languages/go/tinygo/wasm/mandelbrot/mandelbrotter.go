package main

import "github.com/ewaldhorn/tinycanvas/tinycanvas"

// ----------------------------------------------------------------------------
type Mandelbrotter struct {
	canvas *tinycanvas.TinyCanvas
}

// ----------------------------------------------------------------------------
func NewMandelbrotter() *Mandelbrotter {
	brotter := &Mandelbrotter{}

	brotter.setupCanvas()

	return brotter
}

// ----------------------------------------------------------------------------
func (m *Mandelbrotter) setupCanvas() {
	m.canvas = tinycanvas.NewTinyCanvas(800, 600)
}

// ----------------------------------------------------------------------------
func (m *Mandelbrotter) Start() {}
