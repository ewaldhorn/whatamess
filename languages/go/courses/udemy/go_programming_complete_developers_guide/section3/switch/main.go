package main

import "fmt"

func main() {
	x := 3

	switch x {
	case 1, 2:
		fmt.Println("We don't care!")
	default:
		fmt.Println("Not something I was expecting!")
	}

	y := "blue"
	switch y {
	case "blue":
		fmt.Println("#0000FF")
	case "red":
		fmt.Println("#FF0000")
	default:
		fmt.Println("#000000")
	}

	switch z := 16; {
	case z > 20:
		fmt.Println("Too big")
	case z < 15:
		fmt.Println("Too small")
	default:
		fmt.Println("In range")
	}

	switch tmp := 15; {
	case tmp%2 == 0:
		fmt.Println("Is even")
	default:
		fmt.Println("Is odd")
	}

	switch x {
	case 3:
		fallthrough
	case 1:
		fmt.Println("X is 1 or 3")
	default:
		fmt.Println("Nope. Don't know what to do!")
	}
}
