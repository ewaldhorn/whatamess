package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// -------------------------------------------------------------------------------------------------
func handleHTTPRequests(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "You requested: %s\n", r.URL.Path)
}

// -------------------------------------------------------------------------------------------------
func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", handleHTTPRequests)

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.ListenAndServe(":9000", nil)
}
