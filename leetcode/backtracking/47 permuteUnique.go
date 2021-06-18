package backtracking

import "sort"

// 给定一个可包含重复数字的序列 nums ，按任意顺序 返回所有不重复的全排列。
func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	var ans [][]int
	used := make([]bool, n)
	var backtrack func([]int)
	backtrack = func(vals []int) {
		if len(vals) == n {
			ans = append(ans, append([]int{}, vals...))
			return
		}

		for i := 0; i < n; i++ {
			if i > 0 && nums[i-1] == nums[i] && used[i-1] {
				continue
			}
			if used[i] {
				continue
			}

			used[i] = true
			backtrack(append(vals, nums[i]))
			used[i] = false
		}
	}
	backtrack([]int{})
	return ans
}
