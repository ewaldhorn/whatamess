package pxcanvas

import "fyne.io/fyne/v2"

func (p *PxCanvas) scale(direction int) {
	switch {
	case direction > 0:
		if p.PxSize < 50 {
			p.PxSize += 1
		}
	case direction < 0:
		if p.PxSize > 2 {
			p.PxSize -= 1
		}
	default:
		p.PxSize = 10
	}
}

func (p *PxCanvas) pan(prevCoord, curCoord fyne.PointEvent) {
	xDiff := curCoord.Position.X - prevCoord.Position.X
	yDiff := curCoord.Position.Y - prevCoord.Position.Y

	p.CanvasOffset.X += xDiff
	p.CanvasOffset.Y += yDiff
	p.Refresh()
}
