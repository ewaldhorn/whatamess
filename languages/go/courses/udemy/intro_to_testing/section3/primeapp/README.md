# Testing in Go

Run with `go test .`

Can also use `go test -v .` for more details (verbose).

For code coverage `go test -cover .` prints out test coverage percentage.

For a code coverage report, first run `go test -coverprofile=coverage.out .`. Once you have a
coverage report, `go tool cover -html=coverage.out` generates a viewable coverage report.
