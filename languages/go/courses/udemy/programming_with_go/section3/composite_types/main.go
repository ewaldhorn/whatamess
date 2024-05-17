package main

import (
	"fmt"
	"math"
	"strconv"
)

type MyData struct {
	descr   string
	notes   string
	version int
}

func main() {
	castIntToBiggerInt()
	convertIntToString()
	convertStringToInt()
	someMath()
}

func castIntToBiggerInt() {
	var smallInt int8 = 20
	var bigInt int64 = int64(smallInt)

	if bigInt != int64(smallInt) {
		fmt.Println("This is a problem!")
	}
}

func convertIntToString() {
	var start int32 = 500
	end := strconv.Itoa(int(start))

	if end != "500" {
		fmt.Println("This is a problem:", end)
	}
}

func convertStringToInt() {
	var start = "500"
	end, err := strconv.Atoi(start)

	if err != nil {
		fmt.Println("Problem converting str to i:", err)
	}

	if end != 500 {
		fmt.Println("Oh my!")
	}
}

func someMath() {
	twoPower8 := math.Pow(2.0, 8.0)
	fmt.Println("Two ** Eight =", twoPower8)
}
