package main

import "testing"

func Test_parseDictionary(t *testing.T) {
	input := []byte{}
	output := parseDictionary(input)

	if len(output) != 0 {
		t.Fatalf("error. expected len(0), got len(%d)", len(output))
	}
}
