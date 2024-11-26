package main

import (
	"net/http"
)

// ----------------------------------------------------------------------------
type HeaderValue struct {
	key   string
	value string
}

// ----------------------------------------------------------------------------
const prefix = "X-Test-"
const prefixLen = len(prefix)

// ----------------------------------------------------------------------------
func headerModifier(headers http.Header) http.Header {
	updatedHeaders := http.Header{}

	// now loop through the headers and update them
	for header := range headers {
		newName := header[prefixLen:]
		newVal := "non-" + headers.Get(header)
		updatedHeaders.Set(newName, newVal)
	}

	return updatedHeaders
}
