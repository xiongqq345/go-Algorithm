package string

func lengthOfLongestSubstring(s string) int {
	var set [128]bool
	var ans int
	for i, j := 0, 0; i < len(s); i++ {
		if i > 0 {
			set[s[i-1]] = false
		}
		for j < len(s) && !set[s[j]] {
			set[s[j]] = true
			j++
		}
		ans = max(ans, j-i)
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
