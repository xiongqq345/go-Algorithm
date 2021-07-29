package backtracking

import "sort"

// 给定一个可包含重复数字的序列 nums ，按任意顺序 返回所有不重复的全排列。
func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	var ans [][]int
	vis := make([]bool, n)
	var helper func([]int)
	helper = func(vals []int) {
		if len(vals) == n {
			ans = append(ans, append([]int{}, vals...))
			return
		}

		for i := 0; i < n; i++ {
			if i > 0 && nums[i-1] == nums[i] && vis[i-1] {
				continue
			}
			if vis[i] {
				continue
			}

			vis[i] = true
			helper(append(vals, nums[i]))
			vis[i] = false
		}
	}
	helper([]int{})
	return ans
}
