package main

import (
	"log"
	"os"
	"strings"
)

const expectedWordLength = 5

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
	if len(content) > 4 {
		entries := strings.Split(string(content), "\n")

		for _, e := range entries {
			tmp := strings.TrimSpace(e)
			if !strings.Contains(tmp, "--") && len(tmp) == expectedWordLength {
				dictionary = append(dictionary, strings.ToUpper(tmp))
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
