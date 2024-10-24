package main

import (
	"flag"
	"log"
	"webapp/pkg/repository"
	"webapp/pkg/repository/databaserepo"
)

// ----------------------------------------------------------------------------
const port = 9090

// ----------------------------------------------------------------------------
type application struct {
	DSN       string
	DB        repository.DatabaseRepo
	Domain    string
	JWTSecret string
}

// ----------------------------------------------------------------------------
func main() {
	var app application

	// read CLI flags, substitute sane defaults when not found
	flag.StringVar(&app.Domain, "domain", "example.com", "Domain for the API")
	flag.StringVar(&app.DSN, "dsn", "host=localhost port=5432 user=postgres password=postgres dbname=users sslmode=disable timezone=UCT+2 connect_timeout=5", "Postgres connection")
	flag.StringVar(&app.JWTSecret, "jwt-secret", "7iS3j7V9hW5xM2pL4oK0aE8fT1qU6dB9", "signing secret")
	flag.Parse()

	// now perform the database connection
	dbConnection, err := app.connectToDatabase()
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}
	defer dbConnection.Close()

	app.DB = &databaserepo.PostgresDBRepo{DB: dbConnection}
}
