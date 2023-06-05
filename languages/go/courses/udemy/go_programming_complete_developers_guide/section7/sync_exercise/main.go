package main

import (
	"fmt"
	"strings"
	"sync"
	"unicode"
)

type Count struct {
	count int
	sync.Mutex
}

func getWords(line string) []string {
	return strings.Split(line, " ")
}

func countLetters(word string) int {
	letters := 0

	for _, v := range word {
		if unicode.IsLetter(v) {
			letters += 1
		}
	}

	return letters
}

func main() {
	fmt.Printf("\nSynchronization Exercise.\n")
	inputStrings := []string{"The quick brown fox", "jumped over the lazy dog."}

	totalLetters := Count{}

	var wg sync.WaitGroup

	for _, s := range inputStrings {
		words := getWords(s)
		for _, word := range words {
			wordCopy := word
			wg.Add(1)
			go func() {
				totalLetters.Lock()
				defer totalLetters.Unlock()
				defer wg.Done()
				totalLetters.count += countLetters(wordCopy)
			}()
		}
	}

	wg.Wait()
	totalLetters.Lock()
	fmt.Println("Total letters:", totalLetters.count)
	totalLetters.Unlock()

	fmt.Println()
}
