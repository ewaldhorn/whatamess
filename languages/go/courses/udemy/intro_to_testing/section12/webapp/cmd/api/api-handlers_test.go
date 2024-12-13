package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// ----------------------------------------------------------------------------
func Test_app_authenticate(t *testing.T) {
	tests := []struct {
		name               string
		requestBody        string
		expectedStatusCode int
	}{
		{"valid user", `{"email":"admin@example.com","password":"secret"}`, http.StatusOK},
		{"valid user - wrong password", `{"email":"admin@example.com","password":"nosecret"}`, http.StatusUnauthorized},
		{"invalid user - valid password", `{"email":"administer@example.com","password":"secret"}`, http.StatusUnauthorized},
		{"invalid user - invalid password", `{"email":"administer@example.com","password":"ohsosecret"}`, http.StatusUnauthorized},
		{"invalid JSON", `"email":"admin@example.com","password":"secret"}`, http.StatusBadRequest},
		{"empty JSON", `{}`, http.StatusUnauthorized},
		{"missing email", `{"password":"secret"}`, http.StatusUnauthorized},
		{"missing password", `{"email":"admin@example.com"}`, http.StatusUnauthorized},
	}

	for _, test := range tests {
		// var reader io.Reader
		reader := strings.NewReader(test.requestBody)
		req, _ := http.NewRequest("POST", "/auth", reader)
		rr := httptest.NewRecorder()

		handler := http.HandlerFunc(app.authenticate)
		handler.ServeHTTP(rr, req)

		if test.expectedStatusCode != rr.Code {
			t.Errorf("%s returned %d, expected %d", test.name, rr.Code, test.expectedStatusCode)
		}
	}
}
