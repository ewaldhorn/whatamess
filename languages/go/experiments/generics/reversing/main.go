package main

func reverse[E any](collection []E) []E {
	// make a new collection of the same type and size
	result := make([]E, 0, len(collection))

	// populate the new collection in reverse order
	for i := len(collection) - 1; i >= 0; i-- {
		result = append(result, collection[i])
	}

	// now return it
	return result
}
