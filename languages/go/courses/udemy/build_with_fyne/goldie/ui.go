package main

import "fyne.io/fyne/v2/container"

func (cfg *Config) makeUI() {
	// get current price of gold
	openPrice, currentPrice, priceChange := cfg.getPriceText()

	// put price info into a container
	priceContent := container.NewGridWithColumns(3, openPrice, currentPrice, priceChange)
	cfg.PriceContainer = priceContent

	// add container to window
	finalContent := container.NewVBox(priceContent)
	cfg.MainWindow.SetContent(finalContent)
}
