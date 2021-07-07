package array

// 给你一个按 非递减顺序 排序的整数数组 nums，返回 每个数字的平方 组成的新数组，要求也按 非递减顺序 排序。
func sortedSquares(nums []int) []int {
	ans := make([]int, len(nums))
	pos := len(nums) - 1
	i, j := 0, len(nums)-1
	for i <= j {
		if abs(nums[i]) < abs(nums[j]) {
			ans[pos] = nums[j] * nums[j]
			j--
		} else {
			ans[pos] = nums[i] * nums[i]
			i++
		}
		pos--
	}
	return ans
}
