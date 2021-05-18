package array

import "sort"

//给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置。
//
//如果数组中不存在目标值 target，返回 [-1, -1]。
//
//进阶：你可以设计并实现时间复杂度为 O(log n) 的算法解决此问题吗？
func searchRange(nums []int, target int) []int {
	insert := sort.SearchInts(nums, target)
	if insert == len(nums) || nums[insert] != target {
		return []int{-1, -1}
	}
	return []int{insert, sort.SearchInts(nums, target+1) - 1}
}
