# Benchmarking in Go

Take from <https://dev.to/hyperskill/testing-and-benchmarking-in-go-34en>

## To run

`go test -bench . -benchmem -count 10 > 10_runs_bench.txt`
`go test -bench . -benchmem -count 20 > 20_runs_bench.txt`

Generates memory benchmarks for 10 and 20 runs.
