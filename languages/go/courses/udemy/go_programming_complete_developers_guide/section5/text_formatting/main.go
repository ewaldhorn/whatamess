package main

import "fmt"

func surround(msg string, left rune, right rune) string {
	return fmt.Sprintf("%c%v%c", left, msg, right)
}

func main() {
	const number = 10

	fmt.Println()
	fmt.Println("Text Formatting!")

	fmt.Println()
	fmt.Printf("%d in hex is %x or %X\n", number, number, number)
	fmt.Printf("%[1]d in hex is %[1]x or %[1]X\n", number)

	fmt.Println()
	fmt.Printf("Is number > 10? %v\n", number > 10)
	fmt.Printf("Is number > 10? %t\n", number > 10)

	fmt.Println()
	fmt.Println("Escaped 41 yields character \x41 ")
	fmt.Println("Escaped 61 yields character \x61 ")

	fmt.Println()
	fmt.Println(surround("My message", '#', '#'))

	fmt.Println()
}
