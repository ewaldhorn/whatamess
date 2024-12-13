package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"webapp/pkg/data"
)

// ----------------------------------------------------------------------------
func Test_app_getAndVerifyTokenFromHeader(t *testing.T) {
	testUser := data.User{ID: 1, FirstName: "AdminF", LastName: "AdminL", Email: "admin@example.com"}
	tokens, _ := app.generateTokenPair(&testUser)

	tests := []struct {
		name          string
		token         string
		errorExpected bool
		setHeader     bool
		issuer        string
	}{
		{"valid", fmt.Sprintf("Bearer %s", tokens.Token), false, true, app.Domain},
		{"valid - expired", fmt.Sprintf("Bearer %s", createExpiredToken()), true, true, app.Domain},
		{"no header", "", true, true, app.Domain},
		{"invalid", fmt.Sprintf("Bearer 1%s1", tokens.Token), true, true, app.Domain},
		{"no bearer", fmt.Sprintf("%s", tokens.Token), true, true, app.Domain},
		{"bad bearer", fmt.Sprintf("Bearrer %s", tokens.Token), true, true, app.Domain},
		{"bad header", fmt.Sprintf("Token Bearer %s", tokens.Token), true, true, app.Domain},
		{"valid - bad issuer", fmt.Sprintf("Bearer %s", tokens.Token), true, true, "nolansdonuts"},
	}

	for _, test := range tests {
		if test.issuer != app.Domain {
			app.Domain = test.issuer
			tokens, _ = app.generateTokenPair(&testUser)
		}

		req, _ := http.NewRequest("GET", "/", nil)
		if test.setHeader {
			req.Header.Set("Authorization", test.token)
		}

		rr := httptest.NewRecorder()

		_, _, err := app.getAndVerifyTokenFromHeader(rr, req)
		if err != nil && !test.errorExpected {
			t.Errorf("%s was not supposed to error, got %s", test.name, err.Error())
		}

		if err == nil && test.errorExpected {
			t.Errorf("%s was supposed to error, it didn't!", test.name)
		}

		app.Domain = "example.com"
	}
}
