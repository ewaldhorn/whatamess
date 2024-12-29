package main

import "testing"

// ----------------------------------------------------------------------------
func TestSimple_formatting(t *testing.T) {
	const input = `{"website":"golangbot.com", "tutorials": [{"title": "Strings", "url":"/strings/"}, {"title":"maps", "url":"/maps/"}, {"title": "goroutines","url":"/goroutines/"}]}`
	const expected = `{
  "website": "golangbot.com",
  "tutorials": [
    {
      "title": "Strings",
      "url": "/strings/"
    },
    {
      "title": "maps",
      "url": "/maps/"
    },
    {
      "title": "goroutines",
      "url": "/goroutines/"
    }
  ]
}`
	result, err := prettifyJSON(input)
	if err != nil {
		t.Errorf("simple_formatting failed with %s", err.Error())
	}
	if result != expected {
		t.Errorf("simple_formatting failed due to unexpected format: `%s`", result)
	}
}
