package dp

import "math"

// 给定不同面额的硬币 coins 和一个总金额 amount。编写一个函数来计算可以凑成总金额所需的最少的硬币个数。如果没有任何一种硬币组合能组成总金额，返回-1。
//
//你可以认为每种硬币的数量是无限的。

func coinChange(coins []int, amount int) int {
	f := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		f[i] = math.MaxInt32
		for _, v := range coins {
			if i >= v && f[i] > f[i-v]+1 {
				f[i] = f[i-v] + 1
			}
		}
	}
	if f[amount] == math.MaxInt32 {
		return -1
	}
	return f[amount]
}
