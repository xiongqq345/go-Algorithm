package jz_offerII

import "sort"

//给定一个可能有重复数字的整数数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
//
//candidates 中的每个数字在每个组合中只能使用一次，解集不能包含重复的组合。
//
func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	var ans [][]int
	var helper func([]int, int, int)
	helper = func(path []int, pos, sum int) {
		if sum > target {
			return
		}
		if sum == target {
			ans = append(ans, append([]int(nil), path...))
			return
		}

		for i := pos; i < len(candidates); i++ {
			if i > pos && candidates[i] == candidates[i-1] {
				continue
			}
			helper(append(path, candidates[i]), i+1, sum+candidates[i])
		}
	}
	helper(nil, 0, 0)
	return ans
}
