package jz_offerII

import "sort"

//给定一个可包含重复数字的整数集合 nums ，按任意顺序 返回它所有不重复的全排列。
func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	vis := make([]bool, n)
	var ans [][]int
	var helper func(path []int)
	helper = func(path []int) {
		if len(path) == n {
			ans = append(ans, append([]int(nil), path...))
			return
		}

		for i := range nums {
			if i > 0 && nums[i] == nums[i-1] && vis[i-1] {
				continue
			}
			if !vis[i] {
				vis[i] = true
				helper(append(path, nums[i]))
				vis[i] = false
			}
		}
	}
	helper(nil)
	return ans
}
