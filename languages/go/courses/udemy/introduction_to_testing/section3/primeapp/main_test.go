package main

import (
	"io"
	"os"
	"testing"
)

func Test_isPrime(t *testing.T) {
	primeTests := []struct {
		name     string
		testNum  int
		expected bool
		msg      string
	}{
		{"not prime", -1, false, "Negative numbers are not prime, by definition!"},
		{"not prime", 0, false, "0 is not prime, by definition!"},
		{"not prime", 1, false, "1 is not prime, by definition!"},
		{"prime", 7, true, "7 is a prime number!"},
		{"not prime", 8, false, "8 is not a prime number because it is divisible by 2"},
	}

	for _, e := range primeTests {
		result, msg := isPrime(e.testNum)
		if e.expected && !result {
			t.Errorf("%s: expected true but got false", e.name)
		}

		if !e.expected && result {
			t.Errorf("%s: expected false but got true", e.name)
		}

		if e.msg != msg {
			t.Errorf("%s: expected %s but got %s", e.name, e.msg, msg)
		}
	}
}

func Test_prompt(t *testing.T) {
	// keep a copy of Stdout around
	oldOut := os.Stdout

	// create a read/write pipeline
	r, w, _ := os.Pipe()

	// set output to the new pipe
	os.Stdout = w

	prompt()

	// now close the writer
	_ = w.Close()

	// restore Stdout
	os.Stdout = oldOut

	// now read the output of the prompt function from the pipe
	out, _ := io.ReadAll(r)

	if string(out) != "-> " {
		t.Errorf("incorrect prompt: expected -> but got %s", string(out))
	}
}
