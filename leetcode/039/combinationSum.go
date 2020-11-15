package leetcode

import (
	"sort"
)

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	var res [][]int
	findCombination(candidates, []int{}, 0, target, &res)
	return res
}

func findCombination(candidates []int, combination []int, index, target int, res *[][]int) {
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
		combination = append(combination, candidates[i])
		findCombination(candidates, combination, i, target-candidates[i], res)
		combination = combination[:len(combination)-1]
	}
}
