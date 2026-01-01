package main

import (
	"fmt"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/solarlune/resolv"
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
	space            *resolv.Space // collision detection space
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
		space:            resolv.NewSpace(GAME_WIDTH, GAME_HEIGHT, 16, 16),
	}

	g.player = NewPlayer(g)
	g.space.Add(g.player.playerObject)

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
	g.checkIfPlayerCollidedWithMeteor()
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
			g.space.Add(g.meteors[g.meteorCount].meteorObject)
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

// ------------------------------------------------------------------------------------------------
func (g *GameScene) checkIfPlayerCollidedWithMeteor() {
	for _, m := range g.meteors {
		if m.meteorObject.IsIntersecting(g.player.playerObject) {
			data := m.meteorObject.Data().(*ObjectData)
			fmt.Printf("Player collided with meteor #%d!\n", data.index)
		}
	}
}
