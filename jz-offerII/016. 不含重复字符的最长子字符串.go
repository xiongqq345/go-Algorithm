package jz_offerII

//给定一个字符串 s ，请你找出其中不含有重复字符的 最长连续子字符串 的长度。
func lengthOfLongestSubstring(s string) int {
	var ans int
	var set [128]bool
	for i, j := 0, 0; j < len(s); j++ {
		for set[s[j]] {
			set[s[i]] = false
			i++
		}
		set[s[j]] = true
		ans = max(ans, j-i+1)
	}
	return ans
}
