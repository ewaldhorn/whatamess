package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_application_AddIPToContext(t *testing.T) {
	tests := []struct {
		headerName  string
		headerValue string
		addr        string
		emptyAddr   bool
	}{
		{"", "", "", false},
		{"", "", "", true},
		{"X-Forwarded-For", "192.3.2.1", "", false},
		{"", "", "hello:world", false},
	}

	var app application

	// create a handler just to check the context
	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// check that the IP address exists in the context
		val := r.Context().Value(contextUserKey)
		if val == nil {
			t.Error(contextUserKey, "not present")
		}

		// ensure we got a string back
		ip, ok := val.(string)

		if !ok {
			t.Error("not a string")
		}
		t.Log(ip)
	})

	for _, test := range tests {
		handlerToTest := app.addIPToContext(nextHandler)
		req := httptest.NewRequest("GET", "http://testing", nil)

		if test.emptyAddr {
			req.RemoteAddr = ""
		}

		if len(test.headerName) > 0 {
			req.Header.Add(test.headerName, test.headerValue)
		}

		if len(test.addr) > 0 {
			req.RemoteAddr = test.addr
		}

		handlerToTest.ServeHTTP(httptest.NewRecorder(), req)
	}
}
