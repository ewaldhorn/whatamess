package main

import "os"

func main() {
	mainApp := App{}
	mainApp.Initialize(
		os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_NAME"))

	mainApp.Run(":9000")
}
