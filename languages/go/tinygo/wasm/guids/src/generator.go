package main

import (
	"crypto/rand"
	"fmt"
)

// ----------------------------------------------------------------------------
func generateGUID() (string, error) {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	// Set version (4) and variant (RFC 4122) bits
	b[6] = (b[6] & 0x0F) | 0x40 // Version 4
	b[8] = (b[8] & 0x3F) | 0x80 // RFC 4122 variant

	return fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:]), nil
}

// ----------------------------------------------------------------------------
func print_five() {
	for i := 0; i < 5; i++ {
		guid, err := generateGUID()
		if err != nil {
			fmt.Println("Error generating GUID:", err)
			return
		}
		fmt.Println(guid)
	}
}
