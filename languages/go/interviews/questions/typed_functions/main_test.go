package main

import "testing"

func Test_serverTypedFunctions(t *testing.T) {
	tests := []struct {
		input       string
		transformer TransformFunc
		expected    string
	}{
		{"yolo.lol", hashFilename, "50cfb0146be621fcf07b39a1d8e1a838acfb3fe8f6da068847358f0bce928d8d"},
		{"lol.lol", prefixFilename, "prefix_lol.lol"},
		{"wat.wot", genericPrefix("oh_"), "oh_wat.wot"},
	}

	for _, test := range tests {
		tmp := &Server{filenameTransformFunc: test.transformer}
		result := tmp.handleRequest(test.input)
		if result != test.expected {
			t.Fatalf("fail: expected '%s', got '%s'.\n", test.expected, result)
		}
	}
}
