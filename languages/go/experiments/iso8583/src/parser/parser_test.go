package parser

import "testing"

func Test_readMTIFromString(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"all zeroes", "0000000000000", "0000"},
		{"all ones", "111100000000000", "1111"},
		{"sequence", "123456700000000", "1234"},
		{"real message", "02003220000000808000000010000000001500120604120000000112340001840", "0200"},
	}

	for _, test := range tests {
		result := readMTIFromString(test.input)

		if result.MTI != test.expected {
			t.Errorf("%s failed.  expected '%s', got '%s'\n", test.name, test.expected, result.MTI)
		}
	}
}
