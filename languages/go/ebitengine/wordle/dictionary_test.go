package main

import (
	"testing"
)

func Test_parseDictionary(t *testing.T) {
	input := []byte{}
	output := parseDictionary(input)

	if len(output) != 0 {
		t.Fatalf("error. expected len(0), got len(%d)\n", len(output))
	}
}

func Test_parseDictionaryBadInput(t *testing.T) {
	inputString := `this
	is
	a
	stupid
	input
	`

	inputBytes := []byte(inputString)
	output := parseDictionary(inputBytes)

	if len(output) != 1 {
		t.Fatalf("error. expected len(1), got len(%d)\n", len(output))
	}

	if output[0] != "input" {
		t.Fatalf("error: expected 'input', got '%s'\n", output)
	}
}
