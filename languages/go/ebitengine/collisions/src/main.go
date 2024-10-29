package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

// ----------------------------------------------------------------------------
const (
	APP_VERSION              = "0.0.1"
	SCREEN_WIDTH     float32 = 800
	SCREEN_HEIGHT    float32 = 600
	TICKS_PER_SECOND int     = 60
	IS_DEBUGGING     bool    = false
)

// -------------------------------------------------------------------------
func setupWindow() error {
	ebiten.SetWindowSize(int(SCREEN_WIDTH), int(SCREEN_HEIGHT))
	ebiten.SetWindowTitle("Collisions " + APP_VERSION)
	ebiten.SetTPS(TICKS_PER_SECOND)

	return nil
}

// -------------------------------------------------------------------------
func RunCollisions() error {
	if err := setupWindow(); err != nil {
		return err
	}

	return ebiten.RunGame(NewCollision())
}

// ----------------------------------------------------------------------------
func main() {
	if err := RunCollisions(); err != nil {
		log.Fatalf("Error running screensaver: %v", err)
	}
}
