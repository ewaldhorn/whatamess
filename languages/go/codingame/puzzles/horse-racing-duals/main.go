package main

import "fmt"
import "math"
import "sort"

func main() {
	var N int
	fmt.Scan(&N)

	horses := make([]int, N)

	for i := 0; i < N; i++ {
		var pi int
		fmt.Scan(&pi)
		horses[i] = pi
	}

	answer := horses[0]
	sort.Ints(horses)

	for i := 1; i < N; i++ {
		tmp := int(math.Abs(float64(horses[i] - horses[i-1])))
		if tmp < answer {
			answer = tmp
		}
	}

	fmt.Println(answer)
}

