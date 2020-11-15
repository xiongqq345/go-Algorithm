package leetcode

import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	res, n, diff := 0, len(nums), math.MaxInt32
	sort.Ints(nums)
	for i := 0; i < n-1; i++ {
		for j, k := i+1, n-1; j < k; {
			sum := nums[i] + nums[j] + nums[k]
			if diff > abs(target-sum) {
				res, diff = sum, abs(target-sum)
			}
			if sum == target {
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
