# 279. 完全平方数

## 题目

给定正整数 n，找到若干个完全平方数（比如 1, 4, 9, 16, ...）使得它们的和等于 n。你需要让组成和的完全平方数的个数最少。

**示例 1:**

```
输入: n = 12
输出: 3
解释: 12 = 4 + 4 + 4.
```

**示例 2:**

```
输入: n = 13
输出: 2
解释: 13 = 4 + 9.
```

## 题解

**动态规划**

利用一个数组dp来存储i对应的完全平方数`1<i<=n`, 另起一个循环来遍历平方数`j*j <= i+1`，如果`j*j == i+1`, `dp[i]=1`, 否则`dp[i] = min(dp[i-j*j], dp[i])+1`

```go
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
```