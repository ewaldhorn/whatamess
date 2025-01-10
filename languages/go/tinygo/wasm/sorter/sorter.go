package main

import (
	"math/rand"
	"syscall/js"

	"github.com/ewaldhorn/tinycanvas/colour"
	"github.com/ewaldhorn/tinycanvas/tinycanvas"
)

// ----------------------------------------------------------------------------
type Sorter struct {
	canvas             *tinycanvas.TinyCanvas
	width, height      int
	barWidth, gapWidth int
	bottomOffset       int
	done               bool
	count              int
	numbers            []int8
}

// ----------------------------------------------------------------------------
func NewSorter(width, height int) *Sorter {
	sorter := &Sorter{width: width, height: height, bottomOffset: 2, count: 50}
	sorter.setup()
	return sorter
}

// ----------------------------------------------------------------------------
func (s *Sorter) setup() {
	s.setupRandomNumbers()
	s.setupCanvas()
	s.setupRenderFrameCallback()
}

// ----------------------------------------------------------------------------
func (s *Sorter) setupRandomNumbers() {
	for range s.count {
		s.numbers = append(s.numbers, int8(1+rand.Intn(100)))
	}

	// calculate the width of each bar, minus the margins
	// and the gap between bars
	s.gapWidth = 2
	s.barWidth = (s.width - ((s.count - 1) * s.gapWidth)) / s.count
	// for debugging
	// dom.Log(fmt.Sprintf("bar: %d, gap: %d", s.barWidth, s.gapWidth))
}

// ----------------------------------------------------------------------------
func (s *Sorter) setupCanvas() {
	s.canvas = tinycanvas.NewTinyCanvas(s.width, s.height)
	s.canvas.ClearScreen(*colour.NewColour(50, 50, 200, 255))
}

// ----------------------------------------------------------------------------
func (s *Sorter) setupRenderFrameCallback() {
	js.Global().Set("renderFrame", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		s.Refresh()
		return nil
	}))
}

// ----------------------------------------------------------------------------
func (s *Sorter) step() {
	s.done = true
	barColour := colour.NewColour(240, 235, 10, 255)
	for idx := range s.count {
		xpos := (idx * s.barWidth) + (idx * s.gapWidth) + 1
		for th := range s.barWidth {
			s.canvas.ColourLine(xpos+th, s.height-s.bottomOffset, xpos+th, (s.height-s.bottomOffset)-int(s.numbers[idx]), *barColour)
		}
	}
}

// ----------------------------------------------------------------------------
func (s *Sorter) Refresh() {
	if !s.done {
		s.step()
		s.canvas.Render()
	}
}
