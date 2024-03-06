package main

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func Test_application_addIPToContext(t *testing.T) {
	tests := []struct {
		headerName   string
		headerValue  string
		address      string
		emptyAddress bool
	}{
		{"", "", "", false},
		{"", "", "", true},
		{"X-Forwarded-For", "192.3.2.1", "", false},
		{"X-Forwarded-For", "", "", true},
		{"", "", "hello:world", false},
	}

	var app application

	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		val := r.Context().Value(contextUserKey)
		if val == nil {
			t.Error(contextUserKey, "not present")
		}

		ip, ok := val.(string)
		if !ok {
			t.Error("not a string")
		}

		t.Log(ip)
	})

	for _, test := range tests {
		handlerToTest := app.addIPToContext(nextHandler)
		req := httptest.NewRequest("GET", "http://testing", nil)

		if test.emptyAddress {
			req.RemoteAddr = ""
		}

		if len(test.headerName) > 0 {
			req.Header.Add(test.headerName, test.headerValue)
		}

		if len(test.address) > 0 {
			req.RemoteAddr = test.address
		}

		handlerToTest.ServeHTTP(httptest.NewRecorder(), req)
	}
}

// ------------------------------------------------------------------------------------------------
func Test_application_ipFromContext(t *testing.T) {
	var app application
	var ctx = context.Background()

	ctx = context.WithValue(ctx, contextUserKey, "whatever")
	ip := app.ipFromContext(ctx)

	if !strings.EqualFold("whatever", ip) {
		t.Error("expected 'whatever', got:", ip)
	}
}
