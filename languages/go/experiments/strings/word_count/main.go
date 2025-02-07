package main

const data = `
The quick brown fox jumps over the lazy dog
`

func main() {
	count := 0
	for _, c := range data {
		if c == ' ' {
			count++
		}
	}
	count++ // for the last word

	println(count)
}
