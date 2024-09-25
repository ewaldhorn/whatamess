package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"webapp/pkg/data"
)

// ----------------------------------------------------------------------------
func Test_app_auth(t *testing.T) {
	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})

	tests := []struct {
		name   string
		isAuth bool
	}{
		{"logged in", true},
		{"not logged in", false},
	}

	for _, test := range tests {
		handlerToTest := app.auth(nextHandler)
		req := httptest.NewRequest("GET", "http://testing", nil)
		req = addContextAndSessionToRequest(req, app)
		if test.isAuth {
			app.Session.Put(req.Context(), "user", data.User{ID: 1})
		}
		rr := httptest.NewRecorder()
		handlerToTest.ServeHTTP(rr, req)

		if test.isAuth && rr.Code != http.StatusOK {
			t.Errorf("%s: expected %d got %d", test.name, http.StatusOK, rr.Code)
		}

		if !test.isAuth && rr.Code != http.StatusUnauthorized {
			t.Errorf("%s: expected %d got %d", test.name, http.StatusUnauthorized, rr.Code)
		}
	}
}
