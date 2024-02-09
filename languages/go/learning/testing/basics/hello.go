package main

import "fmt"

func Hello(who string) string {
	return "Hello, " + who
}

func main() {
	fmt.Println(Hello("Yowsers!"))
}
