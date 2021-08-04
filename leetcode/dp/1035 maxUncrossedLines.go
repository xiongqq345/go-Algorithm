package dp

// 在两条独立的水平线上按给定的顺序写下 nums1 和 nums2 中的整数。
//
//现在，可以绘制一些连接两个数字 nums1[i] 和 nums2[j] 的直线，这些直线需要同时满足满足：
//
// nums1[i] == nums2[j]
//且绘制的直线不与任何其他连线（非水平线）相交。
//请注意，连线即使在端点也不能相交：每个数字只能属于一条连线。
//
//以这种方法绘制线条，并返回可以绘制的最大连线数。
func maxUncrossedLines(nums1 []int, nums2 []int) int {
	dp := make([]int, len(nums2)+1)
	for i := range nums1 {
		var upLeft int
		for j := range nums2 {
			tmp := dp[j+1]
			if nums1[i] == nums2[j] {
				dp[j+1] = upLeft + 1
			} else {
				dp[j+1] = max(dp[j+1], dp[j])
			}
			upLeft = tmp
		}
	}
	return dp[len(nums2)]
}
