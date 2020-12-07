package numsquares

func numSquares(n int) int {
	if n == 0 {
		return 0
	}

	dp := make([]int, n)
	for i := 0; i < n; i++ {
		num := i + 1
		dp[i] = num
		for j := 1; j*j <= num; j++ {
			if num == j*j {
				dp[i] = 1
			} else {
				dp[i] = min(dp[i-j*j]+1, dp[i])
			}
		}
	}
	return dp[n-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
