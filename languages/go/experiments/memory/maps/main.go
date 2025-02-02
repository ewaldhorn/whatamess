package main

import (
	"fmt"
	"runtime"
)

const COUNT = 100_000

// ----------------------------------------------------------------------------
func main() {
	m := make(map[int]Product)

	// add some items to the map
	for i := range COUNT {
		m[i] = *NewProduct(fmt.Sprintf("Product #%d", i+1), fmt.Sprintf("It is number %d in the map", i), i)
	}
	displayMemoryStatistics("After first allocation")

	// now delete all the entries from the map
	for i := range COUNT {
		delete(m, i)
	}

	displayMemoryStatistics("After deleting all the map entries")

	runtime.GC()
	displayMemoryStatistics("After forcing a garbage collection event")

	// now we make a copy of the map, so there are no references hanging around
	m = refreshMap(m)
	runtime.GC()
	displayMemoryStatistics("After copying map and triggering a GC event")

	// The runtime.KeepAlive() call at the end ensures that the map 'm' is not garbage collected before this point.
	// Without KeepAlive, the compiler could theoretically determine that 'm' is not used after the deletion loop
	// and collect it early. By keeping a reference to 'm' until the end, we can accurately observe the memory
	// impact of the empty but allocated map structure.
	runtime.KeepAlive(m)

	// Now that we no longer need the map, run another GC
	runtime.GC()
	displayMemoryStatistics("One last time after a final GC")
}

// ----------------------------------------------------------------------------
// refreshMap creates a new map and copies the contents of the provided map to it.
// This ensures a completely new map structure is created and returned, effectively
// making a deep copy of the original map. Used to test memory allocation patterns
// by creating fresh maps.
// Parameters:
//   - m: source map[int]Product to copy from
//
// Returns:
//   - map[int]Product containing a copy of all key/value pairs from the input map
func refreshMap(m map[int]Product) map[int]Product {
	newMap := make(map[int]Product)
	for key, val := range m {
		newMap[key] = val
	}
	return newMap
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
