package dp

// 你有一个整数数组 nums。你只能将一个元素 nums[i] 替换为 nums[i] * nums[i]。
//
// 返回替换后的最大子数组和。
func maxSumAfterOperation(nums []int) int {
	var dp0, dp1 int
	var ans int
	for _, num := range nums {
		dp1 = max(dp0+num*num, max(dp1+num, num*num))
		dp0 = max(num, dp0+num)
		ans = max(ans, dp1)
	}
	return ans
}
