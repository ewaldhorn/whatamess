package main

/*
https://gist.github.com/JeremyMorgan/1081d3ffb7fe6235a455
Jeremy Morgan
*/

import (
	"fmt"
)

func loopit(phrase string, repeat int) string {

	var outputphrase string = ""

	for i := 0; i < repeat; i++ {
		outputphrase = outputphrase + " " + phrase
	}

	return string(outputphrase)
}

func main() {

	var x [6]string
	x[0] = "Cause the players gonna"
	x[1] = "And the haters gonna"
	x[2] = "I'm just gonna"
	x[3] = "Heart-breakers gonna"
	x[4] = "And the fakers gonna"
	x[5] = "Baby, I'm just gonna"

	var y [6]string
	y[0] = "play"
	y[1] = "hate"
	y[2] = "shake"
	y[3] = "break"
	y[4] = "fake"
	y[5] = "shake"

	for i := 0; i < 6; i++ {
		fmt.Printf("%[1]s%[2]s \n", x[i], loopit(y[i], 5))
		if i == 2 || i == 5 {
			fmt.Printf("%s \n", loopit("Shake it off", 2))
		}
	}
}
