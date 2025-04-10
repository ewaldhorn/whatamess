package main

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

// ----------------------------------------------------------------------------
type Game struct {
	count       int
	bouncers    []Bouncer
	pressedKeys []ebiten.Key
	lineWidth   float32
}

// ----------------------------------------------------------------------------
func (g *Game) initBouncers() {
	g.bouncers = make([]Bouncer, g.count)

	for bouncerPosition := range g.bouncers {
		tmpBouncer := Bouncer{}
		tmpBouncer.init()
		g.bouncers[bouncerPosition] = tmpBouncer
	}
}

// ----------------------------------------------------------------------------
func (g *Game) Update() error {
	if ebiten.IsFocused() {
		g.pressedKeys = inpututil.AppendJustPressedKeys(g.pressedKeys[:0])

		for _, key := range g.pressedKeys {
			switch key.String() {
			case "ArrowDown":
				if g.lineWidth > 0.20 {
					g.lineWidth -= 0.10
				}
			case "ArrowUp":
				if g.lineWidth < 50.0 {
					g.lineWidth += 0.10
				}
			}
		}

		for pos, bouncer := range g.bouncers {
			bouncer.update()
			g.bouncers[pos] = bouncer
		}
	}

	return nil
}

// ----------------------------------------------------------------------------
func (g *Game) Draw(screen *ebiten.Image) {
	str := fmt.Sprintf("We are at roughly %.0f FPS, more or less. (Line: %0.2f) Focus: %t", ebiten.ActualFPS(), g.lineWidth, ebiten.IsFocused())
	ebitenutil.DebugPrint(screen, str)

	for i := 1; i < len(g.bouncers); i++ {
		vector.StrokeLine(screen,
			float32(g.bouncers[i-1].positionX),
			float32(g.bouncers[i-1].positionY),
			float32(g.bouncers[i].positionX),
			float32(g.bouncers[i].positionY),
			g.lineWidth,
			g.bouncers[i].colour, true)
	}

	lastBouncer := g.bouncers[len(g.bouncers)-1]
	vector.StrokeLine(screen,
		float32(lastBouncer.positionX), float32(lastBouncer.positionY),
		float32(g.bouncers[0].positionX), float32(g.bouncers[0].positionY),
		g.lineWidth, lastBouncer.colour, true)
}

// ----------------------------------------------------------------------------
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return SCREEN_WIDTH, SCREEN_HEIGHT
}
