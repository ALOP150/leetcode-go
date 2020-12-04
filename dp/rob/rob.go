package rob

func rob(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}

	dp := make([]int, n)
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])

	res := max(dp[0], dp[1])
	for i := 2; i < n; i++ {
		dp[i] = max(dp[i-2]+nums[i], dp[i-1])
		if dp[i] > res {
			res = dp[i]
		}
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
