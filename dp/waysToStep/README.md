# 面试题 08.01. 三步问题

## 题目
三步问题。有个小孩正在上楼梯，楼梯有n阶台阶，小孩一次可以上1阶、2阶或3阶。实现一种方法，计算小孩有多少种上楼梯的方式。结果可能很大，你需要对结果模1000000007。

**示例1:**

```
 输入：n = 3 
 输出：4
 说明: 有四种走法
 ```

**示例2:**

```
 输入：n = 5
 输出：13
 ```

**提示:**

+ n范围在[1, 1000000]之间

> 转自力扣（LeetCode）


## 题解

同[爬楼梯](https://github.com/ALOP150/leetcode-go/blob/master/dp/climbStairs/README.md)基本类似，区别是变成了三步，对应的状态转移方程变成: 

```
f(n) = f(n-1) + f(n-2) + f(n-3)
```

需要注意的是可能结果太大，会越界，需要对相加的结果进行取模。

```go
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
```
