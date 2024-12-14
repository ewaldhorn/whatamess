package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"webapp/pkg/data"
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

// ----------------------------------------------------------------------------
func Test_app_authRequired(t *testing.T) {
	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})

	testUser := data.User{
		ID: 1, FirstName: "Bob", LastName: "Soap", Email: "bob@soap.com",
	}

	tokens, err := app.generateTokenPair(&testUser)
	if err != nil {
		t.Errorf("failed to generate tokens: %s", err.Error())
	}

	tests := []struct {
		name             string
		token            string
		expectAuthorized bool
		setHeader        bool
	}{
		{"valid token", fmt.Sprintf("Bearer %s", tokens.Token), true, true},
		{"no token", "", false, false},
		{"invalid token", fmt.Sprintf("Bearer 1%s1", tokens.Token), false, true},
		{"expired token", fmt.Sprintf("Bearer %s", createExpiredToken()), false, true},
	}

	for _, test := range tests {
		req, err := http.NewRequest("GET", "/", nil)
		if err != nil {
			t.Errorf("%s did not expect to fail now: %s", test.name, err.Error())
		}

		if test.setHeader {
			req.Header.Set("Authorization", test.token)
		}

		rr := httptest.NewRecorder()
		handlerToTest := app.authRequired(nextHandler)
		handlerToTest.ServeHTTP(rr, req)

		if test.expectAuthorized && rr.Code == http.StatusUnauthorized {
			t.Errorf("%s is unauthorized, this is unexpected", test.name)
		}

		if !test.expectAuthorized && rr.Code != http.StatusUnauthorized {
			t.Errorf("%s should be unauthorized, it's not!", test.name)
		}
	}

}
