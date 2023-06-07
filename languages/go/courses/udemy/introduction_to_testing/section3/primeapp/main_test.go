package main

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"strings"
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

func Test_printIntro(t *testing.T) {
	// keep a copy of Stdout around
	oldOut := os.Stdout

	// create a read/write pipeline
	r, w, _ := os.Pipe()

	// set output to the new pipe
	os.Stdout = w

	printIntro()

	// now close the writer
	_ = w.Close()

	// restore Stdout
	os.Stdout = oldOut

	// now read the output of the prompt function from the pipe
	out, _ := io.ReadAll(r)

	if !strings.Contains(string(out), "Enter whole numbers to check if they are prime (q to quit).") {
		t.Errorf("incorrect prompt: expected \"Enter whole numbers to check if they are prime (q to quit).\" but got %s", string(out))
	}
}

func Test_checkNumbers(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{name: "empty", input: "", expected: "Please enter a whole number!"},
		{name: "negative", input: "-1", expected: "Negative numbers are not prime, by definition!"},
		{name: "negative", input: "-11", expected: "Negative numbers are not prime, by definition!"},
		{name: "zero", input: "0", expected: "0 is not prime, by definition!"},
		{name: "one", input: "1", expected: "1 is not prime, by definition!"},
		{name: "two", input: "2", expected: "2 is a prime number!"},
		{name: "three", input: "3", expected: "3 is a prime number!"},
		{name: "four", input: "4", expected: "4 is not a prime number because it is divisible by 2"},
		{name: "quit", input: "q", expected: ""},
	}

	for _, test := range tests {
		input := strings.NewReader(test.input)
		reader := bufio.NewScanner(input)
		res, _ := checkNumbers(reader)

		if !strings.EqualFold(res, test.expected) {
			t.Errorf("incorrect value returned: got %s, expected %s", res, test.expected)
		}
	}
}

func Test_readUserInput(t *testing.T) {
	doneChan := make(chan bool)

	var stdin bytes.Buffer
	stdin.Write([]byte("1\nq\n"))

	go readUserInput(&stdin, doneChan)
	<-doneChan
	close(doneChan)
}
