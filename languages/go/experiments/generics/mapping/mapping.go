package main

// first create a new type so that we can pass a function effect to the Map fuction
type mappingFunction[E any] func(E) E

func Map[E any](collection []E, apply mappingFunction[E]) []E {
	// create a new collection with the same size
	result := make([]E, len(collection))

	// range over every entry, applying the effect function, store resulting value
	for i := range collection {
		result[i] = apply(collection[i])
	}

	// return modified collection
	return result
}
