package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"image"
	"image/png"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"path"
	"strings"
	"sync"
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
		name                    string
		url                     string
		expectedStatusCode      int
		expectedURL             string
		expectedFirstStatusCode int
	}{
		{"home", "/", http.StatusOK, HOME_URL, http.StatusOK},
		{"not found", "/thisIsNotAvailable", http.StatusNotFound, "/thisIsNotAvailable", http.StatusNotFound},
		{"profile", "/user/profile", http.StatusUnauthorized, "/user/profile", http.StatusUnauthorized},
	}

	routes := app.routes()

	testServer := httptest.NewTLSServer(routes)
	defer testServer.Close()

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{
		Transport: tr,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	for _, test := range tests {
		rsp, err := testServer.Client().Get(testServer.URL + test.url)
		if err != nil {
			t.Log(err)
			t.Fatal(err)
		}

		if rsp.StatusCode != test.expectedStatusCode {
			t.Errorf("%s failed with a status code of %d, expected %d", test.name, rsp.StatusCode, test.expectedStatusCode)
		}

		if rsp.Request.URL.Path != test.expectedURL {
			t.Errorf("%s failed with path '%s' vs expected path '%s'", test.name, rsp.Request.URL.Path, test.expectedURL)
		}

		resp2, _ := client.Get(testServer.URL + test.url)
		if resp2.StatusCode != test.expectedFirstStatusCode {
			t.Errorf("%s: expected first returned status code to be %d but got %d", test.name, test.expectedFirstStatusCode, resp2.StatusCode)
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
		{
			name:               "invalid login - wrong password",
			postedData:         url.Values{"email": {"admin@example.com"}, "password": {"secrets"}},
			expectedStatusCode: http.StatusSeeOther,
			expectedLocation:   "/",
		},
		{
			name:               "invalid login - no password",
			postedData:         url.Values{"email": {"admin@example.com"}},
			expectedStatusCode: http.StatusSeeOther,
			expectedLocation:   "/",
		},
		{
			name:               "invalid login - no email",
			postedData:         url.Values{"password": {"secret"}},
			expectedStatusCode: http.StatusSeeOther,
			expectedLocation:   "/",
		},
		{
			name:               "invalid login - no such user",
			postedData:         url.Values{"email": {"wntwrk@example.com"}, "password": {"sucker"}},
			expectedStatusCode: http.StatusSeeOther,
			expectedLocation:   "/",
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

// ----------------------------------------------------------------------------
func Test_app_UploadFiles(t *testing.T) {
	// setup pipes for IO
	readPipe, writePipe := io.Pipe()
	writer := multipart.NewWriter(writePipe)

	// create a waitgroup and add '1' to it
	waitGroup := &sync.WaitGroup{}
	waitGroup.Add(1)

	// simulate uploading a file with a goroutine
	go simulatePNGFileUpload("./testdata/img.png", writer, t, waitGroup)

	// read from the pipe
	request := httptest.NewRequest("POST", "/", readPipe)
	request.Header.Add("Content-Type", writer.FormDataContentType())

	// call UploadFiles
	uploadedFiles, err := app.UploadFiles(request, "./testdata/uploads/")
	if err != nil {
		t.Error("error uploading file", err)
	}

	uploadedFilePath := fmt.Sprintf("./testdata/uploads/%s", uploadedFiles[0].OriginalFilename)
	// run tests - basically, make sure the file got uploaded
	if _, err := os.Stat(uploadedFilePath); os.IsNotExist(err) {
		t.Errorf("expected file to exist: %s", err.Error())
	}

	// clean up by removing the uploaded file so the next test does not pass by accident
	err = os.Remove(uploadedFilePath)
	if err != nil {
		t.Error(err)
	}
}

// ----------------------------------------------------------------------------
func simulatePNGFileUpload(fileToUpload string, writer *multipart.Writer, t *testing.T, waitGroup *sync.WaitGroup) {
	defer writer.Close()
	defer waitGroup.Done()

	// create form data field nammed 'file' with the filename as the value
	part, err := writer.CreateFormFile("file", path.Base(fileToUpload))
	if err != nil {
		t.Error("error creating form data field", err)
	}

	// open the file to upload
	file, err := os.Open(fileToUpload)
	if err != nil {
		t.Error("error opening file", err)
	}
	defer file.Close()

	// decode the image
	decodedImage, _, err := image.Decode(file)
	if err != nil {
		t.Error("error decoding image", err)
	}

	// now write the decoded image to the IO writer
	err = png.Encode(part, decodedImage)
	if err != nil {
		t.Error("error writing file to writer", err)
	}
}
