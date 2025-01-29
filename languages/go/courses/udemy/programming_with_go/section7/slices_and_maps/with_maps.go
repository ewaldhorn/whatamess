package main

import "fmt"

// ----------------------------------------------------------------------------
func withMaps() {
	products := map[string]float32{}

	displayProducts(products)
	addProduct("macaroni", 10.05, products)
	displayProducts(products)
	addProduct("cheese", 41.15, products)
	addProduct("eggs", 18.85, products)
	addProduct("olive oil", 38.95, products)
	displayProducts(products)
}

// ----------------------------------------------------------------------------
func displayProducts(p map[string]float32) {
	fmt.Printf("\nThere are %d products:\n", len(p))

	for prod, price := range p {
		fmt.Printf("%20s cost R%.2f\n", prod, price)
	}
}

// ----------------------------------------------------------------------------
func addProduct(name string, price float32, p map[string]float32) {
	p[name] = price
}
