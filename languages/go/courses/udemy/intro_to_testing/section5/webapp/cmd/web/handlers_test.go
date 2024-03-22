package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// ------------------------------------------------------------------------------------------------
func Test_application_handlers(t *testing.T) {
	var tests = []struct {
		name               string
		url                string
		expectedStatusCode int
	}{
		{"home", "/", http.StatusOK},
		{"not found", "/thisIsNotAvailable", http.StatusNotFound},
	}

	routes := app.routes()

	testServer := httptest.NewTLSServer(routes)
	defer testServer.Close()

	pathToTemplates = "./../../templates/"

	for _, test := range tests {
		rsp, err := testServer.Client().Get(testServer.URL + test.url)
		if err != nil {
			t.Log(err)
			t.Fatal(err)
		}

		if rsp.StatusCode != test.expectedStatusCode {
			t.Errorf("%s failed with a status code of %d, expected %d", test.name, rsp.StatusCode, test.expectedStatusCode)
		}
	}
}
