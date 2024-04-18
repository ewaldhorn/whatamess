package main

import (
	"log"
	"os"
	"strings"
)

// -------------------------------------------------------------------------------------------------
func loadDictionary() {
	content, err := os.ReadFile(dictionaryFile)
	if err != nil {
		log.Fatal("unable to load dictionary:", err)
	}
	entries := strings.Split(string(content), "\n")

	for _, e := range entries {
		if !strings.Contains(e, "--") {
			dictionary = append(dictionary, e)
		}
	}
}

// -------------------------------------------------------------------------------------------------
func printDictionary() {
	println("Dictionary:")
	for _, e := range dictionary {
		println(e)
	}
	println()
}
