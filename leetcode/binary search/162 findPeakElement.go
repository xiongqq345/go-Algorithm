package binary_search

import "math"

// 峰值元素是指其值大于左右相邻值的元素。
//
//给你一个输入数组nums，找到峰值元素并返回其索引。数组可能包含多个峰值，在这种情况下，返回 任何一个峰值 所在位置即可。
//
//你可以假设nums[-1] = nums[n] = -∞ 。
func findPeakElement(nums []int) int {
	l, r := 0, len(nums)-1
	for l < r {
		mid := (l + r) / 2
		if nums[mid] < nums[mid+1] {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return l
}

func findPeakElement(nums []int) int {
	nums = append(nums, math.MinInt32)
	nums = append([]int{math.MinInt32}, nums...)
	l, r := 1, len(nums)-2
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] < nums[mid+1] {
			l = mid + 1
		} else if nums[mid] < nums[mid-1] {
			r = mid - 1
		} else {
			return mid - 1
		}
	}
	return 0
}
