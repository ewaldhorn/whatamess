package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

const (
	PORT = 9000
)

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting server at port %d\n", PORT)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", PORT), nil); err != nil {
		log.Println("error starting server", err)
		os.Exit(2)
	}
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/form" {
		http.Error(w, "Path not found", http.StatusNotFound)
		return
	}

	if r.Method != "POST" {
		http.Error(w, "Unsupported operation", http.StatusNotFound)
		return
	}

	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "Form parsing error: %v", err)
		return
	}

	fmt.Fprintf(w, "<h1>Form received</h1>\n")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name as: %s<br>", name)
	fmt.Fprintf(w, "Address as: %s<br>", address)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "Path not found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Unsupported operation", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "Hello there!")
}
