package array

import (
	"math"
	"sort"
)

// 给定一个包括 n 个整数的数组 nums 和 一个目标值 target。找出 nums 中的三个整数，使得它们的和与 target 最接近。返回这三个数的和。假定每组输入只存在唯一答案。
func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	res, diff, n := 0, math.MaxInt32, len(nums)
	for i := 0; i < n-2; i++ {
		for j, k := i+1, n-1; j < k; {
			sum := nums[i] + nums[j] + nums[k]
			if diff > abs(sum-target) {
				res, diff = sum, abs(sum-target)
			}
			if diff == 0 {
				return res
			}
			if sum > target {
				k--
			} else {
				j++
			}
		}
	}
	return res
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}
