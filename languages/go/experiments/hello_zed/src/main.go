package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello there Zed!")

	startTime := time.Now()
	sirDoesALot()
	elapsedTime := time.Since(startTime)

	fmt.Println("It took", elapsedTime, "to run 'sirDoesALot'")
}

func sirDoesALot() {
	time.Sleep(10 * time.Second)
}
