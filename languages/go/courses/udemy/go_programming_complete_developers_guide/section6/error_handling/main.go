package main

import "fmt"

func main() {
	answer, error := divide(5, 0)

	if error == nil {
		fmt.Println("5/0 is ", answer)
	} else {
		fmt.Println(error)
	}

	testStructError()
	testUserError()
}
