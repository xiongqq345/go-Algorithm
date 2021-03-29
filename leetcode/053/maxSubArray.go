package leetcode

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	lens, max := len(nums), nums[0]
	dp := make([]int, lens)
	dp[0] = max
	for i := 1; i < lens; i++ {
		if dp[i-1] > 0 {
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
