package main

const digits = "0123456789abcdef"

type Point struct {
	x, y int
	tag  string
}

var s [32]byte
var msgs = []string{"Hello", "World", "These are", "all", "strings!"}
