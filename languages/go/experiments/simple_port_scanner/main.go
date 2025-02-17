package main

import (
	"fmt"
	"net"
	"sync"
	"time"
)

// ----------------------------------------------------------------------------
func scanPort(port int, wg *sync.WaitGroup) {
	defer wg.Done()

	address := fmt.Sprintf("localhost:%d", port)
	conn, err := net.DialTimeout("tcp", address, 1*time.Second)

	if err != nil {
		fmt.Printf("Port %3d errored: %v\n", port, err.Error())
		return
	}

	conn.Close()

	fmt.Printf("Port %d is open\n", port)
}

// ----------------------------------------------------------------------------
func main() {
	var wg sync.WaitGroup

	for port := 1; port <= 1024; port++ {
		wg.Add(1)

		go scanPort(port, &wg)
	}

	wg.Wait()
}
