package main

import (
	"database/sql"
	"fmt"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

func (a *App) Initialize(user, password, dbname string) {
	fmt.Printf("Received:\nuser:%s\npassword:%s\ndbname:%s\n", user, password, dbname)
}

func (a *App) Run(addr string) {}
