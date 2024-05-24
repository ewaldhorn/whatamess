package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	fmt.Println("Playing with slices")
	watSlice()
	makeSlices()
	sliceSize()
	sliceAndDice()
	appendGotcha()
	useAppendToRemoveElementsFromSlice()
	testRemoveFromSlice()
}

// safely(ish) remove an element at index from a slice
func removeFromSlice(index int, s []int) []int {
	local := append(make([]int, 0, len(s)), s...)
	return append(local[:index], local[index+1:]...)
}

func testRemoveFromSlice() {
	fmt.Println("\nTest removing from slice, do not affect original")
	s1 := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println("s1:", s1)
	fmt.Println("Removing element #4 (5)")
	s2 := removeFromSlice(4, s1)

	fmt.Println("s1:", s1)
	fmt.Println("s2:", s2)

	fmt.Println("\nTest removing from slice, affect original")
	s1 = []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println("s1:", s1)
	fmt.Println("Removing element #4 (5)")
	s1 = removeFromSlice(4, s1)

	fmt.Println("s1:", s1)
}

func useAppendToRemoveElementsFromSlice() {
	fmt.Println("\nUse append to remove elements from slices")
	// yeah, you can "append" to remove elements from a slice
	// s = s[x:] remove from the top
	// s = s[:len(s)-x] remove from the bottom
	// s = append(s[:x], s[len(s)-x:]...) remove from the middle
	s1 := []int{1, 2, 3, 4, 5, 6}
	fmt.Println("s1:", s1)

	s1 = s1[2:] // remove first two elements
	fmt.Println("s1: ", s1, "(removed first two elements)")

	s1 = []int{1, 2, 3, 4, 5, 6}
	fmt.Println("\ns1:", s1, "(reset to original)")
	s1 = s1[:len(s1)-2] // remove last two elements
	fmt.Println("s1: ", s1, "(removed last two elements)")

	s1 = []int{1, 2, 3, 4, 5, 6}
	fmt.Println("\ns1:", s1, "(reset to original)")
	s1 = append(s1[:2], s1[len(s1)-2:]...) // remove middle two elements
	fmt.Println("s1: ", s1, "(removed middle two elements)")
}

func watSlice() {
	s1 := []int{1, 2, 3}
	fmt.Println("S1:", s1)
	head := s1[:1] // {1}
	tail := s1[1:] // {2, 3}
	fmt.Println("Head is", head)
	fmt.Println("Tail is", tail)

	head = append(head, 4)
	fmt.Println("Head is now:", head)
	fmt.Println("Tail is now:", tail)
	fmt.Println("S1 is now  :", s1)
}

func makeArrays() {
	var a1 [1]int                    // declared but not initialised
	var a2 = [3]int{1, 2, 3}         // declare and initialise
	var a3 = [...]int{1, 2, 3, 4, 5} // declare but you count it for me thanks
	a4 := [...]int{1, 2}             // declare and initialise with shorthand

	fmt.Println(a1, a2, a3, a4)
}

func makeSlices() {
	// useful to remember slices are views sitting on top of an underlying supporting array of the same type
	var s1 []int             // uninitialised slice, value is nil
	var s2 = []int{1, 2, 3}  // initialised slice
	s3 := []int{1, 2, 3}     // shorthand initialisation
	s4 := make([]int, 0, 10) // pre-allocate capacity of 10, handy for performance tuning

	fmt.Printf("s1 len: %d, cap: %d\n", len(s1), cap(s1))
	fmt.Printf("s2 len: %d, cap: %d\n", len(s2), cap(s2))
	fmt.Printf("s3 len: %d, cap: %d\n", len(s3), cap(s3))
	fmt.Printf("s4 len: %d, cap: %d\n", len(s4), cap(s4))
}

func sliceAndDice() {
	fmt.Println()
	// s := a[start:end] includes start, excludes end, indexing starts at 0
	a := [...]int{1, 2, 3} // array [3]int {1,2,3}
	s1 := a[:2]            // slice []int{1,2} (from 0..1)
	s2 := a[1:]            // slice []int{2,3} (from 1..end)  will have same start as s3
	s3 := a[1:2]           // slice []int{2} (from 1..1) will have same start as s2

	fmt.Print("s1:")
	printSliceStructInformation(s1)

	fmt.Print("s2:")
	printSliceStructInformation(s2)

	fmt.Print("s3:")
	printSliceStructInformation(s3)
}

func printSliceStructInformation[S ~[]E, E any](s S) {
	sh := (*reflect.SliceHeader)(unsafe.Pointer(&s))
	fmt.Printf("%#v\n", sh)
}

func sliceSize() {
	fmt.Println("\nSlice growth")
	s1 := make([]int, 0, 5)
	fmt.Printf("s1 len: %d, cap: %d\n", len(s1), cap(s1))
	fmt.Println("s1:", s1)

	fmt.Println("\nNow we add five numbers.")
	for i := range 5 {
		s1 = append(s1, i+1)
	}
	fmt.Printf("s1 len: %d, cap: %d\n", len(s1), cap(s1))
	fmt.Println("s1:", s1)

	fmt.Println("\nNow we add another five numbers.")
	for i := range 5 {
		s1 = append(s1, i+6)
	}
	fmt.Printf("s1 len: %d, cap: %d\n", len(s1), cap(s1))
	fmt.Println("s1:", s1)

	fmt.Println("\nNow we add another three numbers. (Forces underlying array capacity to grow)")
	for i := range 3 {
		s1 = append(s1, i+11)
	}
	fmt.Printf("s1 len: %d, cap: %d\n", len(s1), cap(s1))
	fmt.Println("s1:", s1)

	fmt.Println("\nNow we add another 11 numbers. (Forces underlying array capacity to grow again)")
	for i := range 11 {
		s1 = append(s1, i+14)
	}
	fmt.Printf("s1 len: %d, cap: %d\n", len(s1), cap(s1))
	fmt.Println("s1:", s1)

	fmt.Println("Seems like small array size doubles...")
}

func appendGotcha() {
	fmt.Println("\nSlice appending gotcha")
	s1 := []int{1, 2, 3}
	tail := s1[1:]
	fmt.Println("s1  :", s1)
	fmt.Println("tail:", tail)

	fmt.Println("Appending to tail")

	tail = append(tail, 4, 5, 6)

	fmt.Println("s1  :", s1)
	fmt.Println("tail:", tail)

}
