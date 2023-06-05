package main

import (
	"fmt"
	"time"
)

type Title string
type Name string

type LendAudit struct {
	checkOut time.Time
	checkIn  time.Time
}

type Member struct {
	name  Name
	books map[Title]LendAudit
}

type BookEntry struct {
	total int
	lent  int
}

type Library struct {
	member map[Name]Member
	books  map[Title]BookEntry
}

func printMemberAudit(member *Member) {
	for title, audit := range member.books {
		var returnTime string
		if audit.checkIn.IsZero() {
			returnTime = "[not returned yet]"
		} else {
			returnTime = audit.checkIn.String()
		}
		fmt.Println(member.name, ":", title, ":", audit.checkOut.String(), "through", returnTime)
	}
}

func printMemberAudits(library *Library) {
	for _, member := range library.member {
		printMemberAudit(&member)
	}
}

func printLibraryBooks(library *Library) {
	fmt.Println()
	for title, book := range library.books {
		fmt.Println(title, "/ total:", book.total, "/ lended:", book.lent)
	}
	fmt.Println()
}

func checkoutBook(library *Library, title Title, member *Member) bool {
	book, found := library.books[title]

	if !found {
		fmt.Println("Library does not have a book titled: ", title)
		return false
	}

	if book.lent == book.total {
		fmt.Println("The library has no more stock of the book at the moment!")
		return false
	}

	book.lent += 1
	library.books[title] = book
	member.books[title] = LendAudit{checkOut: time.Now()}
	return true
}

func returnBook(library *Library, title Title, member *Member) bool {
	book, found := library.books[title]

	if !found {
		fmt.Println("This is not a book from this library!")
		return false
	}

	audit, found := member.books[title]

	if !found {
		fmt.Println("This member didn't lend this book!")
		return false
	}

	book.lent -= 1
	library.books[title] = book
	audit.checkIn = time.Now()
	member.books[title] = audit
	return true
}

func main() {
	// create our library
	library := Library{
		books:  make(map[Title]BookEntry),
		member: map[Name]Member{},
	}

	// add some books
	library.books["Java for Dummies"] = BookEntry{total: 3, lent: 0}
	library.books["JavaScript 1"] = BookEntry{total: 1, lent: 0}
	library.books["Node"] = BookEntry{total: 1, lent: 0}
	library.books["Kotlin"] = BookEntry{total: 5, lent: 5}

	// add some members
	library.member["Jim"] = Member{name: "Jim", books: map[Title]LendAudit{}}
	library.member["Bob"] = Member{name: "Bob", books: map[Title]LendAudit{}}
	library.member["Ann"] = Member{name: "Ann", books: map[Title]LendAudit{}}

	fmt.Println("\nInitial Status of Library:")
	printLibraryBooks(&library)
	printMemberAudits(&library)

	member := library.member["Jim"]
	checkedOut := checkoutBook(&library, "Node", &member)
	fmt.Println(member.name, "checked out book:", checkedOut)

	fmt.Println("\nNew Status of Library:")
	printLibraryBooks(&library)
	printMemberAudits(&library)

	returned := returnBook(&library, "Node", &member)
	fmt.Println(member.name, "returned book:", returned)

	fmt.Println("\nNew Status of Library:")
	printLibraryBooks(&library)
	printMemberAudits(&library)

}
