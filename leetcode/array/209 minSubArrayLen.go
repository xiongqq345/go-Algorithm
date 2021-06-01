package array

import "math"

// 给定一个含有n个正整数的数组和一个正整数 target 。
//
//找出该数组中满足其和 ≥ target 的长度最小的 连续子数组[numsl, numsl+1, ..., numsr-1, numsr] ，并返回其长度。如果不存在符合条件的子数组，返回 0 。

func minSubArrayLen(target int, nums []int) int {
	l, sum, ans := 0, 0, math.MaxInt32
	for r := 0; r < len(nums); r++ {
		sum += nums[r]
		for sum >= target {
			ans = min(ans, r-l+1)
			sum -= nums[l]
			l++
		}
	}
	if ans == math.MaxInt32 {
		return 0
	}
	return ans
}
