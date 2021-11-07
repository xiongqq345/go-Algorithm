package dp

//给你一个整数数组 arr 和一个整数 difference，请你找出并返回 arr 中最长等差子序列的长度，该子序列中相邻元素之间的差等于 difference 。
//
//子序列 是指在不改变其余元素顺序的情况下，通过删除一些元素或不删除任何元素而从 arr 派生出来的序列。
//

func longestSubsequence(arr []int, difference int) int {
	dp := make(map[int]int)
	var ans int
	for _, v := range arr {
		dp[v] = dp[v-difference] + 1
		ans = max(ans, dp[v])
	}
	return ans
}
