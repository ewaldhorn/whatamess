package main

import (
	"database/sql"
	"fmt"
	"log"
	"yoyosql/dbase"

	"github.com/dixonwille/wmenu/v5"
)

// ----------------------------------------------------------------------------
func main() {
	db := dbase.OpenDatabase("./hellolite.db")
	defer db.Close() // remember to close database when we are done

	fmt.Println("")
	menu := wmenu.NewMenu("Select an action:")
	menu.Action(func(options []wmenu.Opt) error { handleMenuSelection(options, db); return nil })

	menu.Option("List all Agents", 0, true, nil)
	menu.Option("Add an Agent", 1, false, nil)

	err := menu.Run()
	if err != nil {
		log.Fatal(err)
	}
}

// ----------------------------------------------------------------------------
func handleMenuSelection(options []wmenu.Opt, db *sql.DB) {
	switch options[0].Value {
	case 0:
		dbase.ListAllAgents(db)
	case 1:
		fmt.Println("Add an agent")
	}
}
