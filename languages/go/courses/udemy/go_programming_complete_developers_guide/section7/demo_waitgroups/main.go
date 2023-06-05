package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	fmt.Printf("\nWaitGroups Demo.\n")

	var wg sync.WaitGroup
	counter := 0

	for i := 0; i < 25; i++ {
		wg.Add(1)
		counter += 1
		go func() {
			defer func() {
				fmt.Println(counter, "goroutines remaining")
				counter -= 1
				wg.Done()
			}()

			duration := time.Duration(rand.Intn(6000)*int(time.Millisecond) + 50)
			fmt.Println("Wait for", duration, "second(s).")
			time.Sleep(duration)
		}()
	}
	fmt.Println("Waiting for goroutines to finish...")
	wg.Wait()
	fmt.Println("Done!")

	fmt.Println()
}
