package main

import "testing"

func Test_playerStruct(t *testing.T) {
	player := Player{email: "ready@player.one", handle: "Neo"}

	if player.handle != "Neo" {
		t.Fatalf("expected 'Neo', got '%s'\n", player.handle)
	}
}

func Test_pointerToPlayerStruct(t *testing.T) {
	var player *Player

	if player != nil {
		t.Fatal("expected player to be nil!\n")
	}

	player = &Player{email: "mail@me.com", handle: "Mimoo"}

	if player.email != "mail@me.com" {
		t.Fatal("expected 'mail@me.com', got:", player.email)
	}
}
