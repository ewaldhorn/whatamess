package main

import (
	"fmt"
	"interpreter/src/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()

	if err != nil {
		panic(err)
	}

	fmt.Printf("Howdy %s. This is the %s v%s.\nTry out some commands! (\\h - help, \\q - quit.)\n", user.Username, NAME, VERSION)
	repl.Start(os.Stdin, os.Stdout)
}
