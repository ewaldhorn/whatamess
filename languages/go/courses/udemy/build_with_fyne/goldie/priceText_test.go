package main

import "testing"

func TestConfig_getPriceText(t *testing.T) {
	open, _, _ := testApp.getPriceText()

	if open == nil {
		t.Error("text component is nil, should not be!")
	}

	//if open.Text != "Open: 1234.5678 USD" {
	//	t.Error("wrong price text field returned:", open.Text)
	//}
}
