package main

import (
	"errors"
	"fmt"
	"strings"
)

// createCombinationVariants takes an email address as input and returns a slice of
// string combinations based on the name portion (before '@') of the email address.
// Each combination consists of two strings arranged in multiple possible ways.
// If the input email address is empty or doesn't contain '@', it returns an error.
//
// The returned combinations include:
// - First character removed from name
// - Last character removed from name
// - Last character plus remaining characters
// - All but last character plus last character
// - All but first character plus first character
// - First character plus remaining characters
func createCombinationVariants(emailAddress string) ([][2]string, error) {
	if emailAddress == "" || !strings.Contains(emailAddress, "@") {
		return nil, errors.New("invalid email address supplied")
	}

	nameSection := strings.Split(emailAddress, "@")[0]

	combinations := [][2]string{
		{nameSection[1:], ""},
		{nameSection[:len(nameSection)-1], ""},
		{string(nameSection[len(nameSection)-1]), nameSection[:len(nameSection)-1]},
		{nameSection[:len(nameSection)-1], string(nameSection[len(nameSection)-1])},
		{nameSection[1:], string(nameSection[0])},
		{string(nameSection[0]), nameSection[1:]},
	}

	return combinations, nil
}

func main() {
	fmt.Println(createCombinationVariants("simplesimon@toodles.com"))
	fmt.Println(createCombinationVariants("brinks@there.com"))
}
