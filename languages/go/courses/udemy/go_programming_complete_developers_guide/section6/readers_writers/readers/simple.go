package readers

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// SimpleReader reads a string in chunks, simulating a network connection stream, for example.
func SimpleReader() {
	reader := strings.NewReader("This is a sample string")

	var newString strings.Builder
	buffer := make([]byte, 4)

	for {
		numBytes, err := reader.Read(buffer)
		chunk := buffer[:numBytes]
		newString.Write(chunk)
		fmt.Printf("Read %v bytes: %c\n", numBytes, chunk)

		if err == io.EOF {
			break
		}
	}

	fmt.Printf("%v\n", newString.String())

}

func SimpleReaderUsingBUFIO() {
	source := strings.NewReader("This is a sample string")
	buffered := bufio.NewReader(source)
	newString, err := buffered.ReadString('\n')
	if err == io.EOF {
		fmt.Println(newString)
	} else {
		fmt.Println("Something went wrong!", err)
	}
}

func TestScanner() {
	scanner := bufio.NewScanner(os.Stdin)
	lines := make([]string, 5)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if scanner.Err() != nil {
		fmt.Println("Error:", scanner.Err())
	}

	fmt.Println("Lines read:", len(lines))

	for idx, line := range lines {
		fmt.Printf("Line %d: %v\n", idx+1, line)
	}
}
