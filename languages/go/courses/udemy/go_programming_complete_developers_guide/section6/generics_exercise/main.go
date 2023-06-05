package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type Distance int32
type Velocity float64

type Number interface {
	constraints.Float | constraints.Integer
}

func clamp[T Number](value, min, max T) T {
	if value > max {
		return max
	}

	if value < min {
		return min
	}

	return value
}

func testClampInt8() {
	var (
		min int8 = -10
		max int8 = 10
	)

	if clamp(-30, min, max) != min {
		panic("clamp failed for int8")
	} else {
		fmt.Println("-30 clamped is:", clamp(-30, min, max))
	}
}

func main() {
	fmt.Printf("Generics Exercise\n")

	testClampInt8()

	fmt.Println()
	fmt.Println()
}
