package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func main() {
	// 1. Read the JSON file
	filePath := "input.json"
	jsonData, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	// 2. Unmarshal the JSON into a generic interface{}
	// This will typically become a map[string]interface{} for JSON objects
	// or a []interface{} for JSON arrays at the top level.
	var result interface{}
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		log.Fatalf("Error unmarshaling JSON: %v", err)
	}

	// 3. Print the unmarshaled data
	// To print it nicely, you can re-marshal it with indentation
	prettyJSON, err := json.MarshalIndent(result, "", "  ") // Changed non-breaking space to regular space for clarity
	if err != nil {
		log.Fatalf("Error re-marshalling JSON for pretty print: %v", err)
	}

	fmt.Println("Original JSON (read from file):")
	fmt.Println(string(jsonData))
	fmt.Println("\nParsed JSON (re-printed with indentation):")
	fmt.Println(string(prettyJSON))

	// 4. (Optional) Accessing data dynamically using type assertions
	fmt.Println("\nAccessing data dynamically:")
	if m, ok := result.(map[string]interface{}); ok {
		if name, ok := m["name"].(string); ok {
			fmt.Printf("Name: %s\n", name)
		}

		if age, ok := m["age"].(float64); ok { // JSON numbers default to float64
			fmt.Printf("Age: %.0f\n", age) // Print as integer
		}

		if courses, ok := m["courses"].([]interface{}); ok {
			fmt.Println("Courses:")
			for i, course := range courses {
				if c, ok := course.(string); ok {
					fmt.Printf("  - %d: %s\n", i+1, c) // Changed non-breaking space to regular space
				}
			}
		}

		if address, ok := m["address"].(map[string]interface{}); ok {
			if city, ok := address["city"].(string); ok {
				fmt.Printf("City: %s\n", city)
			}
		}

		// Handling a nested array of objects
		if scores, ok := m["scores"].([]interface{}); ok {
			fmt.Println("Scores:")
			for _, scoreEntry := range scores {
				if s, ok := scoreEntry.(map[string]interface{}); ok {
					if subject, ok := s["subject"].(string); ok {
						if grade, ok := s["grade"].(float64); ok {
							fmt.Printf("  - Subject: %s, Grade: %.0f\n", subject, grade) // Changed non-breaking space to regular space
						}
					}
				}
			}
		}

		if metadata := m["metadata"]; metadata == nil {
			fmt.Println("Metadata: nil (JSON null)")
		}

	} else {
		fmt.Println("Top-level JSON is not an object (e.g., it's an array or a primitive).")
	}
}
