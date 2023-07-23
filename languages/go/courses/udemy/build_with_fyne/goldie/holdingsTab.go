package main

import (
	"goldie/repository"

	"fyne.io/fyne/v2"
)

func (app *Config) holdingsTab() *fyne.Container {
	return nil
}

func (app *Config) getHoldingSlice() [][]interface{} {
	var slice [][]interface{}
	return slice
}

func (app *Config) currentHoldings() ([]repository.Holding, error) {
	var holdings []repository.Holding
	return holdings, nil
}
