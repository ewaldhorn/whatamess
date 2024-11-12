package main

import "fmt"

// ----------------------------------------------------------------------------
// contains checks if a byte slice contains a specific byte
// Input: input - byte slice to check in, char - value to check for
// Output: returns true if char is found in input, false otherwise
func contains(input []byte, char byte) bool {
	for i := 0; i < len(input); i++ {
		if input[i] == char {
			return true
		}
	}

	return false
}

// ----------------------------------------------------------------------------
func lengthOfLongestNonRepeatingSubString(s string) int {
	var list []byte
	var position, longest, workingLength int

	startind := -1

	for {

		if len([]rune(s)) == 0 {
			return 0
		}

		if contains(list, s[position]) == false {
			list = append(list, s[position])
			workingLength++
		} else if contains(list, s[position]) == true {
			list = nil
			startind++
			position = startind
			workingLength = 0
		}

		if longest < workingLength {
			longest = workingLength
		}

		if len([]rune(s))-1 == position {
			break
		}

		position++

	}

	return longest
}

// ----------------------------------------------------------------------------
func main() {
	fmt.Println("Use the tests, doofus!")
}
