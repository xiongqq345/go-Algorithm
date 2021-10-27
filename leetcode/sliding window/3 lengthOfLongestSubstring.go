package sliding_window

func lengthOfLongestSubstring(s string) int {
	var ans int
	var pos [256]int
	for i, j := 0, 0; j < len(s); j++ {
		i = max(i, pos[s[j]])
		ans = max(ans, j-i+1)
		pos[s[j]] = j + 1
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
