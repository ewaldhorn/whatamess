package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", ":2007")
	if err != nil {
		fmt.Println("error with listener:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Listening on :2007")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("error accepting incoming:", err)
			return
		}
		fmt.Println("Client connection from ", conn.RemoteAddr())

		go HandleClient(conn, "reference.txt")
	}
}

func HandleClient(conn net.Conn, path string) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("error opening source file:", err)
		return
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Fprintln(conn, line)
		time.Sleep(750 * time.Millisecond)
	}

	conn.Close()
	fmt.Println("Client disconnected from ", conn.RemoteAddr())
	file.Close()
}
