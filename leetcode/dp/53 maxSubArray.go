package dp

// 给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
func maxSubArray(nums []int) int {
	res, dp := nums[0], 0
	for _, v := range nums {
		dp = max(v, dp+v)
		res = max(res, dp)
	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
