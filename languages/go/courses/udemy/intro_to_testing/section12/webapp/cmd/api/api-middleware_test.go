package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// ----------------------------------------------------------------------------
func Test_app_enableCORS(t *testing.T) {
	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})

	tests := []struct {
		name           string
		method         string
		expectedHeader bool
	}{
		{"preflight", "OPTIONS", true},
		{"get", "GET", false},
	}

	for _, test := range tests {
		handlerToTest := app.enableCORS(nextHandler)
		req := httptest.NewRequest(test.method, "http://testing", nil)
		rr := httptest.NewRecorder()
		handlerToTest.ServeHTTP(rr, req)

		if test.expectedHeader && rr.Header().Get("Access-Control-Allow-Credentials") == "" {
			t.Errorf("%s expected header 'Access-Control-Allow-Credentials', did not find it, ", test.name)
		}

		if !test.expectedHeader && rr.Header().Get("Access-Control-Allow-Credentials") != "" {
			t.Errorf("%s did not expect a header!", test.name)
		}
	}
}
