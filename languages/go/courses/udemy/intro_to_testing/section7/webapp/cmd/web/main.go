package main

import (
	"log"
	"net/http"

	"github.com/alexedwards/scs/v2"
)

// ----------------------------------------------------------------------------
type application struct {
	Session *scs.SessionManager
}

// ----------------------------------------------------------------------------
func main() {
	const port = "9000"

	app := application{}
	mux := app.routes()
	app.Session = getSession()

	log.Print("Starting server on:", port)
	err := http.ListenAndServe(string(":"+port), mux)
	if err != nil {
		log.Fatal(err)
	}
}
