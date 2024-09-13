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
func Test_Home_Improved(t *testing.T) {
	var tests = []struct {
		name         string
		putInSession string
		expectedHTML string
	}{
		{"first visit", "", "<small>From Session:"},
		{"second visit", "look at me", "<small>From Session: look at me"},
	}

	for _, test := range tests {
		req, _ := http.NewRequest("GET", "/", nil)
		req = addContextAndSessionToRequest(req, app)
		_ = app.Session.Destroy(req.Context())

		if test.putInSession != "" {
			app.Session.Put(req.Context(), "test", test.putInSession)
		}

		responseRecorder := httptest.NewRecorder()
		handler := http.HandlerFunc(app.Home)
		handler.ServeHTTP(responseRecorder, req)

		if responseRecorder.Code != http.StatusOK {
			t.Errorf("Test_Home_Improved returned status code %d instead of %d", responseRecorder.Code, http.StatusOK)
		}

		body, _ := io.ReadAll(responseRecorder.Body)
		if !strings.Contains(string(body), test.expectedHTML) {
			t.Errorf("(%s) Did not find '%s' in response body", test.name, test.expectedHTML)
		}
	}
}
