package main

import (
	"database/sql"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v5"
	_ "github.com/jackc/pgx/v5/stdlib"
)

// ----------------------------------------------------------------------------
func openDatabase(dataSourceName string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dataSourceName)

	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, err
}

// ----------------------------------------------------------------------------
func (app *application) connectToDatabase() (*sql.DB, error) {
	connection, err := openDatabase(app.DSN)
	if err != nil {
		return nil, err
	}

	return connection, nil
}
