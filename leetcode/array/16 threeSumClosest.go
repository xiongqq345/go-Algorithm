package array

import (
	"math"
	"sort"
)

// 给定一个包括n 个整数的数组nums和 一个目标值target。找出nums中的三个整数，使得它们的和与target最接近。返回这三个数的和。假定每组输入只存在唯一答案。
func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	ans, diff, n := 0, math.MaxInt32, len(nums)
	for i := 0; i < n-2; i++ {
		for j, k := i+1, n-1; j < k; {
			sum := nums[i] + nums[j] + nums[k]
			if diff > abs(sum-target) {
				ans, diff = sum, abs(sum-target)
			}
			if diff == 0 {
				return ans
			}
			if sum > target {
				k--
			} else {
				j++
			}
		}
	}
	return ans
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}
