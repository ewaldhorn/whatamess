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
	TICKS_PER_SECOND int     = 10
	IS_DEBUGGING     bool    = false
)

// -------------------------------------------------------------------------
func setupWindow() error {
	ebiten.SetWindowSize(int(SCREEN_WIDTH), int(SCREEN_HEIGHT))
	ebiten.SetWindowTitle("Tree " + APP_VERSION)
	ebiten.SetTPS(TICKS_PER_SECOND)

	return nil
}

// -------------------------------------------------------------------------
func RunTrees() error {
	if err := setupWindow(); err != nil {
		return err
	}

	return ebiten.RunGame(NewTree())
}

// ----------------------------------------------------------------------------
func main() {
	if err := RunTrees(); err != nil {
		log.Fatalf("Error running trees: %v", err)
	}
}
