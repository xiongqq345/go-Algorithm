package jz_offer

// 请从字符串中找出一个最长的不包含重复字符的子字符串，计算该最长子字符串的长度。
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
