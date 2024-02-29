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
