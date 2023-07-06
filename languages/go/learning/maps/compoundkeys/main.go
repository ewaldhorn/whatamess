package main

import "fmt"

type Point struct {
	X, Y int
}

func main() {
	colours := make(map[Point]int)
	pt := Point{10, 20}
	colours[pt] = 0xFF0000
	fmt.Printf("%X\n", colours[pt])
}
