package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	fmt.Println("Go Web Assembly")
}

const PREFIX = ""

func prettyJson(input string) (string, error) {
	var raw any
	if err := json.Unmarshal([]byte(input), &raw); err != nil {
		return "", err
	}
	
	pretty, err := json.MarshalIndent(raw, PREFIX, "  ")
	if err != nil {
		return "", err
	}

	return string(pretty), nil
}
