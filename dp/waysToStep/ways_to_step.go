package waystostep

func waysToStep(n int) int {
	if n == 0 {
		return 0
	}

	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}

	if n == 3 {
		return 4
	}

	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1
	dp[2] = 2
	dp[3] = 4
	for i := 4; i < n+1; i++ {
		dp[i] = (dp[i-1] + (dp[i-2]+dp[i-3])%1000000007) % 1000000007
	}

	return dp[n]
}
