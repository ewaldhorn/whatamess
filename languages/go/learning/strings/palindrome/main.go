package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s1 := "pop"
	s2 := "topspot"
	s3 := "nope"

	fmt.Printf("%v is a palindrome: %t\n", s1, isPalindrome(s1))
	fmt.Printf("%v is a palindrome: %t\n", s2, isPalindrome(s2))
	fmt.Printf("%v is a palindrome: %t\n", s3, isPalindrome(s3))
}

func isPalindrome(s string) bool {
	runeCount := utf8.RuneCountInString(s)
	middle := runeCount / 2

	for pos := 0; pos < middle; pos++ {
		if s[pos] != s[runeCount-pos-1] {
			return false
		}
	}

	return true
}
