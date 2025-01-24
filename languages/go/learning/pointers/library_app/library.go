package main

import "errors"

// ----------------------------------------------------------------------------
type Library struct {
	name    string
	address string
	books   []*Book
}

// ----------------------------------------------------------------------------
func NewLibrary(name string, address string) (*Library, error) {
	if name == "" || address == "" {
		return nil, errors.New("library needs a name and address")
	}

	return &Library{name: name, address: address, books: []*Book{}}, nil
}

// ----------------------------------------------------------------------------
func (l *Library) addBook(b Book) {
	l.books = append(l.books, &b)
}
