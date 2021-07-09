package string

// 给你一个字符串 s 、一个字符串 t 。返回 s 中涵盖 t 所有字符的最小子串。如果 s 中不存在涵盖 t 所有字符的子串，则返回空字符串 "" 。
//
//注意：如果 s 中存在这样的子串，我们保证它是唯一的答案。
func minWindow(s string, t string) string {
	if len(s) < len(t) {
		return ""
	}
	mp := make([]int, 256)
	for _, ch := range t {
		mp[ch]++
	}

	l, count, min, res := 0, len(t), len(s)+1, ""
	for r, ch := range s {
		mp[ch]--
		if mp[ch] >= 0 {
			count--
		}
		for l < r && mp[s[l]] < 0 {
			mp[s[l]]++
			l++
		}
		if count == 0 && min > r-l+1 {
			min = r - l + 1
			res = s[l : r+1]
		}
	}

	return res
}
