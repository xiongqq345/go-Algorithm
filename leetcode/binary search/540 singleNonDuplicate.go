package binary_search

import "sort"

// 给定一个只包含整数的有序数组，每个元素都会出现两次，唯有一个数只会出现一次，找出这个数。
func singleNonDuplicate(nums []int) int {
	return nums[sort.Search(len(nums)-1, func(i int) bool {
		if i%2 == 0 {
			return nums[i] < nums[i+1]
		}
		return nums[i] > nums[i-1]
	})]
}
