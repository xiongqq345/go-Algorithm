package jz_offer

import "sort"

// 统计一个数字在排序数组中出现的次数。
func search(nums []int, target int) int {
	return sort.SearchInts(nums, target+1) - sort.SearchInts(nums, target)
}
