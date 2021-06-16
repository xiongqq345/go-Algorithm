package array

import "sort"

//数组中有一个数字出现的次数超过数组长度的一半，请找出这个数字。
//你可以假设数组是非空的，并且给定的数组总是存在多数元素。
func majorityElement(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)/2]
}

func majorityElement2(nums []int) int {
	var count, candidate int
	for _, num := range nums {
		if count == 0 {
			candidate = num
			count++
		} else {
			if num == candidate {
				count++
			} else {
				count--
			}
		}
	}
	return candidate
}
