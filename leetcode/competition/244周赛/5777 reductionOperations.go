package _44周赛

import "sort"

// 使数组元素相等的减少操作次数
func reductionOperations(nums []int) int {
	sort.Ints(nums)
	var base, ans int
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			base++
		}
		ans += base
	}
	return ans
}
