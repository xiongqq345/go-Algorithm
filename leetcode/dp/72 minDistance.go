package dp

//给你两个单词word1 和word2，请你计算出将word1转换成word2 所使用的最少操作数。
//
//你可以对一个单词进行如下三种操作：
//
//插入一个字符
//删除一个字符
//替换一个字符

func minDistance(word1 string, word2 string) int {
	dp := make([]int, len(word2)+1)
	for i := range dp {
		dp[i] = i
	}

	for i := range word1 {
		upLeft := i
		dp[0]++
		for j := range word2 {
			tmp := dp[j+1]
			if word1[i] == word2[j] {
				dp[j+1] = upLeft
			} else {
				dp[j+1] = min(upLeft, min(dp[j], dp[j+1])) + 1
			}
			upLeft = tmp
		}
	}
	return dp[len(word2)]
}
