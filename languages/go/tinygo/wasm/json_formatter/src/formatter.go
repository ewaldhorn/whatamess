package main

import "encoding/json"

// ----------------------------------------------------------------------------
func prettifyJSON(inputJSON string) (string, error) {
	var raw any

	if err := json.Unmarshal([]byte(inputJSON), &raw); err != nil {
		return "", err
	}

	output, err := json.MarshalIndent(raw, "", "  ")
	if err != nil {
		return "", err
	}

	return string(output), nil
}
