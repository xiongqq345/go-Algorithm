package greedy

import "math"

// 给你一个整数数组 nums ，判断这个数组中是否存在长度为 3 的递增子序列。
//
//如果存在这样的三元组下标 (i, j, k) 且满足 i < j < k ，使得 nums[i] < nums[j] < nums[k] ，返回 true ；否则，返回 false 。

func increasingTriplet(nums []int) bool {
	small, mid := math.MaxInt32, math.MaxInt32
	for _, v := range nums {
		switch {
		case v <= small:
			small = v
		case v <= mid:
			mid = v
		default:
			return true
		}
	}
	return false
}
