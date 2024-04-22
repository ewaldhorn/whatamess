package main

import (
	"log"
	"os"
	"strings"
)

const expectedWordLength = 5
const hardcoded_dictionary = `-- Five letter words
mains
email
lists
small
snail
plead
pleat
plume
barks
staff
stare
stale
femur
plank
pluck
prune
pupil
spark
teale
-- Just a comment
tease
flunk
heart
snuck
sneak
snort
beard
-- This next one is too short
peek
-- This one too
    four
spoon
dryer
drain
paper
brain
sleet
-- Another comment
smart
sneak
tulip
tripe
`

// -------------------------------------------------------------------------------------------------
func loadDictionaryFromDisk() []byte {
	content, err := os.ReadFile(dictionaryFile)

	if err != nil {
		log.Print("unable to load dictionary (might be on wasm):", err)
		content = []byte(hardcoded_dictionary)
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
