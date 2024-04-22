package main

import (
	"fmt"
	"strconv"
)

func main() {
	var x = 3
	var y = 3.1

	var answer = x + int(y)

	fmt.Println(x, "+", y, "=", answer)
	fmt.Printf("X as a string: '%s'\n", fmt.Sprintf("%d", x))
	fmt.Printf("Y as a string: '%s'\n", fmt.Sprintf("%.2f", y))

	fmt.Println("Parsing strings to numbers:")

	strVal := "1.2345678"
	val, _ := strconv.ParseFloat(strVal, 64)
	fmt.Printf("'%s' converted to a 3 decimal float: %.3f\n", strVal, val)
}
