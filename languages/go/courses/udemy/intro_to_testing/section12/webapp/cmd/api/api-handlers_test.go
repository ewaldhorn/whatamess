package main

import (
	"context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"time"
	"webapp/pkg/data"

	"github.com/go-chi/chi/v5"
	"github.com/golang-jwt/jwt/v4"
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

// ----------------------------------------------------------------------------
func Test_app_refreshToken(t *testing.T) {
	tests := []struct {
		name               string
		token              string
		expectedStatusCode int
		resetRefreshTime   bool
	}{
		{"valid", "", http.StatusOK, true},
		{"valid but not ready for a refresh", "", http.StatusTooEarly, false},
		{"expired token", createExpiredToken(), http.StatusBadRequest, false},
	}

	testUser := data.User{
		ID:        1,
		FirstName: "Admin",
		LastName:  "User",
		Email:     "admin@example.com",
	}

	oldRefreshTime := refreshTokenExpiry

	for _, test := range tests {
		var token string

		if test.token == "" {
			if test.resetRefreshTime {
				refreshTokenExpiry = time.Second * 1
			}

			tokens, _ := app.generateTokenPair(&testUser)
			token = tokens.RefreshToken
		} else {
			token = test.token
		}

		claims := &Claims{}

		jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(app.JWTSecret), nil
		})

		postedData := url.Values{
			"refresh_token": {token},
		}

		req, _ := http.NewRequest("POST", "/refreshToken", strings.NewReader(postedData.Encode()))
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		rr := httptest.NewRecorder()

		handler := http.HandlerFunc(app.refreshToken)
		handler.ServeHTTP(rr, req)

		if rr.Code != test.expectedStatusCode {
			t.Errorf("%s: expected %d, got %d", test.name, test.expectedStatusCode, rr.Code)
		}

		if test.resetRefreshTime {
			refreshTokenExpiry = oldRefreshTime
		}
	}
}

// ----------------------------------------------------------------------------
func Test_app_user_handlers(t *testing.T) {
	tests := []struct {
		name           string
		method         string
		json           string
		paramID        string
		handler        http.HandlerFunc
		expectedStatus int
	}{
		{"all users", "GET", "", "", app.allUsers, http.StatusOK},
		{"delete user - valid", "DELETE", "", "1", app.deleteUser, http.StatusNoContent},
		{"delete user - invalid", "DELETE", "", "2", app.deleteUser, http.StatusBadRequest},
		{"delete user - bad url param", "DELETE", "", "Q", app.deleteUser, http.StatusBadRequest},
		{"get user - valid", "GET", "", "1", app.retrieveUser, http.StatusOK},
		{"get user - invalid", "GET", "", "2", app.retrieveUser, http.StatusBadRequest},
		{"get user - bad url param", "GET", "", "Red", app.retrieveUser, http.StatusBadRequest},
		{"update user - valid", "PATCH", `{"id":1}`, "", app.updateUser, http.StatusNoContent},
		{"update user - invalid", "PATCH", `{"id":2}`, "", app.updateUser, http.StatusBadRequest},
		{"update user - invalid json", "PATCH", `{id:2}`, "", app.updateUser, http.StatusBadRequest},
		{"insert user - valid", "PUT", `{"first_name":"Pete"}`, "", app.createUser, http.StatusNoContent},
		{"insert user - invalid", "PUT", `{"first_name":"Pop"}`, "", app.createUser, http.StatusBadRequest},
		{"insert user - invalid json", "PUT", `{first_name:"Pete"}`, "", app.createUser, http.StatusBadRequest},
	}

	for _, test := range tests {
		var req *http.Request
		if test.json == "" {
			req, _ = http.NewRequest(test.method, "/", nil)
		} else {
			req, _ = http.NewRequest(test.method, "/", strings.NewReader(test.json))
		}

		if test.paramID != "" {
			chiCtx := chi.NewRouteContext()
			chiCtx.URLParams.Add("userID", test.paramID)
			req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiCtx))
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(test.handler)
		handler.ServeHTTP(rr, req)

		if rr.Code != test.expectedStatus {
			t.Errorf("%s failed with code %d, expected %d", test.name, rr.Code, test.expectedStatus)
		}
	}
}
