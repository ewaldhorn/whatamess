package main

import (
	"context"
	"fmt"
	"strings"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func resolveName(name string) string {
	if strings.Compare("ross", strings.ToLower(name)) == 0 {
		return "HipHip"
	} else if strings.Compare("dale", strings.ToLower(name)) == 0 {
		return "DuckDuck"
	} else {
		return name
	}
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", resolveName(name))
}
