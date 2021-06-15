package jz_offer

// 输入一个整型数组，数组中的一个或连续多个整数组成一个子数组。求所有子数组的和的最大值。
//
//要求时间复杂度为O(n)。
func maxSubArray(nums []int) int {
	ans, curSum := nums[0], 0
	for _, num := range nums {
		curSum = max(num, curSum+num)
		ans = max(ans, curSum)
	}
	return ans
}