package main

import (
	"fmt"
	"runtime"
)

// ----------------------------------------------------------------------------
func main() {
	m := make(map[int]Product)

	// add some items to the map
	for i := range 100_000 {
		m[i] = *NewProduct(fmt.Sprintf("Product #%d", i+1), fmt.Sprintf("It is number %d in the map", i), i)
	}
	displayMemoryStatistics("After first allocation")
}

// ----------------------------------------------------------------------------
// displayMemoryStatistics displays various memory usage statistics.
// It retrieves memory stats from the Go runtime and displays:
// - Currently allocated heap memory in MiB
// - Total allocated memory since start in MiB
// - Total memory reserved by system in MiB
// - Number of garbage collection cycles completed
// Parameters:
//   - msg: string message to display as a heading before the stats
func displayMemoryStatistics(msg string) {
	fmt.Println()
	fmt.Println(msg)

	// get the memory statistics from the Go runtime
	var stats runtime.MemStats
	runtime.ReadMemStats(&stats)

	// display the stats we are interested in
	fmt.Printf("Allocated heap        = %v MiB\n", byteToMegabyte(stats.Alloc))
	fmt.Printf("Total Allocated       = %v MiB\n", byteToMegabyte(stats.TotalAlloc))
	fmt.Printf("Total memory reserved = %v MiB\n", byteToMegabyte(stats.Sys))
	fmt.Printf("Number of GC cycles   = %v\n", stats.NumGC)
}

// ----------------------------------------------------------------------------
// byteToMegabyte converts bytes to megabytes by dividing by 1024 twice
// to get from bytes -> kilobytes -> megabytes
// Parameters:
//   - b: number of bytes to convert as uint64
//
// Returns:
//   - uint64 representing the equivalent number of megabytes
func byteToMegabyte(b uint64) uint64 {
	return b / 1024 / 1024
}
