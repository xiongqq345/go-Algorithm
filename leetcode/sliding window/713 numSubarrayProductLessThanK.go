package sliding_window

//给定一个正整数数组 nums和整数 k 。
//
//请找出该数组内乘积小于 k 的连续的子数组的个数。

func numSubarrayProductLessThanK(nums []int, k int) int {
	if k <= 1 {
		return 0
	}
	cnt, product := 0, 1
	for i, j := 0, 0; j < len(nums); j++ {
		product *= nums[j]
		for product >= k {
			product /= nums[i]
			i++
		}
		cnt += j - i + 1
	}
	return cnt
}
