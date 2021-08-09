package dp

import "math"

// 超级丑数 是一个正整数，并满足其所有质因数都出现在质数数组 primes 中。
//
//给你一个整数 n 和一个整数数组 primes ，返回第 n 个 超级丑数 。
//
//题目数据保证第 n 个 超级丑数 在 32-bit 带符号整数范围内。
//
func nthSuperUglyNumber(n int, primes []int) int {
	dp := make([]int, n+1)
	dp[1] = 1
	m := len(primes)
	pointers := make([]int, m)
	for i := range pointers {
		pointers[i] = 1
	}

	for i := 2; i <= n; i++ {
		nums := make([]int, m)
		minNum := math.MaxInt64
		for j, p := range pointers {
			nums[j] = dp[p] * primes[j]
			minNum = min(minNum, nums[j])
		}
		dp[i] = minNum
		for j, num := range nums {
			if minNum == num {
				pointers[j]++
			}
		}
	}
	return dp[n]
}
