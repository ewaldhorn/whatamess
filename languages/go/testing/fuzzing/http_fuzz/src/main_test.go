package main

import (
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
