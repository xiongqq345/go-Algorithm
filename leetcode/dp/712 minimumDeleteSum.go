package dp

// 给定两个字符串s1, s2，找到使两个字符串相等所需删除字符的ASCII值的最小和。
func minimumDeleteSum(s1 string, s2 string) int {
	dp := make([]int, len(s2)+1)
	var sum1, sum2 int
	for _, v := range s1 {
		sum1 += int(v)
	}
	for _, v := range s2 {
		sum2 += int(v)
	}

	for i := range s1 {
		var upLeft int
		for j := range s2 {
			tmp := dp[j+1]
			if s1[i] == s2[j] {
				dp[j+1] = upLeft + int(s2[j])
			} else {
				dp[j+1] = max(dp[j], dp[j+1])
			}
			upLeft = tmp
		}
	}
	return sum1 + sum2 - 2*(dp[len(s2)])
}
