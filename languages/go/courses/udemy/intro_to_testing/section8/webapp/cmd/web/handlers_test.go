package main

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
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

// ----------------------------------------------------------------------------
func Test_Render_Bad_Template(t *testing.T) {
	// set up path to a bad template in the testdata directory
	// keep a copy around to restore later
	tmpPathToTemplates := pathToTemplates
	tmpPathToPages := pathToPages

	pathToTemplates = "./testdata/"
	pathToPages = pathToTemplates + "pages"

	req, _ := http.NewRequest("GET", "/", nil)
	req = addContextAndSessionToRequest(req, app)
	responseRecorder := httptest.NewRecorder()

	err := app.render(responseRecorder, req, "bad.page.gohtml", &TemplateData{})
	if err == nil {
		t.Errorf("Expected an error from the template, none was received")
	}

	// restore page and template paths
	pathToTemplates = tmpPathToTemplates
	pathToPages = tmpPathToPages
}

// ----------------------------------------------------------------------------
func Test_app_Login(t *testing.T) {
	var tests = []struct {
		name               string
		postedData         url.Values
		expectedStatusCode int
		expectedLocation   string
	}{
		{
			name:               "valid login",
			postedData:         url.Values{"email": {"admin@example.com"}, "password": {"secret"}},
			expectedStatusCode: http.StatusSeeOther,
			expectedLocation:   "/user/profile",
		},
	}

	for _, test := range tests {
		req, err := http.NewRequest("POST", "/login", strings.NewReader(test.postedData.Encode()))
		if err != nil {
			t.Errorf("unexpected error: %s ", err.Error())
		}

		req = addContextAndSessionToRequest(req, app)
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		requestRecorder := httptest.NewRecorder()
		handler := http.HandlerFunc(app.Login)
		handler.ServeHTTP(requestRecorder, req)

		if requestRecorder.Code != test.expectedStatusCode {
			t.Errorf("error in test '%s'. Got %d, expected %d", test.name, requestRecorder.Code, test.expectedStatusCode)
		}

		actualLocation, err := requestRecorder.Result().Location()
		if err != nil {
			t.Errorf("unable to retrieve location for '%s'", test.name)
		}

		if actualLocation.String() != test.expectedLocation {
			t.Errorf("location for '%s' is '%s' instead of '%s'", test.name, actualLocation.String(), test.expectedLocation)
		}
	}
}
