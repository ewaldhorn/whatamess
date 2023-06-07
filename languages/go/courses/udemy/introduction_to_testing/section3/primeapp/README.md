# Basic testing in Go

A function that checks whether a number is prime or not.

run with `go run .`
test with `go test .`
get coverage percentage with `go test -cover .`
get coverage information with `go test -coverprofile=coverage.out .`
to use coverage information `go tool cover -html=coverage.out`
to run just a particular test `go test -run Test_alpha_isPrime`
to run a related set of tests `go test -run Test_alpha`
