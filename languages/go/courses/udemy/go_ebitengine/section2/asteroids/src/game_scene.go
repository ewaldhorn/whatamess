package main

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

// ------------------------------------------------------------------------------------------------
const (
	baseMeteorVelocity  = 0.2
	meteorSpawnTime     = 100 * time.Millisecond
	meteorSpeedUpAmount = 0.1
	meteorSpeedUpTime   = 1000 * time.Millisecond
)

// ------------------------------------------------------------------------------------------------
type GameScene struct {
	player           *Player
	baseVelocity     float64
	meteorCount      int
	meteorSpawnTimer *Timer
	meteors          map[int]*Meteor
	meteorsForLevel  int
	velocityTimer    *Timer
}

// ------------------------------------------------------------------------------------------------
func NewGameScene() *GameScene {
	g := &GameScene{
		meteorSpawnTimer: NewTimer(meteorSpawnTime),
		baseVelocity:     baseMeteorVelocity,
		velocityTimer:    NewTimer(meteorSpeedUpTime),
		meteors:          make(map[int]*Meteor),
		meteorCount:      0,
		meteorsForLevel:  2,
	}

	g.player = NewPlayer(g)

	return g
}

// ------------------------------------------------------------------------------------------------
func (g *GameScene) Update(state *State) error {
	g.updateMeteors()
	g.player.Update()

	return nil
}

// ------------------------------------------------------------------------------------------------
func (g *GameScene) updateMeteors() {
	g.spawnMeteors()

	for _, m := range g.meteors {
		m.Update()
	}

	g.speedUpMeteors()
}

// ------------------------------------------------------------------------------------------------
func (g *GameScene) Draw(screen *ebiten.Image) {
	for _, m := range g.meteors {
		m.Draw(screen)
	}

	g.player.Draw(screen)
}

// ------------------------------------------------------------------------------------------------
func (g *GameScene) spawnMeteors() {
	g.meteorSpawnTimer.Update()

	if g.meteorSpawnTimer.IsReady() {
		g.meteorSpawnTimer.Reset()
		if len(g.meteors) < g.meteorsForLevel && g.meteorCount < g.meteorsForLevel {
			g.meteorCount++
			g.meteors[g.meteorCount] = NewMeteor(g.baseVelocity, g, len(g.meteors)-1)
		}
	}
}

// ------------------------------------------------------------------------------------------------
func (g *GameScene) speedUpMeteors() {
	g.velocityTimer.Update()
	if g.velocityTimer.IsReady() {
		g.velocityTimer.Reset()
		g.baseVelocity += meteorSpeedUpAmount
	}
}

// ------------------------------------------------------------------------------------------------
func (g *GameScene) Layout(outsideWidth, outsideHeight int) (ScreenWidth, ScreenHeight int) {
	return outsideWidth, outsideHeight
}
