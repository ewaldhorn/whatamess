// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {
	fmt.Println(countOccurrences("foo", [][]string{
		{"foo", "bar", "foo"},
		{"foo", "$stop", "foo"},
	}))
}

// countOccurrences counts the number of times needle occurs within haystack.
// If it encounters the string `$stop` in haystack, it aborts calculation early.
func countOccurrences(needle string, haystack [][]string) int {
	var count int
stop:
	for _, x := range haystack {
		check := func(needle, word string) bool {
			return needle == word
		}
		for _, y := range x {
			if y == "$stop" {
				break stop
			}
			if check(needle, y) {
				count++
			}
		}
	}
	return count
}
