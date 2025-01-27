package main

import "testing"

// ====================================================================== TESTS
func Test_validateCompound(t *testing.T) {
	tests := []struct {
		name           string
		inputs         []string
		expectedStatus string
	}{
		{"simple", []string{"CH3(1)CH2(1)CH3"}, "VALID"},
	}

	for _, test := range tests {
		result := validateCompound(test.inputs...)

		if result != test.expectedStatus {
			t.Errorf("%s failed with '%s', expected '%s'.", test.name, result, test.expectedStatus)
		}
	}
}
