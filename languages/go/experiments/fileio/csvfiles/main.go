package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const FILENAME = "data.csv"

func main() {
	file, err := os.Open(FILENAME)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		items := strings.Split(line, ",")
		fmt.Printf("Name: %s %s \nEmail: %s\nIP Address: %s\n", items[1], items[2], items[3], items[5])
		fmt.Println(strings.Repeat("-", 20))
	}
}
