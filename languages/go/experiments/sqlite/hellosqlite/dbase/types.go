package dbase

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// ----------------------------------------------------------------------------
type Agent struct {
	id     int
	name   string
	time   string
	email  string
	score  int
	region int
}

// ----------------------------------------------------------------------------
func ReadNewAgentFromCLI() (*Agent, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter a name: ")
	name, err := reader.ReadString('\n')
	if err != nil {
		return nil, err
	}

	fmt.Print("Enter a email: ")
	email, err := reader.ReadString('\n')
	if err != nil {
		return nil, err
	}

	newAgent := Agent{
		name:  strings.TrimSpace(name),
		email: strings.TrimSpace(email),
		score: 0,
	}

	return &newAgent, nil
}
