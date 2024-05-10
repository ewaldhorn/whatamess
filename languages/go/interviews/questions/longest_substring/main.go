package main

func findLongestSubString(s string) int32 {
	charMap := make(map[rune]int) // keep track of the last place we saw a character
	start, maxLen := 0, 0         // sliding window start and the maximum length of non-repeating characters

	// loop over the string with a sliding window
	for end, char := range s {
		// is the char in the map and to the right? update last position
		if index, found := charMap[char]; found && index >= start {
			start = index + 1
		}

		// check if the current length is longer than the previous maximum length
		currentLength := end - start + 1
		if currentLength > maxLen {
			maxLen = currentLength
		}

		// update the character map with the position
		charMap[char] = end
	}

	return int32(maxLen)
}
