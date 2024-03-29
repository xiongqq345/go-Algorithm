package jz_offer

import "sort"

// 一个长度为n-1的递增排序数组中的所有数字都是唯一的，并且每个数字都在范围0～n-1之内。在范围0～n-1内的n个数字中有且只有一个数字不在该数组中，请找出这个数字。
func missingNumber(nums []int) int {
	return sort.Search(len(nums), func(i int) bool {
		return nums[i] > i
	})
}
