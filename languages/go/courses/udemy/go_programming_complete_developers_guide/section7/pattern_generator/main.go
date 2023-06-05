package main

import (
	"fmt"
	"math/rand"
)

func generateRandInt(min, max int) <-chan int {
	out := make(chan int, 3)

	go func() {
		for {
			out <- rand.Intn(max-min+1) + min
		}
	}()

	return out
}

func generateRandIntN(count, min, max int) <-chan int {
	out := make(chan int, 1)

	go func() {
		for i := 0; i < count; i++ {
			out <- rand.Intn(max-min+1) + min
		}
		close(out)
	}()

	return out
}

func main() {
	fmt.Printf("\nGenerator Pattern.\n")

	randInt := generateRandInt(1, 100)

	fmt.Println("Random number generator (infinite):")
	fmt.Println(<-randInt)
	fmt.Println(<-randInt)
	fmt.Println(<-randInt)
	fmt.Println(<-randInt)
	fmt.Println(<-randInt)
	fmt.Println(<-randInt)
	fmt.Println(<-randInt)
	fmt.Println(<-randInt)
	fmt.Println(<-randInt)
	fmt.Println(<-randInt)

	randIntN := generateRandIntN(5, 1, 100)

	fmt.Println("Random number generator (finite):")

	fmt.Println(<-randIntN)
	fmt.Println(<-randIntN)
	fmt.Println(<-randIntN)
	fmt.Println(<-randIntN)
	fmt.Println(<-randIntN)
	fmt.Println(<-randIntN)
	fmt.Println(<-randIntN)

	num, open := <-randIntN

	if open {
		fmt.Println("Read", num)
	} else {
		fmt.Println("The channel is close, we'll always get", num)
	}

	fmt.Println()
}
