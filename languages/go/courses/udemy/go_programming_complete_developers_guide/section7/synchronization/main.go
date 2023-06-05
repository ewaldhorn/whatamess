package main

import "fmt"

func main() {
	fmt.Printf("\nGo Synchronization.\n")

	runMutexExample1()
	runMutexExample2()
	waitGroupExample()

	fmt.Println()
}
