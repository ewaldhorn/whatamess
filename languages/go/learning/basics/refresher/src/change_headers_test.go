package main

import (
	"net/http"
	"testing"
)

// ----------------------------------------------------------------------------
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
func Test_changeHeaders(t *testing.T) {
	// create input headers
	headers := http.Header{}
	for i := 0; i < len(inputHeaders); i++ {
		headers.Set(inputHeaders[i].key, inputHeaders[i].value)
	}

	updatedHeaders := headerModifier(headers)

	// now compare the values
	for _, expectedHeader := range expectedHeaders {
		headerValue := updatedHeaders.Get(expectedHeader.key)

		if headerValue == "" {
			t.Fatal("Expected headerValue to be populated!")
		}

		if headerValue != expectedHeader.value {
			t.Fatalf("Expected '%s', got '%s'!\n", headerValue, expectedHeader.value)
		}
	}
}
