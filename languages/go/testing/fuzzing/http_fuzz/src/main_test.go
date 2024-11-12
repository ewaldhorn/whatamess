package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

// ----------------------------------------------------------------------------
func Test_checkStatus(t *testing.T) {
	req, err := http.NewRequest("GET", "/status", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(status)
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("received wrong status code: got %v, expected %v", status, http.StatusOK)
	}

	expected := `{"isRunning": true}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v, expected %v", rr.Body.String(), expected)
	}
}

// ----------------------------------------------------------------------------
func Test_echoHeaders(t *testing.T) {
	req, err := http.NewRequest("GET", "/echoHeaders", nil)
	if err != nil {
		t.Fatal(err)
	}

	var testHeaders = []struct {
		name  string
		value string
	}{
		{name: "One", value: "1"},
		{name: "Two", value: "two"},
	}

	// add the headers and create the test condition string
	expected := ""
	for _, header := range testHeaders {
		req.Header.Add(header.name, header.value)
		expected += fmt.Sprintf("%v: %v\n", header.name, header.value)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(echoHeaders)
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusSeeOther {
		t.Errorf("received wrong status code: got %v, expected %v", status, http.StatusOK)
	}

	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v, expected %v", rr.Body.String(), expected)
	}
}

// ----------------------------------------------------------------------------
