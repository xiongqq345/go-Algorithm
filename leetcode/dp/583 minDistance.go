package dp

//给定两个单词 word1 和 word2，找到使得 word1 和 word2 相同所需的最小步数，每步可以删除任意一个字符串中的一个字符。
func minDistance(word1 string, word2 string) int {
	dp := make([]int, len(word2)+1)
	for i := range word1 {
		var upLeft int
		for j := range word2 {
			tmp := dp[j+1]
			if word1[i] == word2[j] {
				dp[j+1] = upLeft + 1
			} else {
				dp[j+1] = max(dp[j], dp[j+1])
			}
			upLeft = tmp
		}
	}
	return len(word1) + len(word2) - 2*dp[len(word2)]
}
