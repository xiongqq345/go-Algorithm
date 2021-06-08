package dp

// 给你一个 只包含正整数 的 非空 数组 nums 。请你判断是否可以将这个数组分割成两个子集，使得两个子集的元素和相等。
func canPartition(nums []int) bool {
	var sum int
	for _, v := range nums {
		sum += v
	}
	if sum%2 == 1 {
		return false
	}
	target := sum / 2

	dp := make([]bool, target+1)
	dp[0] = true
	for _, v := range nums {
		for i := target; i >= v; i-- {
			dp[i] = dp[i] || dp[i-v]
		}
	}
	return dp[target]
}
