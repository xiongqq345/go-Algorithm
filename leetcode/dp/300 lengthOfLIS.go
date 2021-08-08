package dp

import "sort"

//给你一个整数数组 nums ，找到其中最长严格递增子序列的长度。
//
//子序列是由数组派生而来的序列，删除（或不删除）数组中的元素而不改变其余元素的顺序。例如，[3,6,2,7] 是数组 [0,3,1,6,2,2,7] 的子序列。
func lengthOfLIS(nums []int) int {
	n, res := len(nums), 0
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		res = max(res, dp[i])
	}
	return res
}

func lengthOfLIS2(nums []int) int {
	d := []int{nums[0]}
	for _, v := range nums {
		if v > d[len(d)-1] {
			d = append(d, v)
		} else {
			d[sort.SearchInts(d, v)] = v
		}
	}
	return len(d)
}
