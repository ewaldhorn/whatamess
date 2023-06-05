package main

import "fmt"

const myConstant = "This is my constant"
const anotherConstant = 42

const (
	CityKTMCode = iota // 0
	CitkBKTCode        // 1
	CityPKHCode        // 2
)

type Weekday int

const (
	Sunday    Weekday = iota // 0
	Monday                   // 1
	Tuesday                  // 2
	Wednesday                // 3
	Thursday                 // 4
	Friday                   // 5
	Saturday                 // 6
)

type ByteSize float64

const (
	_           = iota             // ignore first value by assigning to blank identifier
	KB ByteSize = 1 << (10 * iota) // 1 << (10*1)
	MB                             // 1 << (10*2)
	GB                             // 1 << (10*3)
	TB                             // 1 << (10*4)
	PB                             // 1 << (10*5)
	EB                             // 1 << (10*6)
	ZB                             // 1 << (10*7)
	YB                             // 1 << (10*8)
)

type Direction byte

const (
	North Direction = iota
	South
	East
	West
)

func (d Direction) String() string {
	switch d {
	case North:
		return "North"
	case South:
		return "South"
	case East:
		return "East"
	case West:
		return "West"
	default:
		return "Error"
	}
}

func (d Direction) CustomString() string {
	return []string{"North", "South", "East", "West"}[d]
}

// values can be skipped, iota increments on NEW LINE only
const (
	_ = iota
	ONE
	TWO
	THREE
	_
	FIVE
)

func main() {
	fmt.Println("Constants are:", myConstant, "and", anotherConstant)
	newDir := North
	fmt.Println(newDir)
	fmt.Println(newDir.CustomString())
}
