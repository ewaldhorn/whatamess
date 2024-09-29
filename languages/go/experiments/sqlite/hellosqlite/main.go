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

	menu.Option("List all Agents", 1, true, nil)
	menu.Option("Add an Agent", 2, false, nil)
	menu.Option("Update an Agent", 3, false, nil)
	menu.Option("Remove an Agent", 4, false, nil)

	err := menu.Run()
	if err != nil {
		log.Fatal(err)
	}
}

// ----------------------------------------------------------------------------
func handleMenuSelection(options []wmenu.Opt, db *sql.DB) {
	switch options[0].Value {
	case 1:
		dbase.ListAllAgents(db)
	case 2:
		fmt.Println("Add an agent")
	case 3:
		fmt.Println("Update an agent")
	default:
		fmt.Printf("Option %v is not currently supported.\n", options[0].Value)
	}
}
