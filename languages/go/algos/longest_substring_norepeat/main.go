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
// lengthOfLongestNonRepeatingSubString finds the longest substring without repeating chars
// Input: sourceStr - string to analyze
// Output: returns length of longest non-repeating substring found
func lengthOfLongestNonRepeatingSubString(sourceStr string) int {
	var workingString []byte
	var position, longest, workingLength int

	startind := -1

	for {

		if len([]rune(sourceStr)) == 0 {
			return 0
		}

		if contains(workingString, sourceStr[position]) == false {
			workingString = append(workingString, sourceStr[position])
			workingLength++
		} else if contains(workingString, sourceStr[position]) == true {
			workingString = nil
			startind++
			position = startind
			workingLength = 0
		}

		if longest < workingLength {
			longest = workingLength
		}

		if len([]rune(sourceStr))-1 == position {
			break
		}

		position++

	}

	return longest
}

// ----------------------------------------------------------------------------
// lengthOfLongestNonRepeatingSubStringWithMap finds longest substring without repeats using a map
// Input: sourceStr - string to analyze
// Output: returns length of longest non-repeating substring found using sliding window approach
func lengthOfLongestNonRepeatingSubStringWithMap(sourceStr string) int {
	letterPos := make(map[byte]int)
	maxLength := 0
	start := 0
	byteStr := []byte(sourceStr)

	for end := 0; end < len(byteStr); end++ {
		if pos, exists := letterPos[byteStr[end]]; exists && pos >= start {
			start = pos + 1
		}
		letterPos[byteStr[end]] = end
		if maxLength < end-start+1 {
			maxLength = end - start + 1
		}
	}

	return maxLength
}

// ----------------------------------------------------------------------------
func main() {
	fmt.Println("Use the tests, doofus!")
}
