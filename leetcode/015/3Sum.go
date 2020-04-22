package leetcode

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	var res [][]int
	for i := 0; i < len(nums)-2; i++ {
		first := nums[i]
		if first > 0 {
			break
		}
		p1, p2 := i+1, len(nums)-1
		for p1 < p2 {
			sum := first + nums[p1] + nums[p2]
			if sum == 0 {
				res = append(res, []int{first, nums[p1], nums[p2]})
				for p1 < p2 && nums[p1] == nums[p1+1] {
					p1++
				}
				for p1 < p2 && nums[p2] == nums[p2-1] {
					p2--
				}
				p1++
				p2--
			} else if sum < 0 {
				p1++
			} else if sum > 0 {
				p2--
			}
		}
		for i<len(nums)-3 && nums[i]==nums[i+1] {
			i++
		}
	}
	return res
}
