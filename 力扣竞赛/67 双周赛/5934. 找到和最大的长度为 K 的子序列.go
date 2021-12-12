package _7_双周赛

import "sort"

//给你一个整数数组 nums 和一个整数 k 。你需要找到 nums 中长度为 k 的 子序列 ，且这个子序列的 和最大 。
//
//请你返回 任意 一个长度为 k 的整数子序列。
//
//子序列 定义为从一个数组里删除一些元素后，不改变剩下元素的顺序得到的数组。
func maxSubsequence(nums []int, k int) []int {
	n := len(nums)
	a := make([]int, n)
	copy(a, nums)
	sort.Ints(a)
	mp := make(map[int]int)
	for i := n - 1; i >= 0; i-- {
		if k == 0 {
			break
		}
		mp[a[i]]++
		k--
	}
	var ans []int
	for _, v := range nums {
		if mp[v] > 0 {
			ans = append(ans, v)
			mp[v]--
		}
	}
	return ans
}
