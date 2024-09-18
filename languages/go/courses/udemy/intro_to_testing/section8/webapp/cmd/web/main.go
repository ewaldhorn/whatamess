package main

import (
	"encoding/gob"
	"flag"
	"log"
	"net/http"
	"webapp/pkg/data"
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
const (
	HOME_URL = "/"
	PORT     = "9000"
)

// ----------------------------------------------------------------------------
func main() {
	// register custom types that will be placed in Sessions
	gob.Register(data.User{})

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

	log.Print("Starting server on:", PORT)

	err = http.ListenAndServe(string(":"+PORT), app.routes())
	if err != nil {
		log.Fatal(err)
	}
}
