package array

// 给你一个整数数组 nums ，请你找出数组中乘积最大的连续子数组（该子数组中至少包含一个数字），并返回该子数组所对应的乘积。

// 输入: [2,3,-2,4]
//输出: 6
//解释: 子数组 [2,3] 有最大乘积 6。
func maxProduct(nums []int) int {
	dpMax, dpMin, ans := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		mx, mn := dpMax, dpMin
		dpMax = max(max(nums[i], mx*nums[i]), mn*nums[i])
		dpMin = min(min(nums[i], mx*nums[i]), mn*nums[i])
		ans = max(dpMax, ans)
	}
	return ans
}
