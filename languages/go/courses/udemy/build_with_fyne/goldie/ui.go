package main

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
)

func (cfg *Config) makeUI(myApp Config) {
	// get current price of gold
	openPrice, currentPrice, priceChange := cfg.getPriceText(myApp)

	// put price info into a container
	priceContent := container.NewGridWithColumns(3, openPrice, currentPrice, priceChange)
	cfg.PriceContainer = priceContent

	// get toolbar
	toolbar := cfg.getToolbar(myApp)
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

	// add a goroutine to refresh the prices every 60 seconds
	go func() {
		for range time.Tick(time.Minute * 1) {
			cfg.refreshGoldPrice(myApp)
		}
	}()
}

func (cfg *Config) refreshGoldPrice(myApp Config) {
	open, current, change := cfg.getPriceText(myApp)
	cfg.PriceContainer.Objects = []fyne.CanvasObject{open, current, change}
	cfg.PriceChartContainer.Refresh()

	chart := cfg.getPriceChart()
	cfg.PriceChartContainer.Objects = []fyne.CanvasObject{chart}
	cfg.PriceChartContainer.Refresh()
}
