package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/dixonwille/wmenu/v5"
	_ "github.com/mattn/go-sqlite3"
)

// ----------------------------------------------------------------------------
// Returns a handle to the database, fails if the database could not be opened
func openDatabase(path string) *sql.DB {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		log.Fatal(err)
	}

	return db
}

// ----------------------------------------------------------------------------
func listAllAgents(db *sql.DB) {
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

// ----------------------------------------------------------------------------
func main() {
	db := openDatabase("./hellolite.db")
	defer db.Close() // remember to close database when we are done

	listAllAgents(db)

	menu := wmenu.NewMenu("What is your favorite food?")
	menu.Action(func(opts []wmenu.Opt) error { fmt.Printf(opts[0].Text + " is your favorite food."); return nil })
	menu.Option("Pizza", nil, true, nil)
	menu.Option("Ice Cream", nil, false, nil)
	menu.Option("Tacos", nil, false, func(opt wmenu.Opt) error {
		fmt.Printf("Tacos are great")
		return nil
	})
	err := menu.Run()
	if err != nil {
		log.Fatal(err)
	}
}
