package jz_offer

// 输入一个整数数组，实现一个函数来调整该数组中数字的顺序，使得所有奇数位于数组的前半部分，所有偶数位于数组的后半部分。
func exchange(nums []int) []int {
	for l, r := 0, len(nums)-1; l < r; {
		if nums[l]&1 == 0 {
			nums[l], nums[r] = nums[r], nums[l]
			r--
		} else {
			l++
		}
	}
	return nums
}
