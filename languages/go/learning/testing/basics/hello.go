package main

import "fmt"

const prefix = "Hello, ";

func Hello(who string) string {
	if (who == "") {
		who = "World!"
	}
	
	return prefix + who
}

func main() {
	fmt.Println(Hello("Yowsers!"))
}
