70. 爬楼梯梯

## 题目
假设你正在爬楼梯。需要 n 阶你才能到达楼顶。  

每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？  
注意：给定 n 是一个正整数。

**示例 1** 

```
输入： 2    
输出： 2  
解释： 有两种方法可以爬到楼顶。  
1.  1 阶 + 1 阶  
2.  2 阶  
``` 

**示例 2** 

```
输入： 3
输出： 3
解释： 有三种方法可以爬到楼顶。
1.  1 阶 + 1 阶 + 1 阶
2.  1 阶 + 2 阶
3.  2 阶 + 1 阶
```
> 转自力扣（LeetCode）

## 题解
每次可以爬1或2两个台阶，所以到达第n层可以通过第n-1层和n-2层两种方式，归纳出来的动态转移方程如下：

```
f(n) = f(n-1) + f(n-2)
```