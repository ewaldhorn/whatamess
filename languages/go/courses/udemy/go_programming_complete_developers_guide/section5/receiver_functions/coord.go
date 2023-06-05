package main

import "fmt"

type Coordinate struct {
	X, Y int
}

func shiftBy(x, y int, coord *Coordinate) {
	coord.X += x
	coord.Y += y
}

func testCoord() {
	coord := Coordinate{11, 14}
	shiftBy(1, 1, &coord)
}

func (coord *Coordinate) shiftBy(x, y int) {
	coord.X += x
	coord.Y += y
}

func testReceiver() {
	coord := Coordinate{5, 5}
	coord.shiftBy(1, 1)
}

func (coord *Coordinate) shiftDistance(other Coordinate) Coordinate {
	return Coordinate{other.X - coord.X, other.Y - coord.Y}
}

func testShift() {
	coord := Coordinate{5, 5}
	distance := Coordinate{1, 1}
	newVal := coord.shiftDistance(distance)
	fmt.Println(newVal)
}
