package main

import (
	"fmt"
	"net/http"
)

// ----------------------------------------------------------------------------
func status(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "server up\n")
}

// ----------------------------------------------------------------------------
func echoHeaders(w http.ResponseWriter, req *http.Request) {
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
