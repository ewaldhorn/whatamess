package main

// --------------- -------------------------------------------------------------
func findLCS(left, right string) string {
	// short-circuit if we don't actually need to do any work
	if left == "" || right == "" {
		return ""
	}

	leftLen := len(left)
	rightLen := len(right)

	// Create a 2D array to store the lengths of common subsequences.
	dp := make([][]int, leftLen+1)
	for i := range dp {
		dp[i] = make([]int, rightLen+1)
	}

	// Fill the dp array in a bottom-up manner.
	for i := 1; i <= leftLen; i++ {
		for j := 1; j <= rightLen; j++ {
			if left[i-1] == right[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	// Reconstruct the LCS from the dp array.
	lcs := make([]byte, 0)
	i, j := leftLen, rightLen
	for i > 0 && j > 0 {
		if left[i-1] == right[j-1] {
			lcs = append(lcs, left[i-1])
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

	return string(lcs)
}

// ----------------------------------------------------------------------------
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
