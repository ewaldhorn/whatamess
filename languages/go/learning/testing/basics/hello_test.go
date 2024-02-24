package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("say hello with valid parameter", func(t *testing.T) {
		got := Hello("Testing!")
		want := "Hello, Testing!"

		assertCorrectMessageReceived(t, got, want)
	})

	t.Run("say hello with empty parameter", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World!"

		assertCorrectMessageReceived(t, got, want)
	})

}

// Accept a T(est) or B(enchmark) interface.
func assertCorrectMessageReceived(t testing.TB, got, want string) {
	// tells the test runner this is a helper function
	// when it fails the line number reported will be in our function call rather than inside our test helper
	t.Helper() 

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}