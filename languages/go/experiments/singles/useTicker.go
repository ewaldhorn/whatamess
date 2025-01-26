package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(1 * time.Second)
	count := 0
	fmt.Println()

	for {
		count++
		fmt.Printf("Count at %d\n", count)

		if count >= 10 {
			break
		}

		<-ticker.C
	}

	fmt.Println()
}
