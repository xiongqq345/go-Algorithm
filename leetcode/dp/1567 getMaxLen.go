package dp

// 给你一个整数数组 nums ，请你求出乘积为正数的最长子数组的长度。
//
//一个数组的子数组是由原数组中零个或者更多个连续数字组成的数组。
//
//请你返回乘积为正数的最长子数组长度。
func getMaxLen(nums []int) int {
	var positive, negative, ans int
	for _, v := range nums {
		if v < 0 {
			mx := positive
			if negative > 0 {
				positive = negative + 1
			} else {
				positive = 0
			}
			negative = mx + 1
		} else if v > 0 {
			positive++
			if negative > 0 {
				negative++
			}
		} else {
			positive, negative = 0, 0
		}
		ans = max(ans, positive)
	}
	return ans
}
