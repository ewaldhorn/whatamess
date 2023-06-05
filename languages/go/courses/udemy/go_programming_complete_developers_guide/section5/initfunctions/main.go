package main

import (
	"fmt"
	mypackage "initfunctions/myPackage"
)

func init() {
	fmt.Println("Hello from init")
}

func main() {
	fmt.Println("Main function called")
	mypackage.ThisIsCallable()
}
