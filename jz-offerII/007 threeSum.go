package jz_offerII

import "sort"

// 给定一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a ，b ，c ，使得 a + b + c = 0 ？请找出所有和为 0 且 不重复 的三元组。
//
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	var ans [][]int
	n := len(nums)
	for i := 0; i < n-2; i++ {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		a := nums[i]
		for j, k := i+1, n-1; j < k; {
			b, c := nums[j], nums[k]
			if j > i+1 && b == nums[j-1] {
				j++
				continue
			}
			if a+b+c < 0 {
				j++
			} else if a+b+c > 0 {
				k--
			} else {
				ans = append(ans, []int{a, b, c})
				j++
			}
		}
	}
	return ans
}
