package string

func lengthOfLongestSubstring(s string) int {
	var m [128]bool
	j, ans := 0, 0
	for i := range s {
		if i != 0 {
			m[s[i-1]] = false
		}
		for j < len(s) && !m[s[j]] {
			m[s[j]] = true
			j++
		}
		ans = max(j-i, ans)
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
