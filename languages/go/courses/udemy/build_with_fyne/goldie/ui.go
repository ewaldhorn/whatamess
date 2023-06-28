package main

import (
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
)

func (cfg *Config) makeUI() {
	// get current price of gold
	openPrice, currentPrice, priceChange := cfg.getPriceText()

	// put price info into a container
	priceContent := container.NewGridWithColumns(3, openPrice, currentPrice, priceChange)
	cfg.PriceContainer = priceContent

	// get toolbar
	toolbar := cfg.getToolbar()
	cfg.Toolbar = toolbar

	// create price tab content
	priceTabContent := cfg.priceTabContent()

	// get app tabs
	tabs := container.NewAppTabs(
		container.NewTabItemWithIcon("Prices", theme.HomeIcon(), priceTabContent),
		container.NewTabItemWithIcon("Holdings", theme.InfoIcon(), canvas.NewText("Holdings content", nil)),
	)
	tabs.SetTabLocation(container.TabLocationTop)

	// add container to window
	finalContent := container.NewVBox(priceContent, toolbar, tabs)
	cfg.MainWindow.SetContent(finalContent)
}
