package main

import "fmt"

type UserError struct {
	Msg string
}

func (u *UserError) Error() string {
	return fmt.Sprintf("User error: %v", string(u.Msg))
}

func testUserError() {
}
