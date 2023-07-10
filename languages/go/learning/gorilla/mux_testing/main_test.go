package main

import (
	"log"
	"os"
	"testing"
)

var testApp App

func TestMain(m *testing.M) {
	testApp.Initialize(
		os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_NAME"))

	ensureTableExists()
	code := m.Run()
	clearTable()
	os.Exit(code)
}

func ensureTableExists() {
	if _, err := testApp.DB.Exec(tableCreationQuery); err != nil {
		log.Fatal(err)
	}
}

func clearTable() {
	testApp.DB.Exec("DELETE FROM products")
	testApp.DB.Exec("ALTER SEQUENCE products_id_seq RESTART WITH 1")
}
