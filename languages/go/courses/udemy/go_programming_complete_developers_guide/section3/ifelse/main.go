package main

import "fmt"

const (
	Monday    = 0
	Tuesday   = 1
	Wednesday = 2
	Thursday  = 3
	Friday    = 4
	Saturday  = 5
	Sunday    = 6
)

const (
	Admin      = 10
	Manager    = 20
	Contractor = 30
	Member     = 40
	Guest      = 50
)

func accessGranted() {
	fmt.Println("Granted")
}

func accessDenied() {
	fmt.Println("Denied")
}

func isWeekday(day int) bool {
	return day <= 4
}

func main() {
	today, role := Monday, Guest

	if role == Admin || role == Manager {
		accessGranted()
	} else if role == Contractor && !isWeekday(today) {
		accessGranted()
	} else if role == Member && isWeekday(today) {
		accessGranted()
	} else if role == Guest && (today == Monday || today == Wednesday) {
		accessGranted()
	} else {
		accessDenied()
	}
}
