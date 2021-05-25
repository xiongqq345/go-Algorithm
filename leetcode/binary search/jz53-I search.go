package binary_search

import "sort"

//统计一个数字在排序数组中出现的次数。
func search(nums []int, target int) int {
	l := helper(nums, 0, len(nums)-1, target)
	r := helper(nums, l, len(nums)-1, target+1)
	return r - l
}

func helper(nums []int, l, r, target int) int {
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return l
}

func search(nums []int, target int) int {
	return sort.SearchInts(nums, target+1) - sort.SearchInts(nums, target)
}
