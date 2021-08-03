package array

import "math"

// 给你一个整数数组 nums ，你需要找出一个 连续子数组 ，如果对这个子数组进行升序排序，那么整个数组都会变为升序排序。
//
//请你找出符合题意的 最短 子数组，并输出它的长度。
//
func findUnsortedSubarray(nums []int) int {
	n := len(nums)
	l, r := -1, -1
	maxn, minn := math.MinInt32, math.MaxInt32
	for i := 0; i < n; i++ {
		if nums[i] < maxn {
			r = i
		} else {
			maxn = nums[i]
		}
		if nums[n-i-1] > minn {
			l = n - i - 1
		} else {
			minn = nums[n-i-1]
		}
	}
	if l == -1 {
		return 0
	}
	return r - l + 1
}
