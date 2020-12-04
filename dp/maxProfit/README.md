# 121. 买卖股票的最佳时机机

## 题目
给定一个数组，它的第 i 个元素是一支给定股票第 i 天的价格。

如果你最多只允许完成一笔交易（即买入和卖出一支股票一次），设计一个算法来计算你所能获取的最大利润。

**注意：**你不能在买入股票前卖出股票。

 

**示例 1:**

```
输入: [7,1,5,3,6,4]
输出: 5
解释: 在第 2 天（股票价格 = 1）的时候买入，在第 5 天（股票价格 = 6）的时候卖出，最大利润 = 6-1 = 5 。
     注意利润不能是 7-1 = 6, 因为卖出价格需要大于买入价格；同时，你不能在买入前卖出股票。
```
**示例 2:**

```
输入: [7,6,4,3,1]
输出: 0
解释: 在这种情况下, 没有交易完成, 所以最大利润为 0。
```

> 转自力扣（LeetCode）

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