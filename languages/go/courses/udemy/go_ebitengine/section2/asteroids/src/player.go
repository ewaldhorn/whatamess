package main

import (
	"asteroids/src/assets"
	"math"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/solarlune/resolv"
)

const (
	ROTATION_PER_SECOND = math.Pi
	MAX_ACCELERATION    = 8.0
	shootCoolDown       = 150 * time.Millisecond
	burstCoolDown       = 500 * time.Millisecond
	laserSpawnOffset    = 50.0
)

var currentAcceleration float64
var shotsFired int = 0
var maxShotsPerBurst = 3

// ------------------------------------------------------------------------------------------------
type Player struct {
	gameScene          *GameScene
	sprite             *ebiten.Image
	rotation_angle     float64
	position           Vector
	velocity           float64
	playerObject       *resolv.Circle
	shootCoolDownTimer *Timer
	burstCoolDownTimer *Timer
}

// ------------------------------------------------------------------------------------------------
func (p *Player) Draw(screen *ebiten.Image) {
	bounds := p.sprite.Bounds()
	halfW := float64(bounds.Dx() / 2)
	halfH := float64(bounds.Dy() / 2)

	options := &ebiten.DrawImageOptions{}

	// handle rotation
	options.GeoM.Translate(-halfW, -halfH)
	options.GeoM.Rotate(p.rotation_angle)
	options.GeoM.Translate(halfW, halfH)

	// handle movement
	options.GeoM.Translate(p.position.X, p.position.Y)

	screen.DrawImage(p.sprite, options)
}

// ------------------------------------------------------------------------------------------------
func (p *Player) Update() {
	speed := ROTATION_PER_SECOND / float64(ebiten.TPS())

	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		p.rotation_angle -= speed
	} else if ebiten.IsKeyPressed(ebiten.KeyRight) {
		p.rotation_angle += speed
	}

	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		currentAcceleration -= 4
	}

	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		currentAcceleration += 4
	}
	p.accelerate()
	p.keepOnScreen()
	p.playerObject.SetPosition(p.position.X, p.position.Y)

	// handle shooting
	p.burstCoolDownTimer.Update()
	p.shootCoolDownTimer.Update()
	p.fireLaser()
}

// ------------------------------------------------------------------------------------------------
func (p *Player) fireLaser() {
	if p.burstCoolDownTimer.IsReady() {
		if p.shootCoolDownTimer.IsReady() && ebiten.IsKeyPressed(ebiten.KeySpace) {
			p.shootCoolDownTimer.Reset()
			shotsFired++
			if shotsFired <= maxShotsPerBurst {
				bounds := p.sprite.Bounds()
				halfW := float64(bounds.Dx()) / 2
				halfH := float64(bounds.Dy()) / 2
				spawnPos := Vector{
					p.position.X + halfW + math.Sin(p.rotation_angle)*laserSpawnOffset,
					p.position.Y + halfH + math.Cos(p.rotation_angle)*-laserSpawnOffset,
				}
				p.gameScene.laserCount++
				laser := NewLaser(spawnPos, p.rotation_angle, p.gameScene.laserCount, p.gameScene)
				p.gameScene.lasers[p.gameScene.laserCount] = laser
				p.gameScene.space.Add(laser.laserObj)
			} else {
				p.burstCoolDownTimer.Reset()
				shotsFired = 0
			}
		}
	}
}

// ------------------------------------------------------------------------------------------------
func (p *Player) accelerate() {
	// figure out acceleration
	if currentAcceleration < MAX_ACCELERATION {
		currentAcceleration = p.velocity + 2
	}

	if currentAcceleration > MAX_ACCELERATION {
		currentAcceleration = MAX_ACCELERATION
	}

	p.velocity = currentAcceleration

	// now move in the right direction
	dx := math.Sin(p.rotation_angle) * currentAcceleration
	dy := math.Cos(p.rotation_angle) * -currentAcceleration

	p.position.X += dx
	p.position.Y += dy
}

// ------------------------------------------------------------------------------------------------
func (p *Player) keepOnScreen() {
	if p.position.X >= GAME_WIDTH {
		p.position.X = 0
	} else if p.position.X < 0 {
		p.position.X = GAME_WIDTH
	}

	if p.position.Y >= GAME_HEIGHT {
		p.position.Y = 0
	} else if p.position.Y < 0 {
		p.position.Y = GAME_HEIGHT
	}
}

// ------------------------------------------------------------------------------------------------
func NewPlayer(gameScene *GameScene) *Player {
	bounds := assets.PlayerSprite.Bounds()
	halfW := float64(bounds.Dx()) / 2
	halfH := float64(bounds.Dy()) / 2

	player := &Player{
		gameScene:          gameScene,
		sprite:             assets.PlayerSprite,
		position:           Vector{X: GAME_WIDTH/2 - halfW, Y: GAME_HEIGHT/2 - halfH},
		shootCoolDownTimer: NewTimer(shootCoolDown),
		burstCoolDownTimer: NewTimer(burstCoolDown),
	}

	player.playerObject = resolv.NewCircle(player.position.X, player.position.Y, float64(player.sprite.Bounds().Dx()/2))
	player.playerObject.Tags().Set(TagPlayer)

	return player
}
