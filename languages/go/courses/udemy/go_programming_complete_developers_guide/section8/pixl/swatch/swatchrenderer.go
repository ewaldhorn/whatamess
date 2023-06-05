package swatch

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

type SwatchRenderer struct {
	square  canvas.Rectangle
	objects []fyne.CanvasObject
	parent  *Swatch
}

func (sr *SwatchRenderer) MinSize() fyne.Size {
	return sr.square.MinSize()
}

func (sr *SwatchRenderer) Layout(size fyne.Size) {
	sr.objects[0].Resize(size)
}

func (sr *SwatchRenderer) Refresh() {
	sr.Layout(fyne.NewSize(20, 20))
	sr.square.FillColor = sr.parent.Color
	if sr.parent.Selected {
		sr.square.StrokeWidth = 3
		sr.square.StrokeColor = color.NRGBA{255, 255, 255, 255}
		sr.objects[0] = &sr.square
	} else {
		sr.square.StrokeWidth = 0
		sr.objects[0] = &sr.square
	}
	canvas.Refresh(sr.parent)
}

func (sr *SwatchRenderer) Objects() []fyne.CanvasObject {
	return sr.objects
}

func (sr *SwatchRenderer) Destroy() {

}
