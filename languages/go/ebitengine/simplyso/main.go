package main

import (
	"fmt"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

// ----------------------------------------------------------------------------
const (
	ScreenWidth     = 800
	ScreenHeight    = 600
	meteorSpawnTime = 1 * time.Second

	baseMeteorVelocity  = 0.25
	meteorSpeedUpAmount = 0.1
	meteorSpeedUpTime   = 5 * time.Second
)

// ----------------------------------------------------------------------------
type Game struct {
	player           *Player
	meteorSpawnTimer *Timer
	meteors          []*Meteor
	bullets          []*Bullet
	score            int
	baseVelocity     float64
	velocityTimer    *Timer
}

// ----------------------------------------------------------------------------
func (g *Game) Update() error {
	g.velocityTimer.Update()
	if g.velocityTimer.IsReady() {
		g.velocityTimer.Reset()
		g.baseVelocity += meteorSpeedUpAmount
	}

	g.player.Update()

	g.meteorSpawnTimer.Update()
	if g.meteorSpawnTimer.IsReady() {
		g.meteorSpawnTimer.Reset()

		m := NewMeteor(g.baseVelocity)
		g.meteors = append(g.meteors, m)
	}

	for _, m := range g.meteors {
		m.Update()
	}

	for _, b := range g.bullets {
		b.Update()
	}

	// Check for meteor/bullet collisions
	for i, m := range g.meteors {
		for j, b := range g.bullets {
			if m.Collider().Intersects(b.Collider()) {
				g.meteors = append(g.meteors[:i], g.meteors[i+1:]...)
				g.bullets = append(g.bullets[:j], g.bullets[j+1:]...)
				g.score++
			}
		}
	}

	// Check for meteor/player collisions
	for _, m := range g.meteors {
		if m.Collider().Intersects(g.player.Collider()) {
			g.Reset()
			break
		}
	}
	return nil
}

// ----------------------------------------------------------------------------
func (g *Game) Draw(screen *ebiten.Image) {
	g.player.Draw(screen)

	for _, m := range g.meteors {
		m.Draw(screen)
	}

	for _, b := range g.bullets {
		b.Draw(screen)
	}

	textDrawOptions := text.DrawOptions{}
	textDrawOptions.GeoM.Translate(ScreenWidth/2-100, 50)

	text.Draw(screen, fmt.Sprintf("%06d", g.score), text.NewGoXFace(ScoreFont), &textDrawOptions)
}

// ----------------------------------------------------------------------------
func (g *Game) AddBullet(b *Bullet) {
	g.bullets = append(g.bullets, b)
}

// ----------------------------------------------------------------------------
func (g *Game) Reset() {
	g.player = NewPlayer(g)
	g.meteors = nil
	g.bullets = nil
	g.score = 0
	g.meteorSpawnTimer.Reset()
	g.baseVelocity = baseMeteorVelocity
	g.velocityTimer.Reset()
}

// ----------------------------------------------------------------------------
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return ScreenWidth, ScreenHeight
}

// ----------------------------------------------------------------------------
func main() {

	g := &Game{
		meteorSpawnTimer: NewTimer(meteorSpawnTime),
		baseVelocity:     baseMeteorVelocity,
		velocityTimer:    NewTimer(meteorSpeedUpTime)}

	g.player = NewPlayer(g)

	err := ebiten.RunGame(g)
	if err != nil {
		panic(err)
	}
}
