package writers

import (
	"bytes"
	"fmt"
)

func SimpleWriter() {
	buffer := bytes.NewBufferString("")
	numBytes, err := buffer.WriteString("SAMPLE STRING")

	if err != nil {
		fmt.Println("Something went wrong!", err)
	} else {
		fmt.Printf("Wrote %d bytes: %c\n", numBytes, buffer)
	}
}
