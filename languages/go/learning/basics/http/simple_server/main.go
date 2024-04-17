package main

import (
	"fmt"
	"net/http"
)

// -------------------------------------------------------------------------------------------------
func handleHTTPRequests(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "You requested: %s\n", r.URL.Path)
}

// -------------------------------------------------------------------------------------------------
func main() {
	http.HandleFunc("/", handleHTTPRequests)
	http.ListenAndServe(":9000", nil)
}
