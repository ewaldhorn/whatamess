package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_application_handlers(t *testing.T) {
	var tests = []struct {
		name               string
		url                string
		expectedStatusCode int
	}{
		{"home", "/", http.StatusOK},
		{"about", "/about", http.StatusOK},
		{"nowhere", "/bada_bing", http.StatusNotFound},
	}

	// override path so tests can find templates too
	pathToTemplates = "./../../templates/"

	var app application
	routes := app.routes()

	// create a new test server
	ts := httptest.NewTLSServer(routes)
	defer ts.Close()

	for _, test := range tests {
		rsp, err := ts.Client().Get(ts.URL + test.url)

		if err != nil {
			t.Log(err)
			t.Fatal(err)
		}

		if rsp.StatusCode != test.expectedStatusCode {
			t.Errorf("for %s: expected %d received %d", test.name, test.expectedStatusCode, rsp.StatusCode)
		}
	}
}
