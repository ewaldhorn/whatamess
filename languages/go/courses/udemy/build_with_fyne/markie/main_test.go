package main

import (
	"testing"

	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/test"
)

func Test_makeUI(t *testing.T) {
	var testCfg config

	edit, preview := testCfg.createMainUI()

	test.Type(edit, "Hello")

	if preview.String() != "Hello" {
		t.Error("Failed: Did not find expected value in preview.")
	}
}

func Test_runApp(t *testing.T) {
	var testCfg config
	testApp := test.NewApp()
	testWin := testApp.NewWindow("Test Markie")

	edit, preview := testCfg.createMainUI()
	testCfg.createMainMenu(testWin)

	testWin.SetContent(container.NewHSplit(edit, preview))
	testApp.Run()

	test.Type(edit, "# This is a test.")

	if preview.String() != "# This is a test." {
		t.Error("Expected value not correct.")
	}
}
