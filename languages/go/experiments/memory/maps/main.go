package main

import (
	"fmt"
	"runtime"
)

// ----------------------------------------------------------------------------
func main() {
	m := make(map[int]Product)

	// add some items to the map
	for i := range 10_000 {
		m[i] = *NewProduct(fmt.Sprintf("Product #%d", i+1), fmt.Sprintf("It is number %d in the map", i), i)
	}
	displayMemoryStatistics("After first allocation")
}

// ----------------------------------------------------------------------------
func displayMemoryStatistics(msg string) {
	fmt.Println()
	fmt.Println(msg)

	// get the memory statistics from the Go runtime
	var stats runtime.MemStats
	runtime.ReadMemStats(&stats)

	// display the stats we are interested in
	fmt.Printf("")
	fmt.Printf("")
	fmt.Printf("")
	fmt.Printf("")
}
