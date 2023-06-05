package pxcanvas

import (
	"image"
	"image/color"
	"pixl/apptype"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
)

type PxCanvasMouseState struct {
	previousCoord *fyne.PointEvent
}

type PxCanvas struct {
	widget.BaseWidget
	apptype.PxCanvasConfig
	renderer    *PxCanvasRenderer
	PixelData   image.Image
	mouseState  PxCanvasMouseState
	appState    *apptype.State
	reloadImage bool
}

func (p *PxCanvas) Bounds() image.Rectangle {
	x0 := int(p.CanvasOffset.X)
	y0 := int(p.CanvasOffset.Y)
	x1 := int(p.PxCols*p.PxSize + int(p.CanvasOffset.X))
	y1 := int(p.PxRows*p.PxSize + int(p.CanvasOffset.Y))

	return image.Rect(x0, y0, x1, y1)
}

func InBounds(pos fyne.Position, bounds image.Rectangle) bool {
	return pos.X >= float32(bounds.Min.X) && pos.X < float32(bounds.Max.X) &&
		pos.Y >= float32(bounds.Min.Y) && pos.Y < float32(bounds.Max.Y)
}

func NewBlankImage(cols, rows int, c color.Color) image.Image {
	img := image.NewNRGBA(image.Rect(0, 0, cols, rows))

	for y := 0; y < rows; y++ {
		for x := 0; x < cols; x++ {
			img.Set(x, y, c)
		}
	}

	return img
}

func NewPxCanvas(state *apptype.State, config apptype.PxCanvasConfig) *PxCanvas {
	pxCanvas := &PxCanvas{
		PxCanvasConfig: config,
		appState:       state,
		PixelData:      NewBlankImage(config.PxCols, config.PxRows, color.NRGBA{128, 128, 128, 255}),
	}

	pxCanvas.ExtendBaseWidget(pxCanvas)

	return pxCanvas
}

func (pxCanvas *PxCanvas) CreateRenderer() fyne.WidgetRenderer {
	canvasImage := canvas.NewImageFromImage(pxCanvas.PixelData)
	canvasImage.ScaleMode = canvas.ImageScalePixels
	canvasImage.FillMode = canvas.ImageFillContain
	canvasBorder := make([]canvas.Line, 4)

	for _, b := range canvasBorder {
		b.StrokeColor = color.NRGBA{100, 100, 100, 255}
		b.StrokeWidth = 2
	}

	renderer := &PxCanvasRenderer{
		pxCanvas:     pxCanvas,
		canvasImage:  canvasImage,
		canvasBorder: canvasBorder,
	}
	pxCanvas.renderer = renderer

	return renderer
}

func (p *PxCanvas) TryPan(prevCoord *fyne.PointEvent, ev *desktop.MouseEvent) {
	if prevCoord != nil && ev.Button == desktop.MouseButtonTertiary {
		p.pan(*prevCoord, ev.PointEvent)
	}
}

func (p *PxCanvas) SetColor(c color.Color, x, y int) {
	if ngrba, ok := p.PixelData.(*image.NRGBA); ok {
		ngrba.Set(x, y, c)
	}

	if rgba, ok := p.PixelData.(*image.RGBA); ok {
		rgba.Set(x, y, c)
	}

	p.Refresh()
}

func (p *PxCanvas) MouseToCanvasXY(ev *desktop.MouseEvent) (*int, *int) {
	bounds := p.Bounds()

	if !InBounds(ev.Position, bounds) {
		return nil, nil
	}

	pxSize := float32(p.PxSize)
	xOffset := p.CanvasOffset.X
	yOffset := p.CanvasOffset.Y

	x := int((ev.Position.X - xOffset) / pxSize)
	y := int((ev.Position.Y - yOffset) / pxSize)

	return &x, &y
}

func (p *PxCanvas) LoadImage(img image.Image) {
	dimensions := img.Bounds()

	p.PxCanvasConfig.PxCols = dimensions.Dx()
	p.PxCanvasConfig.PxRows = dimensions.Dy()
	p.PixelData = img
	p.reloadImage = true
	p.Refresh()
}

func (p *PxCanvas) NewDrawing(cols, rows int) {
	p.appState.SetFilePath("")
	p.PxCols = cols
	p.PxRows = rows
	tmpImage := NewBlankImage(cols, rows, color.NRGBA64{128, 128, 128, 255})
	p.LoadImage(tmpImage)
}
