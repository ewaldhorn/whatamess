package main

import (
	"log"
	"os"
	"strings"
)

// -------------------------------------------------------------------------------------------------
func loadDictionaryFromDisk() []byte {
	content, err := os.ReadFile(dictionaryFile)

	if err != nil {
		log.Fatal("unable to load dictionary:", err)
	}

	return content
}

// -------------------------------------------------------------------------------------------------
func parseDictionary(content []byte) []string {
	// we need at least one word to bother unpacking things
	if len(content) > 5 {
		entries := strings.Split(string(content), "\n")

		for _, e := range entries {
			if !strings.Contains(e, "--") {
				dictionary = append(dictionary, e)
			}
		}
	}
	return dictionary
}

// -------------------------------------------------------------------------------------------------
func printDictionary() {
	println("Dictionary:")
	for _, e := range dictionary {
		println(e)
	}
	println()
}
