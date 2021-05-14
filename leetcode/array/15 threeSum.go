package array

import "sort"

//给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有和为 0 且不重复的三元组。
func threeSum(nums []int) [][]int {
	n := len(nums)
	sort.Ints(nums)
	var ans [][]int
	for first := 0; first < n-2; first++ {
		if nums[first] > 0 {
			break
		}
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		third := n - 1
		target := -nums[first]
		for second := first + 1; second < n-1; second++ {
			if second > first+1 && nums[second] == nums[second-1] {
				continue
			}

			for nums[second]+nums[third] > target && second < third {
				third--
			}

			if second == third {
				break
			}

			if nums[second]+nums[third] == target {
				ans = append(ans, []int{nums[first], nums[second], nums[third]})
			}
		}
	}
	return ans
}
