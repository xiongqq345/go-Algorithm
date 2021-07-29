package list_stack

import "math"

// 给你一个整数数组 nums ，数组中共有 n 个整数。132 模式的子序列 由三个整数 nums[i]、nums[j] 和 nums[k] 组成，并同时满足：i < j < k 和 nums[i] < nums[k] < nums[j] 。
//
//如果 nums 中存在 132 模式的子序列 ，返回 true ；否则，返回 false 。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/132-pattern
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func find132pattern(nums []int) bool {
	stack := []int{nums[len(nums)-1]}
	maxK := math.MinInt64
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] < maxK {
			return true
		}
		for len(stack) > 0 && nums[i] > stack[len(stack)-1] {
			maxK = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
		if nums[i] > maxK {
			stack = append(stack, nums[i])
		}
	}
	return false
}
