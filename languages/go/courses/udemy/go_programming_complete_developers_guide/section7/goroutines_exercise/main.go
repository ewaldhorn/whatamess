package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"
)

func sumFile(rd bufio.Reader) int {
	sum := 0

	for {
		line, err := rd.ReadString('\n')

		if err == io.EOF {
			return sum
		}

		if err != nil {
			fmt.Println("Error reading file:", err)
		}

		num, err := strconv.Atoi(line[:len(line)-1])

		if err != nil {
			fmt.Println("Error converting to number:", err)
		} else {
			sum += num
		}
	}
}

func main() {
	fmt.Printf("\nGoroutines Exercise\nRunning...")
	expectedTotal := 50
	files := []string{"num1.txt", "num2.txt", "num3.txt", "num4.txt", "num5.txt"}
	total := 0

	for _, v := range files {
		file, err := os.Open(v)

		if err != nil {
			fmt.Println("Could not open:", v, err)
			return
		}

		rd := bufio.NewReader(file)

		calculate := func() {
			fileSum := sumFile(*rd)
			total += fileSum
		}

		go calculate()
	}

	time.Sleep(1000 * time.Millisecond)

	fmt.Println("Done!")
	fmt.Printf("Expected total of %d versus total of %d : Success %t.\n", expectedTotal, total, total == expectedTotal)
	fmt.Println()
}
