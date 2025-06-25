package main

import (
	"fmt"
	"net/url"
)

type QueryParam struct {
	Key   string
	Value *string
}

// BuildQueryURL constructs a URL with dynamic query parameters.
// It takes a baseURL and a slice of QueryParam structs.
// Only parameters with non-nil values will be added to the query string.
func BuildQueryURL(baseURL string, params []QueryParam) string {
	tmpURL, err := url.Parse(baseURL)
	if err != nil {
		panic(fmt.Sprintf("Error parsing base URL: %v\n", err))
	}

	query := tmpURL.Query()

	for _, param := range params {
		if param.Value != nil {
			query.Set(param.Key, *param.Value)
		}
	}

	tmpURL.RawQuery = query.Encode()

	return tmpURL.String()
}

func main() {
	baseURL := "http://example.com/api"

	// Example usage 1: Both personID and productID are present
	personID1 := "user123"
	productID1 := "brick456"
	params1 := []QueryParam{
		{Key: "personId", Value: &personID1},
		{Key: "productId", Value: &productID1},
		{Key: "status", Value: nil}, // This will be ignored
	}
	url1 := BuildQueryURL(baseURL, params1)
	fmt.Printf("URL 1: %s\n", url1)
	// Expected output: http://example.com/api?personId=user123&productId=conv456

	// Example usage 2: Only personID is present
	personID2 := "user789"
	var productID2 *string // nil
	params2 := []QueryParam{
		{Key: "personId", Value: &personID2},
		{Key: "productId", Value: productID2}, // This will be ignored
	}
	url2 := BuildQueryURL(baseURL, params2)
	fmt.Printf("URL 2: %s\n", url2)
	// Expected output: http://example.com/api?personId=user789

	// Example usage 3: No optional parameters
	params3 := []QueryParam{
		{Key: "page", Value: nil},
		{Key: "limit", Value: nil},
	}
	url3 := BuildQueryURL(baseURL, params3)
	fmt.Printf("URL 3: %s\n", url3)
	// Expected output: http://example.com/api

	// Example usage 4: Multiple additional parameters
	personID4 := "abc"
	name := "Alice"
	age := "30"
	params4 := []QueryParam{
		{Key: "personId", Value: &personID4},
		{Key: "name", Value: &name},
		{Key: "age", Value: &age},
		{Key: "city", Value: nil},
	}
	url4 := BuildQueryURL(baseURL, params4)
	fmt.Printf("URL 4: %s\n", url4)
	// Expected output: http://example.com/api?age=30&personId=abc&name=Alice
}
