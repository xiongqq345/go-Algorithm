package _90

import (
	"sort"
)

func subsetsWithDup(nums []int) [][]int {
	res := make([][]int, 0)
	sort.Ints(nums)
	backtrack(&res, nums, []int{}, 0)
	return res
}

func backtrack(res *[][]int, nums, temp []int, start int) {
	tmp := make([]int, len(temp))
	copy(tmp, temp)
	*res = append(*res, tmp)
	for i := start; i < len(nums); i++ {
		if i > start && nums[i] == nums[i-1] {
			continue
		}
		temp = append(temp, nums[i])
		backtrack(res, nums, temp, i+1)
		temp = temp[:len(temp)-1]
	}
}
