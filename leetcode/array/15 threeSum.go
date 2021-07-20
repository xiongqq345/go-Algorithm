package array

import "sort"

//给你一个包含 n 个整数的数组nums，判断nums中是否存在三个元素 a，b，c ，使得a + b + c = 0 ？请你找出所有和为 0 且不重复的三元组。
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	var ans [][]int
	for i := 0; i < len(nums)-2; i++ {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j, k := i+1, len(nums)-1; j < k; {
			if j != i+1 && nums[j] == nums[j-1] {
				j++
				continue
			}
			if nums[i]+nums[j]+nums[k] < 0 {
				j++
			} else if nums[i]+nums[j]+nums[k] > 0 {
				k--
			} else {
				ans = append(ans, []int{nums[i], nums[j], nums[k]})
				j++
			}
		}
	}
	return ans
}
