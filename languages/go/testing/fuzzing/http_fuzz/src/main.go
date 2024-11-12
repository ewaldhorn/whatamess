package main

import (
	"fmt"
	"io"
	"net/http"
)

// ----------------------------------------------------------------------------
func status(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, `{"isRunning": true}`)
}

// ----------------------------------------------------------------------------
func echoHeaders(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

// ----------------------------------------------------------------------------
func main() {
	http.HandleFunc("/status", status)
	http.HandleFunc("/echoHeaders", echoHeaders)

	http.ListenAndServe(":9090", nil)
}
