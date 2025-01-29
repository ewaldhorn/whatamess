package main

import "fmt"

// ----------------------------------------------------------------------------
func withMaps() {
	products := map[string]float32{}

	displayProducts_map(products)

	addProduct_map("macaroni", 10.05, products)
	displayProducts_map(products)

	addProduct_map("cheese", 41.15, products)
	addProduct_map("eggs", 18.85, products)
	addProduct_map("olive oil", 38.95, products)
	displayProducts_map(products)
}

// ----------------------------------------------------------------------------
func displayProducts_map(p map[string]float32) {
	fmt.Printf("\nThere are %d products:\n", len(p))

	for prod, price := range p {
		fmt.Printf("%20s cost R%.2f\n", prod, price)
	}
}

// ----------------------------------------------------------------------------
func addProduct_map(name string, price float32, p map[string]float32) {
	p[name] = price
}
