package math

//给你一个整数 n 。按下述规则生成一个长度为 n + 1 的数组 nums ：
//
//nums[0] = 0
//nums[1] = 1
//当 2 <= 2 * i <= n 时，nums[2 * i] = nums[i]
//当 2 <= 2 * i + 1 <= n 时，nums[2 * i + 1] = nums[i] + nums[i + 1]
//返回生成数组 nums 中的 最大 值。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/get-maximum-in-generated-array
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func getMaximumGenerated(n int) int {
	if n == 0 {
		return 0
	}

	nums := make([]int, n+1)
	nums[1] = 1
	ans := 1
	for i := 2; i <= n; i++ {
		if i%2 == 0 {
			nums[i] = nums[i/2]
		} else {
			nums[i] = nums[i/2] + nums[i/2+1]
		}
		ans = max(ans, nums[i])
	}
	return ans
}
