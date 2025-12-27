package main

import "github.com/hajimehoshi/ebiten/v2"

// ------------------------------------------------------------------------------------------------
var (
	transitionFrom = ebiten.NewImage(GAME_WIDTH, GAME_HEIGHT)
	transitionTo   = ebiten.NewImage(GAME_WIDTH, GAME_HEIGHT)
)

// ------------------------------------------------------------------------------------------------
const (
	TRANSITION_MAX_COUNT = 25
)

// ------------------------------------------------------------------------------------------------
type Scene interface {
	Update(state *State) error
	Draw(screen *ebiten.Image)
}

// ------------------------------------------------------------------------------------------------
type State struct {
	SceneManager *SceneManager
	Input        *Input
}

// ------------------------------------------------------------------------------------------------
type SceneManager struct {
	currentScene    Scene
	nextScene       Scene
	transitionCount int
}

// ------------------------------------------------------------------------------------------------
func (s *SceneManager) Draw(r *ebiten.Image) {
	if s.transitionCount == 0 {
		s.currentScene.Draw(r)
		return
	}

	transitionFrom.Clear()
	s.currentScene.Draw(transitionFrom)

	transitionTo.Clear()
	s.nextScene.Draw(transitionTo)

	r.DrawImage(transitionFrom, nil)

	alpha := 1.0 - float32(s.transitionCount/TRANSITION_MAX_COUNT)
	op := &ebiten.DrawImageOptions{}
	op.ColorScale.ScaleAlpha(alpha)
	r.DrawImage(transitionTo, op)
}

// ------------------------------------------------------------------------------------------------
func (s *SceneManager) Update(_ *Input) error {
	if s.transitionCount == 0 {
		return s.currentScene.Update(&State{
			SceneManager: s,
		})
	}

	s.transitionCount--
	if s.transitionCount > 0 {
		return nil
	}

	s.currentScene = s.nextScene
	s.nextScene = nil
	return nil
}

// ------------------------------------------------------------------------------------------------
func (s *SceneManager) GoToScene(scene Scene) {
	if s.currentScene == nil {
		s.currentScene = scene
	} else {
		s.nextScene = scene
		s.transitionCount = TRANSITION_MAX_COUNT
	}
}
