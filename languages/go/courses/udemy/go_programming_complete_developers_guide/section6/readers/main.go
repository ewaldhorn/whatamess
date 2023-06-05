package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Hit Ctrl-Z on Windows or Ctrl-D on Mac to stop input
// input is expected to be something like:  1 2 3 4 5
// which will yield 15
func main() {
	r := bufio.NewReader(os.Stdin)

	sum := 0

	for {
		inp, inpErr := r.ReadString(' ')
		n := strings.TrimSpace(inp)
		if n == "" {
			continue
		}
		num, convErr := strconv.Atoi(n)
		if convErr != nil {
			fmt.Println(convErr)
		} else {
			sum += num
		}

		if inpErr == io.EOF {
			break
		} else if inpErr != nil {
			fmt.Println(inpErr)
		}
	}

	fmt.Println("The total is:", sum)
}
