package main

import "fmt"

type second uint

func main() {
	var hour second = 3600

	fmt.Printf("An hour contains %d minutes.\n", hour/60)
	fmt.Printf("An hour contains %d seconds.\n", hour)
}
