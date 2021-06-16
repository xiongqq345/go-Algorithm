package array

// 给你一个未排序的整数数组 nums ，请你找出其中没有出现的最小的正整数。
//
//请你实现时间复杂度为 O(n) 并且只使用常数级别额外空间的解决方案。
func firstMissingPositive(nums []int) int {
	n := len(nums)
	for i, v := range nums {
		if v <= 0 {
			nums[i] = n + 1
		}
	}
	for _, v := range nums {
		if abs(v) > 0 && abs(v) <= n {
			if nums[abs(v)-1] > 0 {
				nums[abs(v)-1] *= -1
			}
		}
	}
	for i := 0; i < n; i++ {
		if nums[i] > 0 {
			return i + 1
		}
	}
	return n + 1
}
