package jz_offerII

// 给定一个含有 n 个正整数的数组和一个正整数 target 。
//
//找出该数组中满足其和 ≥ target 的长度最小的 连续子数组 [numsl, numsl+1, ..., numsr-1, numsr] ，并返回其长度。如果不存在符合条件的子数组，返回 0 。
//
func minSubArrayLen(target int, nums []int) int {
	var sum int
	ans := len(nums) + 1
	for i, j := 0, 0; j < len(nums); j++ {
		sum += nums[j]
		for sum >= target {
			ans = min(ans, j-i+1)
			sum -= nums[i]
			i++
		}
	}
	if ans == len(nums)+1 {
		return 0
	}
	return ans
}
