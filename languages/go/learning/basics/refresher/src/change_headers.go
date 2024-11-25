package main

import (
	"fmt"
	"net/http"
)

// ----------------------------------------------------------------------------
type HeaderValue struct {
	key   string
	value string
}

var inputHeaders = []HeaderValue{
	{key: "X-Test-0", value: "test-value-1"},
	{key: "X-Test-1", value: "test-value-2"},
	{key: "X-Test-Alpha", value: "some-thing"},
}

var expectedHeaders = []HeaderValue{
	{key: "0", value: "non-test-value-1"},
	{key: "1", value: "non-test-value-2"},
	{key: "Alpha", value: "non-some-thing"},
}

// ----------------------------------------------------------------------------
const prefix = "X-Test-"
const prefixLen = len(prefix)

// ----------------------------------------------------------------------------
func headerModifier() {
	// create input headers
	headers := http.Header{}
	for i := 0; i < len(inputHeaders); i++ {
		headers.Set(inputHeaders[i].key, inputHeaders[i].value)
	}

	updatedHeaders := http.Header{}
	// now loop through the headers and update them
	for header := range headers {
		newName := header[prefixLen:]
		newVal := "non-" + headers.Get(header)
		updatedHeaders.Set(newName, newVal)
	}

	// now compare the values
	for _, expectedHeader := range expectedHeaders {
		headerValue := updatedHeaders.Get(expectedHeader.key)

		if headerValue == "" {
			fmt.Println("Didn't expect this!")
		}

		if headerValue != expectedHeader.value {
			fmt.Printf("Expected '%s', got '%s'!\n", headerValue, expectedHeader.value)
		}
	}

	fmt.Println("Success!?")
}
