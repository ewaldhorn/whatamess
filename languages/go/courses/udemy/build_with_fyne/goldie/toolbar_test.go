package main

import "testing"

func TestConfig_getToolbar(t *testing.T) {
	tb := testApp.getToolbar(testApp)

	if len(tb.Items) != 4 {
		t.Error("wrong number of items in toolbar")
	}
}
