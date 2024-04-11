package utils

import "strings"

func padLeft(source string, length int, padWith string) string {
	if len(source) >= length {
		return source
	}
	paddedString := strings.Repeat(padWith, length-len(source))
	return paddedString + source
}
