package main

/*
  #include "withgo.h"
*/
import "C"
import "fmt"

// ======================================================================= main
func main() {
	left := 5
	right := 10
	fmt.Printf("Sum of %d and %d is %d.\n", left, right, goSum(left, right))
}

// ---------------------------------------------------------------------- goSum
func goSum(one int, two int) int {
	result := C.summer(C.int(one), C.int(two))
	return int(result)
}
