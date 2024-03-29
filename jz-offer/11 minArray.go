package jz_offer

//把一个数组最开始的若干个元素搬到数组的末尾，我们称之为数组的旋转。输入一个递增排序的数组的一个旋转，输出旋转数组的最小元素。例如，数组 [3,4,5,1,2] 为 [1,2,3,4,5] 的一个旋转，该数组的最小值为1。
func minArray(nums []int) int {
	i, j := 0, len(nums)-1
	for i < j {
		h := int(uint(i+j) >> 1)
		if nums[h] > nums[j] {
			i = h + 1
		} else if nums[h] < nums[j] {
			j = h
		} else {
			j--
		}
	}
	return nums[i]
}
