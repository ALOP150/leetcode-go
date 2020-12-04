# 买卖股票的最佳时机

## 题解

### 暴力解法
计算每一个数与它之后的数的差，比较记录最大的差就是最后的结果。

```go
func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}

	max := 0
	for i := 0; i < len(prices); i++ {
		for j := i + 1; j < len(prices); j++ {
			if prices[j]-prices[i] > max {
				max = prices[j] - prices[i]
			}
		}
	}
	return max
}
```


### 动态规划
计算每一个数与它之前记录的最小的数，计算两个数之间的差，并与之前记录的最大的差做比较，保留较大的差，就是最后的结果。

```go
func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}

	max := 0
	min := prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[i] < min {
			min = prices[i]
		}

		if prices[i]-min > max {
			max = prices[i] - min
		}
	}
	return max
}
```