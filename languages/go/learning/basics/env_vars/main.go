package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Using environment variables in Go")

	fmt.Println("You passed in:", getGreetingEnvironmentVariable())
}

func getGreetingEnvironmentVariable() string {
	maybe := os.Getenv("GREETING")

	return maybe
}
