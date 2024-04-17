package main

import (
	"net"
	"sync"
)

// --------------------------------------------------------------------------------------------------
// const declarations
const VATRate = 0.15
const identityCode = 1010

// --------------------------------------------------------------------------------------------------
// grouping variables
var (
	runningTotal             = 0
	startingPosition float64 = 1.0
	endOfRange               = 100
)

// --------------------------------------------------------------------------------------------------
// functions that error
func parseUserInput(input string) (int, error) {
	value := 0
	_ = input
	return value, nil
}

// --------------------------------------------------------------------------------------------------
// functions that panic
func MustParseOrPanic(input string) int {
	value := 0

	if input == "doh" {
		panic("oops - doh!")
	}

	return value
}

// --------------------------------------------------------------------------------------------------
// structure initialisation
type EntityMovement struct {
	x, y     int
	velocity float64
}

// --------------------------------------------------------------------------------------------------
// mutexes
type Server struct {
	address     string
	isAvailable bool

	clientsMutex sync.RWMutex
	clients      map[string]net.Conn
}

// --------------------------------------------------------------------------------------------------
// constructor functions
func NewServer() *Server {
	return &Server{address: "localhost"}
}

//--------------------------------------------------------------------------------------------------
//--------------------------------------------------------------------------------------------------
//--------------------------------------------------------------------------------------------------
//--------------------------------------------------------------------------------------------------

// --------------------------------------------------------------------------------------------------
func main() {
	entity := EntityMovement{
		x: 10, y: 11, velocity: 1.05,
	}
	_ = entity
}
