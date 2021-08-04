package dp

//给定一个二进制数组，你可以最多将 1 个 0 翻转为 1，找出其中最大连续 1 的个数。
func findMaxConsecutiveOnes(nums []int) int {
	var dp0, dp1 int
	var ans int
	for _, num := range nums {
		if num == 1 {
			dp0++
			dp1++
		} else {
			dp1 = dp0 + 1
			dp0 = 0
		}
		ans = max(ans, dp1)
	}
	return ans
}
