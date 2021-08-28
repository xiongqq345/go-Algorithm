package array

// 给你一个数组 nums 。数组「动态和」的计算公式为：runningSum[i] = sum(nums[0]…nums[i]) 。
//
//请返回 nums 的动态和。
func runningSum(nums []int) []int {
	for i := range nums[1:] {
		nums[i+1] += nums[i]
	}
	return nums
}
