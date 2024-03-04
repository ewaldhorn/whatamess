package main

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

// ------------------------------------------------------------------------------------------------
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

// ------------------------------------------------------------------------------------------------
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
		{"six", 6, false, "6 is not a prime number, divisible by 2."},
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

// ------------------------------------------------------------------------------------------------
func Test_prompt(t *testing.T) {
	oldStdout := os.Stdout // preserve a copy of Stdout to restore when we are done

	// intercept os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	prompt()

	// close the writer
	_ = w.Close()

	os.Stdout = oldStdout // now restore Stdout

	// time to read the output
	out, _ := io.ReadAll(r)

	// test if we got what we expected
	if string(out) != "-> " {
		t.Errorf("incorrect prompt: expected '-> ', got '%s'", string(out))
	}
}

// ------------------------------------------------------------------------------------------------
func Test_helloBye_sayWelcome(t *testing.T) {
	oldStdout := os.Stdout

	r, w, _ := os.Pipe()
	os.Stdout = w

	sayWelcome()

	_ = w.Close()

	os.Stdout = oldStdout

	out, _ := io.ReadAll(r)

	if !strings.Contains(string(out), "Or (Q)uit : ") {
		t.Errorf("incorrect prompt: expected to find 'Or (Q)uit : ', got '%s'", string(out))
	}
}

// ------------------------------------------------------------------------------------------------
func Test_helloBye_sayGoodbye(t *testing.T) {
	oldStdout := os.Stdout

	r, w, _ := os.Pipe()
	os.Stdout = w

	sayGoodbye()

	_ = w.Close()

	os.Stdout = oldStdout

	out, _ := io.ReadAll(r)

	if !strings.Contains(string(out), "That was fun,") {
		t.Errorf("incorrect prompt: expected to find 'That was fun,', got '%s'", string(out))
	}
}

// ------------------------------------------------------------------------------------------------
func Test_checkNumbers(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{name: "word", input: "one", expected: "Please enter a whole, positive number"},
		{name: "empty", input: "", expected: "Please enter a whole, positive number"},
		{name: "negative", input: "-1", expected: "Please enter a whole, positive number"},
		{name: "zero", input: "0", expected: "0 is not a prime number."},
		{name: "one", input: "1", expected: "1 is not a prime number."},
		{name: "two", input: "2", expected: "2 is a prime number!"},
		{name: "quit", input: "q", expected: ""},
	}

	for _, e := range tests {
		input := strings.NewReader(e.input)
		reader := bufio.NewScanner(input)
		res, _ := checkNumbers(reader)

		if !strings.EqualFold(e.expected, res) {
			t.Errorf("%s failed: expected '%s', got '%s'", e.name, e.expected, res)
		}
	}
}

// ------------------------------------------------------------------------------------------------
func Test_readUserInput(t *testing.T) {
	allDoneChannel := make(chan bool)
	var stdin bytes.Buffer

	stdin.Write([]byte("1\nq\n"))

	go readUserInput(&stdin, allDoneChannel)
	<-allDoneChannel
	close(allDoneChannel)
}
