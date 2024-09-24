package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./hellolite.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

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
