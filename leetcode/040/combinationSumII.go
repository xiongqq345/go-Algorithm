package leetcode

import (
	"sort"
)

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	var res [][]int
	findCombination2(candidates, []int{}, 0, target, &res)
	return res
}

func findCombination2(candidates []int, combination []int, index, target int, res *[][]int) {
	if target < 0 {
		return
	}

	if target == 0 {
		c := make([]int, len(combination))
		copy(c, combination)
		*res = append(*res, c)
		return
	}

	for i := index; i < len(candidates); i++ {
		if i > index && candidates[i] == candidates[i-1] {
			continue
		}
		if candidates[i] <= target {
			combination = append(combination, candidates[i])
			findCombination2(candidates, combination, i+1, target-candidates[i], res)
			combination = combination[:len(combination)-1]
		}
	}
}
