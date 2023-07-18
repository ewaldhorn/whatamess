package main

import (
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func (cfg *Config) getToolbar(myApp Config) *widget.Toolbar {
	toolbar := widget.NewToolbar(
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(theme.DocumentCreateIcon(), func() {}),
		widget.NewToolbarAction(theme.ViewRefreshIcon(), func() { cfg.refreshGoldPrice(myApp) }),
		widget.NewToolbarAction(theme.SettingsIcon(), func() {}),
	)

	return toolbar
}
