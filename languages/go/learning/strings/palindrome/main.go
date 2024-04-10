package main

import (
	"unicode/utf8"
)

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
