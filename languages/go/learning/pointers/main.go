package main

import "fmt"

func main() {
	var myPtr *int
	count := 2
	myPtr = &count

	// print the value at the pointer address by dereferencing the pointer using *
	fmt.Println("We have:", *myPtr, "in the pointer at", myPtr)

	// now test structures etc
	lib, err := NewLibrary("Cape Central Library", "Grand Parade")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("The %s at %s has %d books.\n", lib.name, lib.address, len(lib.books))
	fmt.Println("Adding a book...")
	lib.addBook(Book{title: "Readability", author: "Joe Soap", ISBN: "0091000191"})
	fmt.Printf("The %s at %s has %d books.\n", lib.name, lib.address, len(lib.books))
}
