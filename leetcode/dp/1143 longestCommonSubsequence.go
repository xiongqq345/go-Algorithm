package dp

// 给定两个字符串text1 和text2，返回这两个字符串的最长 公共子序列 的长度。如果不存在 公共子序列 ，返回 0 。
//
//一个字符串的子序列是指这样一个新的字符串：它是由原字符串在不改变字符的相对顺序的情况下删除某些字符（也可以不删除任何字符）后组成的新字符串。
//
//例如，"ace" 是 "abcde" 的子序列，但 "aec" 不是 "abcde" 的子序列。
//两个字符串的 公共子序列 是这两个字符串所共同拥有的子序列。
// 最长公共子序列
func longestCommonSubsequence(text1 string, text2 string) int {
	dp := make([]int, len(text2)+1)
	for i := range text1 {
		var upLeft int
		for j := range text2 {
			tmp := dp[j+1]
			if text1[i] == text2[j] {
				dp[j+1] = upLeft + 1
			} else {
				dp[j+1] = max(dp[j+1], dp[j])
			}
			upLeft = tmp
		}
	}
	return dp[len(text2)]
}
