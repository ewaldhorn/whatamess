package main

// need a custom type for the reducer/folding function
type foldFunction[E any] func(E, E) E

// custom reducer, aka fold in scala
func Fold[E any](slice []E, initial E, folder foldFunction[E]) E {
	// get the initial value
	current := initial

	// perform the fold operation(s)
	for _, value := range slice {
		current = folder(current, value)
	}

	// return the end result
	return current
}
