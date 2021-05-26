package string

func lengthOfLongestSubstring(s string) int {
	var set [128]bool
	j, ans := 0, 0
	for i := range s {
		if i != 0 {
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
