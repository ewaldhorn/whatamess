package main

// ----------------------------------------------------------------------------
type Product struct {
	Name string
	Description string
	Count int
}

// ----------------------------------------------------------------------------
func NewProduct(n,d string, c int)*Product{
	return &Product{Name:n, Description:d, Count:c}
}
