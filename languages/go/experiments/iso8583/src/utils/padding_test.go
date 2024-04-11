package utils

import "testing"

func Test_padLeft(t *testing.T) {
	tests := []struct {
		name           string
		input          string
		wantedLength   int
		padWith        string
		expectedResult string
	}{
		{name: "pad with zeroes", input: "1", wantedLength: 5, padWith: "0", expectedResult: "00001"},
		{name: "do nothing", input: "123", wantedLength: 3, padWith: "0", expectedResult: "123"},
		{name: "pad with underscores", input: "123", wantedLength: 8, padWith: "_", expectedResult: "_____123"},
	}

	for _, test := range tests {
		result := padLeft(test.input, test.wantedLength, test.padWith)

		if result != test.expectedResult {
			t.Errorf("%s failed.  expected '%s', got '%s'.\n", test.name, test.expectedResult, result)
		}
	}
}
