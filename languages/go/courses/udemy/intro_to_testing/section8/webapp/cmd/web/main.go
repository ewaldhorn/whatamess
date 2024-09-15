package main

import (
	"flag"
	"log"
	"net/http"
	"webapp/pkg/db"

	"github.com/alexedwards/scs/v2"
)

// ----------------------------------------------------------------------------
type application struct {
	DSN     string
	DB      db.PostgresConn
	Session *scs.SessionManager
}

// ----------------------------------------------------------------------------
func main() {
	const port = "9000"

	app := application{}

	// read command-line "dsn" flag, with a sane default if it is not found
	flag.StringVar(&app.DSN, "dsn", "host=localhost port=5432 user=postgres password=postgres dbname=users sslmode=disable timezone=UCT+2 connect_timeout=5", "Postgres connection")
	flag.Parse()

	dbConnection, err := app.connectToDatabase()
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}
	defer dbConnection.Close()

	app.DB = db.PostgresConn{DB: dbConnection}

	app.Session = getSession()

	log.Print("Starting server on:", port)

	err = http.ListenAndServe(string(":"+port), app.routes())
	if err != nil {
		log.Fatal(err)
	}
}
