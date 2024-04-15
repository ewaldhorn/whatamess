package main

import (
	"consumers/internal/bar"
	"consumers/internal/foo"
)

func main() {
	bar := &bar.Bar{}
	foo := foo.NewFoo(bar)

	foo.Greet()
}
