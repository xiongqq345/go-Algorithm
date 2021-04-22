package _98

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	dp := make([]int, 2)
	dp[0], dp[1] = nums[0], max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		dp[0], dp[1] = dp[1], max(dp[0]+nums[i], dp[1])
	}
	return dp[1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
