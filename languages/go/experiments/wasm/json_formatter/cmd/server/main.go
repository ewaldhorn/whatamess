package main
/*
 * Super basic HTTP server.
 * Only usable during development.
 * It has no security and it's full of exploits, for sure.
 */
import (
	"fmt"
	"net/http"
)

func main() {
	err := http.ListenAndServe(":9000", http.FileServer(http.Dir("../../assets")))
	if err != nil {
		fmt.Println("Failed to start server", err)
		return
	}
}
