package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

var dictionary = []string{}

// --------------------------------------------------------------------------------- loadDictionary
func loadDictionary() {
	input, err := os.Open("dictionary/wordlist.txt")

	if err != nil {
		fmt.Println(err)
	}

	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		dictionary = append(dictionary, scanner.Text())
	}

	input.Close()
}

// ------------------------------------------------------------------------------- reportDictionary
func reportDictionary() {
	for pos, entry := range dictionary {
		fmt.Printf("%2d: %s\n", pos+1, entry)
	}
}

// -------------------------------------------------------------------- getRandomWordFromDictionary
func getRandomWordFromDictionary() string {
	return dictionary[rand.Intn(len(dictionary))]
}
