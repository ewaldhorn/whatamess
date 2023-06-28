package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

func (cfg *Config) priceTabContent() *fyne.Container {
	chart := cfg.getPriceChart()
	chartContainer := container.NewVBox(chart)
	cfg.PriceChartContainer = chartContainer

	return chartContainer
}
