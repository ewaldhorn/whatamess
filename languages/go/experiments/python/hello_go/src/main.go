package main

import "fmt"
import "C"

// ----------------------------------------------------------------------------
//
//export HelloFromGo
func HelloFromGo() {
	fmt.Println("This is Go, answering the call.")
}

//export CalculateLargeNumber
func CalculateLargeNumber() int64 {
	var answer int64 = 0

	for i := range 5_555_555 {
		answer += int64(i)
	}

	return answer
}

// ----------------------------------------------------------------------------
func main() {}
