package backtracking

import "sort"

// 给定一个数组candidates和一个目标数target，找出candidates中所有可以使数字和为target的组合。
//
//candidates中的每个数字在每个组合中只能使用一次。
//
//说明：
//所有数字（包括目标数）都是正整数。
//解集不能包含重复的组合。

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	var ans [][]int
	var helper func(vals []int, sum, index int)
	helper = func(vals []int, sum, index int) {
		if sum > target {
			return
		}
		if sum == target {
			ans = append(ans, append([]int(nil), vals...))
			return
		}
		for i := index; i < len(candidates); i++ {
			if i > index && candidates[i] == candidates[i-1] {
				continue
			}
			helper(append(vals, candidates[i]), sum+candidates[i], i+1)
		}
	}
	helper([]int{}, 0, 0)
	return ans
}
