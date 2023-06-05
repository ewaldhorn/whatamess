package main

import "fmt"

const (
	ADD Operation = iota
	SUBTRACT
	MULTIPLY
	DIVIDE
)

type Operation int

func (o Operation) calculate(v1, v2 int) int {
	switch o {
	case ADD:
		return v1 + v2
	case SUBTRACT:
		return v1 - v2
	case MULTIPLY:
		return v1 * v2
	case DIVIDE:
		return v1 / v2
	default:
		return 0
	}
}

func main() {
	fmt.Println(ADD.calculate(10, 20))
}
