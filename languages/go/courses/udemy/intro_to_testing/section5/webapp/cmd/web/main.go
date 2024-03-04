package main

import (
	"log"
	"net/http"
)

type application struct{}

func main() {
	const port = "9000"

	app := application{}

	mux := app.routes()

	log.Print("Starting server on:", port)

	err := http.ListenAndServe(string(":"+port), mux)

	if err != nil {
		log.Fatal(err)
	}
}
