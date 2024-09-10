package main

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// ----------------------------------------------------------------------------
func getContext(req *http.Request) (requestContext context.Context) {
	requestContext = context.WithValue(req.Context(), contextUserKey, "unknown")
	return
}

// ----------------------------------------------------------------------------
func addContextAndSessionToRequest(req *http.Request, app application) *http.Request {
	req = req.WithContext(getContext(req))
	ctx, _ := app.Session.Load(req.Context(), req.Header.Get("X-Session"))
	return req.WithContext(ctx)
}

// ----------------------------------------------------------------------------
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
	pathToPages = pathToTemplates + "pages"

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

// ----------------------------------------------------------------------------
func Test_Home(t *testing.T) {
	// create request to get home page
	req, _ := http.NewRequest("GET", "/", nil)
	req = addContextAndSessionToRequest(req, app)

	responseRecorder := httptest.NewRecorder()

	handler := http.HandlerFunc(app.Home)
	handler.ServeHTTP(responseRecorder, req)

	// check response values is as expected
	if responseRecorder.Code != http.StatusOK {
		t.Errorf("Test_Home returned wrong status code. Expected %d, got %d", http.StatusOK, responseRecorder.Code)
	}

	requestBody, err := io.ReadAll(responseRecorder.Body)
	if err != nil {
		t.Errorf("Reading request body errored. %s", err.Error())
	}

	const lookFor = "<small>From Session:"

	if !strings.Contains(string(requestBody), lookFor) {
		t.Errorf("Could not find `%s` in the body", lookFor)
	}
}
