package main

import "fmt"

// ----------------------------------------------------------------------------
func longestCommonSubsequence(X string, Y string, m int, n int) int {
	if m == 0 || n == 0 {
		return 0
	}

	if X[m-1] == Y[n-1] {
		return 1 + longestCommonSubsequence(X, Y, m-1, n-1)
	}

	return max(longestCommonSubsequence(X, Y, m, n-1), longestCommonSubsequence(X, Y, m-1, n))
}

// ----------------------------------------------------------------------------
func main() {
	X := "AGGTAB"
	Y := "GXTXAYB"
	m := len(X)
	n := len(Y)

	fmt.Println("Length of LCS is", longestCommonSubsequence(X, Y, m, n))
}
