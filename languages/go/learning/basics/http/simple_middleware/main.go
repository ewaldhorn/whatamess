package main

import (
	"fmt"
	"log"
	"net/http"
)

func withLogging(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path)
		f(w, r)
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "home")
}

func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "about")
}

func main() {
	http.HandleFunc("/", withLogging(home))
	http.HandleFunc("/about", withLogging(about))

	http.ListenAndServe(":9000", nil)
}
