package main

import "fmt"

// ----------------------------------------------------------------------------
func withSlices() {
	names := []string{}

	displayNames(names)

	addName("Colin", &names)
	displayNames(names)

	listNames(names)

	addName("Bob", &names)
	addName("Sarah", &names)
	listNames(names)
}

// ----------------------------------------------------------------------------
func displayNames(names []string) {
	fmt.Printf("\nThere are %d names in the slice.\n", len(names))
}

// ----------------------------------------------------------------------------
func addName(name string, names *[]string) {
	*names = append(*names, name)
}

// ----------------------------------------------------------------------------
func listNames(names []string) {
	fmt.Println("The names are:")

	for idx, name := range names {
		fmt.Printf("%d: %s\n", idx+1, name)
	}
}
