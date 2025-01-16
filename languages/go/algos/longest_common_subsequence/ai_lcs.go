package main

import "errors"

// LCS returns the longest common subsequence of two strings.
func LCS(s1, s2 string) (string, error) {
	if len(s1) == 0 || len(s2) == 0 {
		return "", errors.New("input strings cannot be empty")
	}

	m := len(s1)
	n := len(s2)

	// Create a 2D array to store the lengths of common subsequences.
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	// Fill the dp array in a bottom-up manner.
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if s1[i-1] == s2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	// Reconstruct the LCS from the dp array.
	lcs := make([]byte, 0)
	i, j := m, n
	for i > 0 && j > 0 {
		if s1[i-1] == s2[j-1] {
			lcs = append(lcs, s1[i-1])
			i--
			j--
		} else if dp[i-1][j] > dp[i][j-1] {
			i--
		} else {
			j--
		}
	}

	// Reverse the LCS since we constructed it in reverse order.
	for i, j := 0, len(lcs)-1; i < j; i, j = i+1, j-1 {
		lcs[i], lcs[j] = lcs[j], lcs[i]
	}

	return string(lcs), nil
}
