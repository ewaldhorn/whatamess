package dbase

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

// ----------------------------------------------------------------------------
// Returns a handle to the database, fails if the database could not be opened
func OpenDatabase(path string) *sql.DB {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		log.Fatal(err)
	}

	return db
}

// ----------------------------------------------------------------------------
func ListAllAgents(db *sql.DB) {
	rows, err := db.Query("select id, name, score from agents")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var name string
		var score int
		err = rows.Scan(&id, &name, &score)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(id, name, score)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}
