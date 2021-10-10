package jz_offerII

//给定一个排序的整数数组 nums 和一个整数目标值 target ，请在数组中找到 target ，并返回其下标。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。
//
func searchInsert(nums []int, target int) int {
	i, j := 0, len(nums)
	for i < j {
		h := int(uint(i+j) >> 1)
		if !(nums[h] >= target) {
			i = h + 1
		} else {
			j = h
		}
	}
	return i
}
