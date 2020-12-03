# 53. 最大子序和
## 题目
给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。 

**示例:**  

```
输入: [-2,1,-3,4,-1,2,1,-5,4]
输出: 6
解释: 连续子数组 [4,-1,2,1] 的和最大，为 6。
```

**进阶:**  
如果你已经实现复杂度为 O(n) 的解法，尝试使用更为精妙的分治法求解。

## 题解
用一个数组`dp[i]`来记录每一个记录`nums[i]`对应的连续子数组的最大和，`dp[i]`中的值是取`dp[i-1]+nums[i]`和`nums[i]`中的较大的值，对应的动态转移方程：`f(i) = max{f(i-1)+nums[i], nums[i]}`。

## 代码

```go
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	max := dp[0]

	for i := 1; i < len(nums); i++ {
		if dp[i-1]+nums[i] > nums[i] {
			dp[i] = dp[i-1] + nums[i]
		} else {
			dp[i] = nums[i]
		}
		if dp[i] > max {
			max = dp[i]
		}
	}
	return max
}
```


