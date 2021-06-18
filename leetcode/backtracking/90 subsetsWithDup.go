package backtracking

import "sort"

// 给你一个整数数组 nums ，其中可能包含重复元素，请你返回该数组所有可能的子集（幂集）。
//
//解集 不能 包含重复的子集。返回的解集中，子集可以按 任意顺序 排列。
func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	var ans [][]int
	var set []int
	var helper func(int)
	helper = func(index int) {
		ans = append(ans, append([]int{}, set...))
		for i := index; i < len(nums); i++ {
			if i > index && nums[i-1] == nums[i] {
				continue
			}
			set = append(set, nums[i])
			helper(i + 1)
			set = set[:len(set)-1]
		}
	}

	helper(0)
	return ans
}
