package main

import (
	"testing"

	"fyne.io/fyne/v2/test"
	"fyne.io/fyne/v2/widget"
	"github.com/stretchr/testify/assert"
)

// ----------------------------------------------------------------------------
func Test_Text_Selected(t *testing.T) {
	entry := widget.NewEntry()
	test.Type(entry, "Hello World")
	assert.Equal(t, "Hello World", entry.Text)

	test.DoubleTap(entry) // expect to select last word
	assert.Equal(t, "World", entry.SelectedText())
	assert.Equal(t, 0, entry.CursorRow)
	assert.Equal(t, 11, entry.CursorColumn)
}
