// Package utils houses various utility functions required to parse and process Iso8583 messages.
package utils

import "strings"

// padLeft returns a string padded to the required length with the specified padding string.
//
// The original string its length returned if it is >= length.
func padLeft(source string, wantedLength int, padWith string) string {
	if len(source) >= wantedLength {
		return source
	}

	paddingString := strings.Repeat(padWith, wantedLength-len(source))

	return paddingString + source
}
