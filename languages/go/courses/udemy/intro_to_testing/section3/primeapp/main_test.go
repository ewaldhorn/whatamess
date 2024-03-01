package main

import "testing"

func Test_isPrime(t *testing.T) {
	result, msg := isPrime(0)

	if result {
		t.Errorf("with %d as test parameter, got true, expected false", 0)
	}

	if msg != "0 is not a prime number." {
		t.Errorf("with %d, expected '%s', got '%s'", 0, "0 is not a prime number.", msg)
	}

	result, msg = isPrime(1)

	if result {
		t.Errorf("with %d as test parameter, got true, expected false", 1)
	}

	if msg != "1 is not a prime number." {
		t.Errorf("with %d, expected '%s', got '%s'", 1, "1 is not a prime number.", msg)
	}

	result, msg = isPrime(2)

	if !result {
		t.Errorf("with %d as test parameter, got true, expected false", 2)
	}

	if msg != "2 is a prime number!" {
		t.Errorf("with %d, expected '%s', got '%s'", 2, "2 is a prime number!", msg)
	}
}

func Test_isPrime_WithTables(t *testing.T) {
	primeTests := []struct {
		name     string
		number   int
		expected bool
		msg      string
	}{
		{"negative", -10, false, "-10 is not a prime number."},
		{"zero", 0, false, "0 is not a prime number."},
		{"one", 1, false, "1 is not a prime number."},
		{"two", 2, true, "2 is a prime number!"},
		{"three", 3, true, "3 is a prime number!"},
		{"four", 4, false, "4 is not a prime number, divisible by 2."},
		{"five", 5, true, "5 is a prime number!"},
	}

	for _, test := range primeTests {
		result, msg := isPrime(test.number)

		if test.expected != result {
			t.Errorf("%s: expected '%t' but got '%t'", test.name, test.expected, result)
		}

		if msg != test.msg {
			t.Errorf("with %d, expected '%s', got '%s'", test.number, test.msg, msg)
		}
	}
}
