package main

// Generic "contains" function
// Returns true if the collection contains the specified value, false otherwise
func contains[E comparable](collection []E, target E) bool {
	for _, value := range collection {
		if value == target {
			return true
		}
	}
	return false
}
