package main

import "fmt"

func main() {
	myMap := map[string]string{
		"one":   "this is one",
		"two":   "this is two",
		"three": "and this is three",
	}

	fmt.Println("myMap has", len(myMap), "items:")

	for key, val := range myMap {
		fmt.Printf("\t'%s' maps to '%s'\n", key, val)
	}

	fmt.Println("\nIt's likely that the elements are returned in a random order:")
	for key, val := range myMap {
		fmt.Printf("\t'%s' maps to '%s'\n", key, val)
	}

}
