package main

import "fmt"

func main() {
	print("\n")

	// the usual suspects
	fmt.Println("This is a text line")
	fmt.Print("This", "is another", "text", "line\n")
	fmt.Println("This", "is", "yet another", "text", "line")
	fmt.Printf("A %s string\n", "formatted")

	// determine the Type of a variable
	myType := struct {
		name string
		age  int
	}{name: "Ben", age: 10}

	fmt.Printf("myType contains (%s, %d) and is of type `%T`\n", myType.name, myType.age, myType)

	fmt.Printf("We also have: '%#v'\n", myType)
	fmt.Printf("We also have: '%#v'\n", 11)

	// booleans
	fmt.Printf("Dividing 12 by four equals three: %t\n", 12/4 == 3)

	// rounding, decimal places
	print("\n")
	theFraction := 1.23456789

	fmt.Printf("theFraction is %f\n", theFraction)
	fmt.Printf("theFraction is %.1f\n", theFraction)
	fmt.Printf("theFraction is %.2f\n", theFraction)
	fmt.Printf("theFraction is %.3f\n", theFraction)
	fmt.Printf("theFraction is %.4f\n", theFraction)
	fmt.Printf("theFraction is %.5f\n", theFraction)
	fmt.Printf("theFraction is %.6f\n", theFraction)
	fmt.Printf("theFraction is %.7f\n", theFraction)
	fmt.Printf("theFraction is %.8f\n", theFraction)
	fmt.Printf("theFraction is %.9f\n", theFraction)
	fmt.Printf("theFraction is %.10f\n", theFraction)

	print("\n")
	println("Or with a loop:")

	for i := range 10 {
		fmt.Printf("theFraction is %.*f\n", i+1, theFraction)
	}

	// various number formats
	print("\n")
	theNumber := 11

	fmt.Printf("The number in binary is %b\n", theNumber)
	fmt.Printf("The number in binary is %08b, as 0 padded 8 bit byte\n", theNumber)
	fmt.Printf("The number in binary is %016b, as 0 padded 16 bit byte\n", theNumber)

	fmt.Printf("The number in hex is %x, lower case\n", theNumber)
	fmt.Printf("The number in hex is %X, upper case\n", theNumber)
	fmt.Printf("The number in hex is %04X, upper case, 0 padded to length 4\n", theNumber)

	fmt.Printf("The number in octal is %o\n", theNumber)
	fmt.Printf("The number in octal is %05o, zero padded to length of 5\n", theNumber)

	// quote strings
	print("\n")
	fmt.Printf("We can also quote something: %q\n", "abc")

	println()
}
