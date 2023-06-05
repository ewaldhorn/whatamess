package main

import "fmt"

func footer() {
	fmt.Printf("\nThat's all folks!\n")
}

func header() {
	fmt.Println("Where does this show up?")
}

func body() {
	fmt.Println("\nThe data body")
	fmt.Println("Should be in the middle...")
	fmt.Println("But it isn't, because defer walks the stack backwards...")
}

func main() {
	defer footer()
	fmt.Printf("\nDefer Lesson\n")

	defer header()
	defer body()
}
