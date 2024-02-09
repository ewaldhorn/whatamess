package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Testing!")
	want := "Hello, Testing!"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}